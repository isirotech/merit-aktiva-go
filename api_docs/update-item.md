# Update Item

## Source

- Private: Integrations: https://api.merit.ee/?page_id=6
- AKTIVA Reference Manual: https://api.merit.ee/connecting-robots/reference-manual/
- Items: https://api.merit.ee/connecting-robots/reference-manual/items/
- Update Item: https://api.merit.ee/connecting-robots/reference-manual/items/update-item/

## Endpoints

- Estonia: https://aktiva.merit.ee/api/v1/updateitem
- Poland: https://program.360ksiegowosc.pl/api/v1/updateitem

## Query Payload

| Field | Type | Comment |
|---|---|---|
| Id | Guid | Required |
| Code | Str | Max length 20 |
| Description | Str | Max length 100 |
| SalesPrice | Dec | Decimal 18.2 |
| ItemGrCode | Str | If used, must exist in the company database |
| DiscountPct | Dec | Decimal 18.2 |
| EANCode | Str | Max length 13 |
| NameEN | Str | Max length 100 |
| LastPurchasePrice | Dec | Decimal 18.2 |
| SalesAccountCode | Str | Max length 10; if used, must exist in the company database |
| InventoryAccountCode | Str | Max length 10; if used, must exist in the company database |
| ItemCostAccountCode | Str | Max length 10; if used, must exist in the company database |
| TaxId | Guid | |
| GTUCode | Int | Poland only; for usage `1` (sales), allowed values `1...13` |
