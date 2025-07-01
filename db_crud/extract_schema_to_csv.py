import sqlite3
import csv

DB_FILE = '../db/sales_quotation_ai_mock.db'
CSV_FILE = 'db_schema.csv'

conn = sqlite3.connect(DB_FILE)
c = conn.cursor()

# Get all table names
c.execute("SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%';")
tables = [row[0] for row in c.fetchall()]

schema_rows = []

for table in tables:
    # Get columns info
    c.execute(f'PRAGMA table_info({table});')
    columns = c.fetchall()  # cid, name, type, notnull, dflt_value, pk
    col_dict = {col[1]: col for col in columns}

    # Get foreign keys
    c.execute(f'PRAGMA foreign_key_list({table});')
    fkeys = c.fetchall()  # id, seq, table, from, to, on_update, on_delete, match
    fkey_map = {fk[3]: (fk[2], fk[4]) for fk in fkeys}  # from_col: (ref_table, ref_col)

    for col in columns:
        col_name = col[1]
        col_type = col[2]
        is_pk = 'YES' if col[5] else ''
        is_fk = 'YES' if col_name in fkey_map else ''
        ref_table = fkey_map[col_name][0] if col_name in fkey_map else ''
        ref_col = fkey_map[col_name][1] if col_name in fkey_map else ''
        schema_rows.append({
            'Table': table,
            'Column': col_name,
            'Type': col_type,
            'PrimaryKey': is_pk,
            'ForeignKey': is_fk,
            'ReferencesTable': ref_table,
            'ReferencesColumn': ref_col
        })

# Write to CSV
with open(CSV_FILE, 'w', newline='') as csvfile:
    fieldnames = ['Table', 'Column', 'Type', 'PrimaryKey', 'ForeignKey', 'ReferencesTable', 'ReferencesColumn']
    writer = csv.DictWriter(csvfile, fieldnames=fieldnames)
    writer.writeheader()
    for row in schema_rows:
        writer.writerow(row)

print(f"Schema extracted to {CSV_FILE}")
