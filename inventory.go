package merit

import "time"

type Location struct {
	CompanyID     int    `json:"CompanyId"`
	LocationID    string `json:"LocationId"`
	Code          string `json:"Code"`
	Name          string `json:"Name"`
	InBPrefix     string `json:"InBPrefix"`
	InBNextNo     int    `json:"InBNextNo"`
	OutBPrefix    string `json:"OutBPrefix"`
	OutBNextNo    int    `json:"OutBNextNo"`
	Loc2LocPrefix string `json:"Loc2LocPrefix"`
	Loc2LocNextNo int    `json:"Loc2LocNextNo"`
	InvSetPrefix  string `json:"InvSetPrefix"`
	InvSetNextNo  int    `json:"InvSetNextNo"`
}

func (c *Client) GetListOfLocations() ([]Location, error) {
	locations := []Location{}
	err := c.post(epGetListOfLocations, map[string]interface{}{}, &locations)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

type Article struct {
	ItemCode         string  `json:"ItemCode"`
	EANCode          string  `json:"EANCode"`
	ItemName         string  `json:"ItemName"`
	LocName          string  `json:"LocName"`
	Quantity         float64 `json:"Quantity"`
	ReservedQuantity float64 `json:"ReservedQuantity"`
	UnitCode         string  `json:"UnitCode"`
	Amount           float64 `json:"Amount"`
	Price            float64 `json:"Price"`
}

type GetInventoryReportQuery struct {
	ArticleGroups    []string  `json:"ArticleGroups,omitempty"`
	Location         string    `json:"Location,omitempty"`
	RepDate          time.Time `json:"RepDate,omitempty"`
	ShowZero         bool      `json:"ShowZero,omitempty"`
	WithReservations bool      `json:"WithReservations,omitempty"`
}

type getInventoryReportQueryFormated struct {
	ArticleGroups    []string  `json:"ArticleGroups,omitempty"`
	Location         string    `json:"Location,omitempty"`
	RepDate          queryDate `json:"RepDate,omitempty"`
	ShowZero         bool      `json:"ShowZero,omitempty"`
	WithReservations bool      `json:"WithReservations,omitempty"`
}

func (c *Client) GetInventoryReport(query GetInventoryReportQuery) ([]Article, error) {
	queryFormated := getInventoryReportQueryFormated{
		ArticleGroups:    query.ArticleGroups,
		Location:         query.Location,
		RepDate:          queryDate{query.RepDate, "20060102"},
		ShowZero:         query.ShowZero,
		WithReservations: query.WithReservations,
	}
	articles := []Article{}
	err := c.post(epGetInventoryReport, queryFormated, &articles)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

type InventoryMovement struct {
	DocumentID     string             `json:"DocumentId"`
	DocDate        string             `json:"DocDate"`
	DocNo          string             `json:"DocNo"`
	Location1Code  string             `json:"Location1Code"`
	Location2Code  string             `json:"Location2Code"`
	DepartmentCode string             `json:"DepartmentCode"`
	Type           int                `json:"Type"`
	Comment        string             `json:"Comment"`
	Dimensions     []DimensionsObject `json:"Dimensions"`
	Rows           []Row              `json:"Rows"`
	ChangedDate    string             `json:"ChangedDate"`
}

type Row struct {
	HeaderID      string          `json:"HeaderId"`
	LineID        string          `json:"LineId"`
	ArticleCode   string          `json:"ArticleCode"`
	UOMName       string          `json:"UOMName"`
	Quantity      float64         `json:"Quantity"`
	Amount        float64         `json:"Amount"`
	Dimensions    []CostDimension `json:"Dimensions"`
	GLAccountCode string          `json:"GLAccountCode"`
}

type CostDimension struct {
	HeaderID    string  `json:"HeaderId"`
	LineID      string  `json:"LineId"`
	DimID       int     `json:"DimId"`
	Code        string  `json:"Code"`
	AllocPct    float64 `json:"AllocPct"`
	AllocAmount float64 `json:"AllocAmount"`
}

type GetInventoryMovementsQuery struct {
	PeriodStart time.Time `json:"PeriodStart,omitempty"`
	PeriodEnd   time.Time `json:"PeriodEnd,omitempty"`
	WithLines   bool      `json:"WithLines,omitempty"`
	ChangedDate int       `json:"ChangedDate,omitempty"`
}

type getInventoryMovementsQueryFormated struct {
	PeriodStart queryDate `json:"PeriodStart,omitempty"`
	PeriodEnd   queryDate `json:"PeriodEnd,omitempty"`
	WithLines   bool      `json:"WithLines,omitempty"`
	ChangedDate int       `json:"ChangedDate,omitempty"`
}

func (c *Client) GetInventoryMovements(query GetInventoryMovementsQuery) ([]InventoryMovement, error) {
	queryFormated := getInventoryMovementsQueryFormated{
		PeriodStart: queryDate{query.PeriodStart, "20060102"},
		PeriodEnd:   queryDate{query.PeriodEnd, "20060102"},
		WithLines:   query.WithLines,
		ChangedDate: query.ChangedDate,
	}
	inventoryMovements := []InventoryMovement{}
	err := c.post(epGetListOfInventoryMovements, queryFormated, &inventoryMovements)
	if err != nil {
		return nil, err
	}
	return inventoryMovements, nil
}
