# Sales Quotation AI

A full-stack application that uses AI to analyze sales transcripts and generate SQL queries to extract insights from a structured database. The backend is built with Go and SQLite, and the frontend is located in the `frontend/` directory.

---

## Prerequisites

- [Go](https://golang.org/doc/install) (v1.18 or higher recommended)
- [Node.js & npm](https://nodejs.org/) (for the frontend)
- [SQLite3](https://www.sqlite.org/download.html)

---

## 1. Clone the Repository

```bash
git clone https://github.com/yourusername/sales-quotation-ai.git
cd sales-quotation-ai
```

---

## 2. Database Setup

1. **Install SQLite3** if you don't have it:
   - macOS: `brew install sqlite3`
   - Ubuntu: `sudo apt-get install sqlite3`
   - Windows: [Download from sqlite.org](https://www.sqlite.org/download.html)

2. **Create and Populate the Database:**
   - The project expects the database at `db/sales_quotation_ai_mock.db`.
   - If it's missing or you want to recreate it:

```bash
sqlite3 db/sales_quotation_ai_mock.db < schemas/db_schema.csv
```
*Note: If `db_schema.csv` is not a SQL script, but a CSV, you may need to manually create tables and import data. Adjust as needed.*

---

## 3. Backend Setup (Go)

1. **Install Dependencies:**

```bash
cd go-backend
go mod tidy
```

2. **Create a `.env` file** in `go-backend/` with your Google API key:

```
GOOGLE_API_KEY=your_google_api_key_here
```

3. **Run the Go Server:**

```bash
cd cmd/server
go run main.go
```

The server will start on port 8080 by default.

---

## 4. Frontend Setup

1. **Install Dependencies:**

```bash
cd ../../frontend
npm install
# or
# yarn install
```

2. **Run the Frontend:**

```bash
npm start
# or
# yarn start
```

The frontend will typically run on [http://localhost:3000](http://localhost:3000).

---

## 5. Example .env File

```
GOOGLE_API_KEY=your_google_api_key_here
```

---

## 6. Troubleshooting

- Ensure your `.env` file is present and contains a valid Google API key.
- If the backend cannot find the database, make sure `db/sales_quotation_ai_mock.db` exists and is populated.
- The backend expects the schema CSV at `schemas/db_schema.csv`.
- If you encounter CORS errors, ensure the frontend is running on the default port (3000) or update the backend CORS settings.

---

## 7. Project Structure

```
├── db/                        # SQLite database files
├── schemas/                   # Database schema CSV
├── go-backend/                # Go backend source code
├── frontend/                  # Frontend source code (React/Vue/etc.)
├── requirements.txt           # (optional) Python utilities
└── README.md                  # This file
```

---

## 8. License

[MIT](LICENSE)

---

For further questions, please open an issue or contact the maintainer.
