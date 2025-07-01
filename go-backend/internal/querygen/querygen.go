package querygen

import "fmt"

// BuildGeminiPrompt constructs the prompt for Gemini given the schema CSV and transcript.
func BuildGeminiPrompt(schemaCSV, transcript string) string {
	return fmt.Sprintf(`You are a SQL expert. Given the following database schema in CSV format:

%s

And the following user transcript:

%s

Generate a single valid SQL query (SQLite dialect) that answers the user's request. Only output the SQL query.`, schemaCSV, transcript)
}
