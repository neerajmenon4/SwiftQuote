# Sales Quotation AI - Go Backend

This backend service uses Gemini to generate SQL queries from natural language transcripts and executes them on a SQLite database.

## Structure

- `cmd/app/` - Main entry point
- `internal/` - Application logic (Gemini, DB, schema, etc.)
- `testdata/` - Sample CSVs, transcripts, etc.

## Getting Started

1. `cd go-backend`
2. `go run ./cmd/app`
