package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/yourusername/sales-quotation-ai/internal/db"
	"github.com/yourusername/sales-quotation-ai/internal/gemini"
	"github.com/yourusername/sales-quotation-ai/internal/querygen"
	"github.com/yourusername/sales-quotation-ai/internal/schema"
)

// Server configuration
const (
	PORT           = 8080
	UPLOAD_DIR     = "./uploads"
	SCHEMA_CSV     = "/Users/neerajmenon/Documents/Projects/Sales-Quotation-AI/schemas/db_schema.csv"
	DB_PATH        = "/Users/neerajmenon/Documents/Projects/Sales-Quotation-AI/db/sales_quotation_ai_mock.db"
)

// Response structure for the API
type AnalysisResponse struct {
	Query string      `json:"query"`
	Data  interface{} `json:"data"`
}

func main() {
	// Create upload directory if it doesn't exist
	if err := os.MkdirAll(UPLOAD_DIR, 0755); err != nil {
		log.Fatalf("Failed to create upload directory: %v", err)
	}

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, proceeding to use environment variables.")
	}

	// Check for API key
	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		log.Fatal("GOOGLE_API_KEY not set in environment or .env file")
	}

	// Read schema CSV
	schemaCSV, err := schema.ReadSchemaCSV(SCHEMA_CSV)
	if err != nil {
		log.Fatalf("Failed to read schema CSV: %v", err)
	}

	// Set up router
	r := mux.NewRouter()
	r.HandleFunc("/api/analyze", func(w http.ResponseWriter, r *http.Request) {
		handleAnalyzeTranscript(w, r, schemaCSV, apiKey)
	}).Methods("POST")

	// Add CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	// Start server
	fmt.Printf("Server starting on port %d...\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), handler))
}

func handleAnalyzeTranscript(w http.ResponseWriter, r *http.Request, schemaCSV string, apiKey string) {
	// Parse multipart form with 10MB max memory
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Get uploaded file
	file, handler, err := r.FormFile("transcript")
	if err != nil {
		http.Error(w, "Failed to get transcript file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create temporary file
	tempFilePath := filepath.Join(UPLOAD_DIR, handler.Filename)
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		http.Error(w, "Failed to save transcript file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempFilePath) // Clean up after processing

	// Copy uploaded file to temp file
	if _, err := io.Copy(tempFile, file); err != nil {
		http.Error(w, "Failed to save transcript file", http.StatusInternalServerError)
		return
	}
	tempFile.Close() // Close to ensure all data is written

	// Read transcript file
	transcript, err := readTranscript(tempFilePath)
	if err != nil {
		http.Error(w, "Failed to read transcript file", http.StatusInternalServerError)
		return
	}

	// Build prompt and call Gemini API
	prompt := querygen.BuildGeminiPrompt(schemaCSV, transcript)
	result, err := gemini.CallGemini(prompt, apiKey)
	if err != nil {
		http.Error(w, fmt.Sprintf("Gemini API error: %v", err), http.StatusInternalServerError)
		return
	}

	// Clean up the SQL query (remove markdown formatting if present)
	sqlQuery := strings.TrimSpace(result)
	if strings.HasPrefix(sqlQuery, "```") {
		firstLineBreak := strings.Index(sqlQuery, "\n")
		if firstLineBreak != -1 {
			sqlQuery = sqlQuery[firstLineBreak+1:]
		}
		sqlQuery = strings.TrimSuffix(sqlQuery, "```")
	}
	sqlQuery = strings.TrimSpace(sqlQuery)

	// Connect to the database
	database, err := db.OpenDatabase(DB_PATH)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to open database: %v", err), http.StatusInternalServerError)
		return
	}
	defer database.Close()

	// Execute the query
	queryResults, err := db.ExecuteQuery(database, sqlQuery)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to execute query: %v", err), http.StatusInternalServerError)
		return
	}

	// Parse the JSON string returned by ExecuteQuery
	var dataArray []map[string]interface{}
	if err := json.Unmarshal([]byte(queryResults), &dataArray); err != nil {
		// If parsing fails, create an empty array
		dataArray = []map[string]interface{}{}
		log.Printf("Warning: Failed to parse query results: %v", err)
	}
	
	// Prepare response
	response := AnalysisResponse{
		Query: sqlQuery,
		Data:  dataArray,
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// readTranscript reads the entire contents of a transcript file and returns it as a string.
func readTranscript(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
