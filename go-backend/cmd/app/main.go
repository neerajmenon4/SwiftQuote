package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/yourusername/sales-quotation-ai/internal/db"
	"github.com/yourusername/sales-quotation-ai/internal/gemini"
	"github.com/yourusername/sales-quotation-ai/internal/querygen"
	"github.com/yourusername/sales-quotation-ai/internal/schema"
)

// ReadTranscript reads the entire contents of a transcript file and returns it as a string.
func ReadTranscript(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	schemaPath := flag.String("schema", "", "Path to schema CSV file")
	transcriptPath := flag.String("transcript", "", "Path to transcript text file")
	dbPath := flag.String("db", "/Users/neerajmenon/Documents/Projects/Sales-Quotation-AI/db/sales_quotation_ai_mock.db", "Path to SQLite database file")
	executeQuery := flag.Bool("execute", true, "Execute the generated SQL query against the database")
	flag.Parse()

	if *schemaPath == "" || *transcriptPath == "" {
		log.Fatal("Usage: go run ./cmd/app --schema <schema.csv> --transcript <transcript.txt> [--db <database.db>] [--execute=false]")
	}

	schemaCSV, err := schema.ReadSchemaCSV(*schemaPath)
	if err != nil {
		log.Fatalf("Failed to read schema CSV: %v", err)
	}

	transcript, err := ReadTranscript(*transcriptPath)
	if err != nil {
		log.Fatalf("Failed to read transcript: %v", err)
	}

	fmt.Println("=== SCHEMA CSV ===")
	fmt.Println(schemaCSV)
	fmt.Println("=== TRANSCRIPT ===")
	fmt.Println(transcript)

	prompt := querygen.BuildGeminiPrompt(schemaCSV, transcript)
	fmt.Println("\n=== GEMINI PROMPT ===")
	fmt.Println(prompt)

	// Load .env file
	err = godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, proceeding to use environment variables.")
	}

	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		log.Fatal("GOOGLE_API_KEY not set in environment or .env file")
	}

	fmt.Println("\nCalling Gemini API...")
	result, err := gemini.CallGemini(prompt, apiKey)
	if err != nil {
		log.Fatalf("gemini API error: %v", err)
	}

	fmt.Println("\n=== GEMINI SQL QUERY ===")
	fmt.Println(result)

	// Clean up the SQL query (remove markdown formatting if present)
	sqlQuery := strings.TrimSpace(result)

	// Handle markdown code blocks with language specifiers
	if strings.HasPrefix(sqlQuery, "```") {
		// Find the first line break which ends the opening markdown fence
		firstLineBreak := strings.Index(sqlQuery, "\n")
		if firstLineBreak != -1 {
			// Remove everything up to and including that first line break
			sqlQuery = sqlQuery[firstLineBreak+1:]
		}

		// Remove the closing markdown fence
		sqlQuery = strings.TrimSuffix(sqlQuery, "```")
	}

	sqlQuery = strings.TrimSpace(sqlQuery)

	if *executeQuery {
		fmt.Println("\n=== EXECUTING SQL QUERY ===")
		// Connect to the database
		database, err := db.OpenDatabase(*dbPath)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}
		defer database.Close()

		// Execute the query
		queryResults, err := db.ExecuteQuery(database, sqlQuery)
		if err != nil {
			log.Fatalf("Failed to execute query: %v", err)
		}

		fmt.Println("\n=== QUERY RESULTS ===")
		fmt.Println(queryResults)
	}
}
