package merit

import (
	"fmt"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/shopspring/decimal"
)

type ItemType int

const (
	ItemTypeStock ItemType = iota + 1
	ItemTypeService
	ItemTypeItem
)

type ItemObject struct {
	Code            string   `json:"Code"`                      // Required
	Description     string   `json:"Description"`               // Required
	Type            ItemType `json:"Type,omitzero"`             // 1 = stock item, 2 = service, 3 = item. Required.
	UOMName         string   `json:"UOMName,omitempty"`         //
	DefLocationCode string   `json:"DefLocationCode,omitempty"` //
	EANCode         string   `json:"EANCode,omitempty"`         //
}

type Item struct {
	ID                   guid.GUID       `json:"ItemId"`
	Code                 string          `json:"Code"`
	Name                 string          `json:"Name"`
	UnitofMeasureName    string          `json:"UnitofMeasureName"`
	Type                 string          `json:"Type"`
	SalesPrice           decimal.Decimal `json:"SalesPrice"`
	InventoryQty         decimal.Decimal `json:"InventoryQty"`
	ReservedQty          decimal.Decimal `json:"ReservedQty"`
	VatTaxName           string          `json:"VatTaxName"`
	Usage                string          `json:"Usage"`
	SalesAccountCode     string          `json:"SalesAccountCode"`
	PurchaseAccountCode  string          `json:"PurchaseAccountCode"`
	InventoryAccountCode string          `json:"InventoryAccountCode"`
	ItemCostAccountCode  string          `json:"ItemCostAccountCode"`
	DiscountPct          decimal.Decimal `json:"DiscountPct"`
	LastPurchasePrice    decimal.Decimal `json:"LastPurchasePrice"`
	ItemUnitCost         decimal.Decimal `json:"ItemUnitCost"`
	InventoryCost        decimal.Decimal `json:"InventoryCost"`
	ItemGroupName        string          `json:"ItemGroupName"`
	DefLocName           string          `json:"DefLoc_Name"`
	EANCode              string          `json:"EANCode"`
}

type GetItemsQuery struct {
	ID           string `json:"Id,omitempty"`
	Code         string `json:"Code,omitempty"`
	Description  string `json:"Description,omitempty"`
	LocationCode string `json:"LocationCode,omitempty"`
}

func (c *Client) GetItems(query GetItemsQuery) ([]Item, error) {
	items := []Item{}
	err := c.post(epGetItems, query, &items)
	if err != nil {
		return nil, err
	}
	fmt.Println(items)
	return items, nil
}

type ItemGroup struct {
	Code string `json:"Code"`
	Name string `json:"Name"`
	ID   string `json:"Id"`
}

func (c *Client) GetItemGroups() ([]ItemGroup, error) {
	itemGroups := []ItemGroup{}
	err := c.post(epGetItemGroups, map[string]interface{}{}, &itemGroups)
	if err != nil {
		return nil, err
	}

	fmt.Println(itemGroups)
	return itemGroups, nil
}

// ItemUsage describes whether an item is used for sales, purchases, or both.
type ItemUsage int

const (
	ItemUsageSales             ItemUsage = 1
	ItemUsagePurchases         ItemUsage = 2
	ItemUsageSalesAndPurchases ItemUsage = 3
)

// SendItemObject is the payload for a single item in a SendItems request.
type SendItemObject struct {
	Type             ItemType  `json:"Type"`        // Required. 1=stock item, 2=service, 3=item.
	Usage            ItemUsage `json:"Usage"`       // Required. 1=sales, 2=purchases, 3=both.
	Code             string    `json:"Code"`        // Required. Max 20 chars.
	Description      string    `json:"Description"` // Required. Max 100 chars.
	EANCode          string    `json:"EANCode,omitempty"`
	UOMName          string    `json:"UOMName,omitempty"`         // Required for stock items.
	DefLocationCode  string    `json:"DefLocationCode,omitempty"` // Required if multiple stocks.
	DescriptionEN    string    `json:"DescriptionEN,omitempty"`
	DescriptionRU    string    `json:"DescriptionRU,omitempty"`
	DescriptionFI    string    `json:"DescriptionFI,omitempty"`
	TaxId            guid.GUID `json:"TaxId,omitempty"`
	ItemGrCode       string    `json:"ItemGrCode,omitempty"`
	SalesAccCode     string    `json:"SalesAccCode,omitempty"`
	PurchaseAccCode  string    `json:"PurchaseAccCode,omitempty"`
	InventoryAccCode string    `json:"InventoryAccCode,omitempty"`
	CostAccCode      string    `json:"CostAccCode,omitempty"`
}

// SendItemResult is the per-item response from the v2/senditems endpoint.
type SendItemResult struct {
	ItemId guid.GUID `json:"ItemId"`
	Code   string    `json:"Code"`
}

// SendItemsQuery is the request payload for the v2/senditems endpoint.
type SendItemsQuery struct {
	Items []SendItemObject `json:"Items"`
}

// SendItems creates or updates inventory items in Merit.
func (c *Client) SendItems(query SendItemsQuery) ([]SendItemResult, error) {
	results := []SendItemResult{}
	err := c.post(epSendItems, query, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
