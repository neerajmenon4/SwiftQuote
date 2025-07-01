import sqlite3

# Connect to SQLite DB (creates file if not exists)
conn = sqlite3.connect('sales_quotation_ai_mock.db')
c = conn.cursor()

# 1. Roles Table
c.execute('''
CREATE TABLE IF NOT EXISTS Roles (
    RoleID INTEGER PRIMARY KEY AUTOINCREMENT,
    RoleName TEXT NOT NULL UNIQUE
)
''')

# 2. Users Table
c.execute('''
CREATE TABLE IF NOT EXISTS Users (
    UserID INTEGER PRIMARY KEY AUTOINCREMENT,
    UserName TEXT NOT NULL,
    Email TEXT NOT NULL UNIQUE,
    RoleID INTEGER,
    FOREIGN KEY (RoleID) REFERENCES Roles(RoleID)
)
''')

# 3. ProductCatalog Table
c.execute('''
CREATE TABLE IF NOT EXISTS ProductCatalog (
    ProductID TEXT PRIMARY KEY,
    Name TEXT NOT NULL,
    Description TEXT,
    PriceUSD REAL NOT NULL,
    ComplianceDocID TEXT
)
''')

# 4. PricingRules Table
c.execute('''
CREATE TABLE IF NOT EXISTS PricingRules (
    RuleID INTEGER PRIMARY KEY AUTOINCREMENT,
    ProductID TEXT NOT NULL,
    MinQty INTEGER,
    MaxQty INTEGER,
    DiscountPct REAL,
    FOREIGN KEY (ProductID) REFERENCES ProductCatalog(ProductID)
)
''')

# 5. ComplianceDocs Table
c.execute('''
CREATE TABLE IF NOT EXISTS ComplianceDocs (
    ComplianceDocID TEXT PRIMARY KEY,
    Title TEXT NOT NULL,
    Version TEXT,
    Description TEXT
)
''')

# 6. ApprovalMatrix Table
c.execute('''
CREATE TABLE IF NOT EXISTS ApprovalMatrix (
    MatrixID INTEGER PRIMARY KEY AUTOINCREMENT,
    MinValueUSD REAL,
    MaxValueUSD REAL,
    ApproverRole TEXT
)
''')

# 7. QuoteTemplates Table
c.execute('''
CREATE TABLE IF NOT EXISTS QuoteTemplates (
    TemplateID TEXT PRIMARY KEY,
    Name TEXT NOT NULL,
    Format TEXT,
    Version TEXT,
    Description TEXT
)
''')

# Insert sample Roles
roles = [
    ('Salesperson',),
    ('Sales Manager',),
    ('Regional Director',),
    ('VP of Sales',),
    ('Admin',)
]
c.executemany('INSERT OR IGNORE INTO Roles (RoleName) VALUES (?)', roles)

# Insert sample Users
users = [
    ('Alice', 'alice@example.com', 1),
    ('Bob', 'bob@example.com', 2),
    ('Charlie', 'charlie@example.com', 3),
    ('Diana', 'diana@example.com', 4)
]
c.executemany('INSERT OR IGNORE INTO Users (UserName, Email, RoleID) VALUES (?, ?, ?)', users)

# Insert sample ProductCatalog
products = [
    ('P001', 'Laser Printer X100', 'High-speed office printer', 299.99, 'C001'),
    ('P002', 'Smart Scanner S200', 'Compact document scanner', 149.99, 'C002'),
    ('P003', 'Cloud Storage Pro', '1TB secure cloud storage', 99.99, 'C003')
]
c.executemany('INSERT OR IGNORE INTO ProductCatalog (ProductID, Name, Description, PriceUSD, ComplianceDocID) VALUES (?, ?, ?, ?, ?)', products)

# Insert sample PricingRules
pricing_rules = [
    ("P001", 1, 10, 0),
    ("P001", 11, 50, 5),
    ("P002", 1, 20, 0),
    ("P002", 21, 100, 10)
]
c.executemany('INSERT OR IGNORE INTO PricingRules (ProductID, MinQty, MaxQty, DiscountPct) VALUES (?, ?, ?, ?)', pricing_rules)

# Insert sample ComplianceDocs
compliance_docs = [
    ('C001', 'Printer Safety Cert', 'v1.0', 'Meets EU/US standards'),
    ('C002', 'Scanner EMC Cert', 'v2.1', 'Electromagnetic compliance'),
    ('C003', 'Cloud Data Policy', 'v3.2', 'GDPR compliant')
]
c.executemany('INSERT OR IGNORE INTO ComplianceDocs (ComplianceDocID, Title, Version, Description) VALUES (?, ?, ?, ?)', compliance_docs)

# Insert sample ApprovalMatrix
approval_matrix = [
    (0, 10000, 'Sales Manager'),
    (10001, 50000, 'Regional Director'),
    (50001, 1000000, 'VP of Sales')
]
c.executemany('INSERT OR IGNORE INTO ApprovalMatrix (MinValueUSD, MaxValueUSD, ApproverRole) VALUES (?, ?, ?)', approval_matrix)

# Insert sample QuoteTemplates
templates = [
    ('T001', 'Standard LG B2B Template', 'DOCX/PDF', 'v2.3', 'Includes branding, terms, SLA')
]
c.executemany('INSERT OR IGNORE INTO QuoteTemplates (TemplateID, Name, Format, Version, Description) VALUES (?, ?, ?, ?, ?)', templates)

conn.commit()
conn.close()

print("Sample SQLite DB 'sales_quotation_ai_mock.db' created and populated successfully.")
