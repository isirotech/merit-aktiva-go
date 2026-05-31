# Send Inventory Movements

- v1: Send Inventory Movements
- v2: Send Inventory Movements (recommended when dimensions are used)

Reference (v2): https://api.merit.ee/connecting-robots/reference-manual/inventory-movements1/send-invenroty-movements/

## Endpoints v1

- Estonia: https://aktiva.merit.ee/api/v1/SendInvMovement
- Poland: https://program.360ksiegowosc.pl/api/v1/SendInvMovement

## Query Payload v1

Example payload reference: https://api.merit.ee/connecting-robots/reference-manual/inventory-movements/send-invenroty-movements/example-send-inventori-movements/

| Field | Type | Comment |
|---|---|---|
| DocDate | Date | |
| DocNo | Str | Max length 35 |
| Location1Code | Str | Max length 20; stock code |
| Location2Code | Str | Max length 20; destination stock code for `Type = 3` |
| DepartmentCode | Str | Max length 20; if used, must exist in database |
| Type | Int | `1` in, `2` out, `3` between stocks |
| Rows | Array of Row objects | |

### Row Object (v1)

| Field | Type | Comment |
|---|---|---|
| ArticleCode | Str | Max length 20; must exist in database |
| UOMName | Str | Max length 64 |
| ItemUnitCost | Dec | Decimal 18.3; required for `Type = 1` |
| Quantity | Dec | Decimal 18.2 |

## Endpoints v2

- Estonia: https://aktiva.merit.ee/api/v2/SendInvMovement
- Poland: https://program.360ksiegowosc.pl/api/v2/SendInvMovement

## Query Payload v2

| Field | Type | Comment |
|---|---|---|
| DocDate | Date | |
| DocNo | Str | Max length 35 |
| Location1Code | Str | Max length 20; stock code |
| Location2Code | Str | Max length 20; destination stock code for `Type = 3` |
| DepartmentCode | Str | Max length 20; if used, must exist in database |
| Type | Int | `1` in, `2` out, `3` between stocks |
| Rows | Array of Row objects | |
| Dimensions | Array of Dimensions objects | |

### Row Object (v2)

| Field | Type | Comment |
|---|---|---|
| ArticleCode | Str | Max length 20; must exist in database |
| UOMName | Str | Max length 64 |
| ItemUnitCost | Dec | Decimal 18.3; required for `Type = 1` |
| Quantity | Dec | Decimal 18.2 |
| GLAccountCode | Str | Max length 10 |
| Dimensions | Array of Dimensions objects | |

### Dimensions Object

| Field | Type | Comment |
|---|---|---|
| DimId | Int | |
| DimValueId | Guid | |
| DimCode | Str | |
