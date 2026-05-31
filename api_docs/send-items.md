# Send Items

## Endpoints

- Estonia: https://aktiva.merit.ee/api/v2/senditems
- Poland: https://program.360ksiegowosc.pl/api/v2/senditems

## Query Payload

| Field | Type | Comment |
|---|---|---|
| Items | Array of Item objects | |

### Item Object

| Field | Type | Comment |
|---|---|---|
| Type | Int | Required. `1` stock item, `2` service, `3` item |
| Usage | Int | Required. `1` sales, `2` purchases, `3` sales and purchases |
| Code | Str | Max length 20; required |
| EANCode | Str | |
| Description | Str | Max length 100; required |
| UOMName | Str | Max length 64; required for stock item |
| DefLocationCode | Str | Max length 20; required if multiple stocks are in use |
| GTUCode | Int | Poland only; for usage `1` (sales), allowed values `1...13` |
| DescriptionEN | Str | Max length 100 |
| DescriptionRU | Str | Max length 100 |
| DescriptionFI | Str | Max length 100 |
| TaxId | Guid | |
| ItemGrCode | Str | If used, must exist in the company database |
| SalesAccCode | Str | Max length 10 |
| PurchaseAccCode | Str | Max length 10 |
| InventoryAccCode | Str | Max length 10 |
| CostAccCode | Str | Max length 10 |

## Successful Result

Array of answers.

| Field | Type | Comment |
|---|---|---|
| ItemId | Guid | |
| Code | Str | |
