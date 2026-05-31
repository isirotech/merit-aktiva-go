# Get Invoices

## Endpoints

### v1 (`getinvoices`)
- Estonia: https://aktiva.merit.ee/api/v1/getinvoices
- Poland: https://program.360ksiegowosc.pl/api/v1/getinvoices

### v2 (`getinvoices`) - Multiple Dimensions
- Estonia: https://aktiva.merit.ee/api/v2/getinvoices
- Poland: https://program.360ksiegowosc.pl/api/v2/getinvoices

### v2 (`getinvoices2`) - By Invoice Number or Customer
- Estonia: https://aktiva.merit.ee/api/v2/getinvoices2
- Poland: https://program.360ksiegopwosc.pl/api/v2/getinvoices2

## Query Payload v1

Bare minimum example when querying list of invoices:

```json
{
  "PeriodStart": "YYYYmmdd",
  "PeriodEnd": "YYYYmmdd",
  "UnPaid": true
}
```

| Field | Type | Comment |
|---|---|---|
| PeriodStart | Date | Format `yyyymmdd` |
| PeriodEnd | Date | Format `yyyymmdd`; period length max 3 months |
| UnPaid | Bool | |
| DateType | Int | |

## Query Payload v2

| Field | Type | Comment |
|---|---|---|
| PeriodStart | Date | |
| PeriodEnd | Date | |
| UnPaid | Bool | |
| BankId | Guid | |
| DateType | Int | `0` = DocumentDate, `1` = ChangedDate |

## Query Payload (`getinvoices2`)

| Field | Type | Comment |
|---|---|---|
| InvNo | Str | |
| CustName | Str | |
| CustId | Guid | |

## Successful Result v1

Response contains invoice header summary fields:

| Field | Type | Comment |
|---|---|---|
| SIHId | Guid | |
| DepartmentCode | Str | |
| DepartmentName | Str | |
| ProjectCode | Str | |
| ProjectName | Str | |
| BatchInfo | Str | |
| InvoiceNo | Str | |
| DocumentDate | Date | |
| TransactionDate | Date | |
| CustomerId | Guid | |
| CustomerName | Str | |
| CustomerRegNo | Str | |
| HComment | Str | |
| FComment | Str | |
| DueDate | Date | |
| CurrencyCode | Str | |
| CurrencyRate | Dec | |
| TaxAmount | Dec | |
| RoundingAmount | Dec | |
| TotalAmount | Dec | |
| ProfitAmount | Dec | |
| TotalSum | Dec | |
| UserName | Str | |
| ReferenceNo | Str | |
| PriceInclVat | Bool | |
| VatRegNo | Str | |
| PaidAmount | Dec | |
| EInvSent | Bool | |
| EmailSent | Date | |
| Paid | Bool | |

## Successful Result v2 and `getinvoices2`

| Field | Type | Comment |
|---|---|---|
| SIHId | Guid | |
| DepartmentCode | Str | |
| DepartmentName | Str | |
| Dimension1Code | Str | |
| Dimension2Code | Str | |
| Dimension3Code | Str | |
| Dimension4Code | Str | |
| Dimension5Code | Str | |
| Dimension6Code | Str | |
| Dimension7Code | Str | |
| AccountingDocAccDocBatchInfo | Str | |
| InvoiceNo | Str | |
| DocumentDate | Date | |
| TransactionDate | Date | |
| CustomerName | Str | |
| CustomerRegNo | Str | |
| CustomerId | Guid | |
| HComment | Str | |
| FComment | Str | |
| DueDate | Date | |
| CurrencyCode | Str | |
| CurrencyRate | Dec | |
| TaxAmount | Dec | |
| RoundingAmount | Dec | |
| TotalAmount | Dec | |
| ProfitAmount | Dec | |
| TotalSum | Dec | |
| UserName | Str | |
| ReferenceNo | Str | |
| PriceInclVat | Bool | |
| VatRegNo | Str | |
| PaidAmount | Dec | |
| EInvSent | Bool | |
| EmailSent | Date | |
| Paid | Bool | |
| ChangedDate | Date | |
