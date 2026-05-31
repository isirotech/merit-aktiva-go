# Get Invoice Details

## Endpoints

### v1
- Estonia: https://aktiva.merit.ee/api/v1/getinvoice
- Poland: https://program.360ksiegowosc.pl/api/v1/getinvoice

### v2
- Estonia: https://aktiva.merit.ee/api/v2/getinvoice
- Poland: https://program.360ksiegowosc.pl/api/v2/getinvoice

## Query Payload

```json
{
  "Id": "xxxxxxxx-xxxx-Mxxx-Nxxx-xxxxxxxxxxxx",
  "AddAttachment": true
}
```

## Successful Result v1

Top-level response fields:

| Field Name | Type | Comment |
|---|---|---|
| Header | SalesHeader object | |
| Lines | Array of InvoiceRow objects | |
| Payments | Array of Payment objects | |
| Allocations | Array of CostAllocation objects | Project and Cost Center allocations (if existing) |

### SalesHeader Object

| Field Name | Type | Comment |
|---|---|---|
| SIHId | Guid | |
| DepartmentCode | Str | |
| DepartmentName | Str | |
| ProjectCode | Str | |
| ProjectName | Str | |
| BatchInfo | Str | GL transaction code and number |
| InvoiceNo | Str | |
| DocumentDate | Date | |
| TransactionDate | Date | |
| CustomerName | Str | |
| HComment | Str | |
| FComment | Str | |
| DueDate | Date | |
| CurrencyCode | Str | |
| CurrencyRate | Dec | |
| TaxAmount | Dec | VAT amount |
| RoundingAmount | Dec | |
| TotalAmount | Dec | Amount without VAT |
| ProfitAmount | Dec | Margin amount |
| TotalSum | Dec | Total amount with taxes and rounding |
| UserName | Str | |
| ReferenceNo | Str | |
| PriceInclVat | Bool | true/false |
| VatRegNo | Str | |
| OfferId | Guid | |
| OfferDocType | Int | |
| OfferNo | Str | |
| FileName | Str | |
| FileContent | Str | File in base64 |
| PerSHId | Guid | |
| ContractNo | Str | |

### InvoiceRow Object

| Field Name | Type | Comment |
|---|---|---|
| ArticleCode | Str | |
| LocationCode | Str | |
| Quantity | Dec | |
| Price | Dec | |
| TaxName | Str | |
| TaxPct | Dec | |
| AmountExclVat | Dec | |
| AmountInclVat | Dec | |
| VatAmount | Dec | |
| AccountCode | Str | |
| DepartmentCode | Str | |
| DepartmentName | Str | |
| ItemCostAmount | Dec | |
| ProfitAmount | Dec | |
| DiscountPct | Dec | |
| DiscountAmount | Dec | |
| Description | Str | |
| UOMName | Str | |
| FixAsset | Bool | true/false |
| ProjectAllocation | Array of CostAllocation objects | |
| CostCenterAllocation | Array of CostAllocation objects | |

### Payment Object

| Field Name | Type | Comment |
|---|---|---|
| PaymDate | Date | |
| Amount | Dec | |
| PaymentMethod | Str | |
| PaymentId | Guid | |

### CostAllocation Object

| Field Name | Type | Comment |
|---|---|---|
| Code | Str | Code of project or cost center |
| AllocPct | Dec | Allocation percentage |
| AllocAmount | Dec | Allocation amount |

## Successful Result v2

Reference: https://api.merit.ee/connecting-robots/reference-manual/sales-invoices/get-invoice-details/

Top-level response fields:

| Field | Type | Comment |
|---|---|---|
| Header | Header object | |
| Lines | Array of InvoiceRow objects | |
| Payments | Array of Payment objects | |
| Attachment | Attachment object | |

### Invoice Header Object

| Field | Type | Comment |
|---|---|---|
| Dimensions | Array of Dimensions objects | |
| SIHId | Guid | |
| DepartmentCode | Str | |
| DepartmentName | Str | |
| ProjectCode | Str | |
| ProjectName | Str | |
| AccountingDockBatchInfo | Str | |
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
| PriceInclVat | Bool | true/false |
| VatRegNo | Str | |
| PaidAmount | Dec | |
| EInvSent | Bool | true/false |
| EmailSent | Date | |
| EInvOperator | Int | 1 not exist, 2 Omniva bank e-invoice, 3 bank full extent, 4 bank limited extent |
| OfferId | Guid | |
| OfferDocType | Int | 1 quote, 2 sales order, 3 prepayment invoice |
| OfferNo | Str | |
| FileExists | Bool | true/false |
| PerSHId | Guid | |
| ContractNo | Str | |
| Paid | Bool | |
| Contact | Str | |

### Dimensions Object

| Field | Type | Comment |
|---|---|---|
| DimId | Str | |
| DimValueId | Guid | |
| DimCode | Str | |

### InvoiceRow Object

| Field | Type | Comment |
|---|---|---|
| SILId | Guid | |
| ArticleCode | Str | |
| LocationCode | Str | |
| Quantity | Dec | |
| Price | Dec | |
| TaxId | Guid | |
| TaxName | Str | |
| TaxPct | Dec | |
| AmountExclVat | Dec | |
| AmountInclVat | Dec | |
| VatAmount | Dec | |
| AccountCode | Str | |
| DepartmentCode | Str | |
| DepartmentName | Str | |
| ItemCostAmount | Dec | |
| ProfitAmount | Dec | |
| DiscountPct | Dec | |
| DiscountAmount | Dec | |
| Description | Str | |
| UOMName | Str | |
| FixAsset | Bool | true/false |
| DimAllocation | Array of DimensionsAllocation objects | |
| StorLifeDate | Date | |
| BatchNo | Str | |

### DimensionsAllocation Object

| Field | Type | Comment |
|---|---|---|
| DimId | Str | |
| Code | Str | |
| AllocPct | Dec | |
| AllocAmount | Dec | |

### Payment Object

| Field | Type | Comment |
|---|---|---|
| PaymDate | Date | |
| Amount | Dec | |
| PaymentMethod | Str | |
| PaymentId | Guid | |

### Attachment Object

| Field | Type | Comment |
|---|---|---|
| Filename | Str | |
| FileContent | Str | File in base64 |
