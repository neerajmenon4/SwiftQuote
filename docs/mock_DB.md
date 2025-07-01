
## 📦 Sample Use Case

> **Scenario**: A salesperson at LG wants to create a quote for a retail chain looking to buy commercial air conditioners, digital signage, and display panels.

---

## 🗃️ Mock Database Schema with Sample Data

### 1. `Products`

| ProductID | Name                          | Category          | Description                             | BasePrice (USD) | SKU       | Available |
| --------- | ----------------------------- | ----------------- | --------------------------------------- | --------------- | --------- | --------- |
| P001      | LG Commercial AC 5 Ton        | HVAC              | High-efficiency AC for commercial use   | 4,500           | LGCAC5T   | Yes       |
| P002      | LG 55" 4K UHD Digital Signage | Display Solutions | Ultra HD signage for commercial spaces  | 1,200           | LGDS55UHD | Yes       |
| P003      | LG OLED Video Wall Panel      | Display Solutions | Seamless wall display panel             | 3,000           | LGOLEDVW  | Yes       |
| P004      | LG Inverter Window AC         | HVAC              | Energy-efficient window air conditioner | 700             | LGIWACINV | Yes       |

---

### 2. `Customers`

| CustomerID | Name               | Industry      | Tier   | Region        |
| ---------- | ------------------ | ------------- | ------ | ------------- |
| C001       | ABC Retail Group   | Retail        | Gold   | North America |
| C002       | GlobalTech Offices | Corporate B2B | Silver | Europe        |

---

### 3. `PriceRules`

| CustomerID | ProductID | Discount (%) | FinalPrice (USD) |
| ---------- | --------- | ------------ | ---------------- |
| C001       | P001      | 10           | 4,050            |
| C001       | P002      | 5            | 1,140            |
| C002       | P003      | 7            | 2,790            |

---

### 4. `Quotations`

| QuoteID | CustomerID | DateCreated | Status   | TotalValue (USD) |
| ------- | ---------- | ----------- | -------- | ---------------- |
| Q1001   | C001       | 2025-06-20  | Draft    | 21,090           |
| Q1002   | C002       | 2025-06-22  | Reviewed | 16,740           |

---

### 5. `QuoteItems`

| QuoteID | ProductID | Quantity | UnitPrice (USD)  | LineTotal (USD) |
| ------- | --------- | -------- | ---------------- | --------------- |
| Q1001   | P001      | 3        | 4,050            | 12,150          |
| Q1001   | P002      | 5        | 1,140            | 5,700           |
| Q1001   | P004      | 5        | 648 (after 7.5%) | 3,240           |
| Q1002   | P003      | 6        | 2,790            | 16,740          |

---

### 6. `ApprovalMatrix`

| MinValue (USD) | MaxValue (USD) | ApproverRole      |
| -------------- | -------------- | ----------------- |
| 0              | 10,000         | Sales Manager     |
| 10,001         | 50,000         | Regional Director |
| 50,001         | 1,000,000      | VP of Sales       |

---

### 7. `QuoteTemplates`

| TemplateID | Name                     | Format   | Version | Description                   |
| ---------- | ------------------------ | -------- | ------- | ----------------------------- |
| T001       | Standard LG B2B Template | DOCX/PDF | v2.3    | Includes branding, terms, SLA |

---

## 🛠️ How This Helps

With this dataset, you can:

* Simulate the end-to-end flow of creating a quotation.
* Test document lookups (e.g., product catalog filters).
* Validate pricing logic and approval flows.
* Generate mock PDF or web-based quote templates.
