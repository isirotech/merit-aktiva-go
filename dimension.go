package merit

import (
	"github.com/Microsoft/go-winio/pkg/guid"
)

type DimensionsObject struct {
	DimID      int       `json:"DimId"`
	DimValueId guid.GUID `json:"DimValueId"`
	DimCode    string    `json:"DimCode"`
}

type Dimension struct {
	DimID   int       `json:"DimId"`
	DimName string    `json:"DimName"`
	ID      guid.GUID `json:"Id"`
	Code    string    `json:"Code"`
	Name    string    `json:"Name"`
	EndDate string    `json:"EndDate"`
}

func (c *Client) GetDimensions(query Dimension) ([]Dimension, error) {
	dimensions := []Dimension{}
	err := c.post(epGetListOfDimensions, struct{}{}, &dimensions)
	if err != nil {
		return nil, err
	}
	filteredDimensions := []Dimension{}
	for _, dimension := range dimensions {
		if query.DimID != 0 && dimension.DimID != query.DimID {
			continue
		}
		if query.DimName != "" && dimension.DimName != query.DimName {
			continue
		}
		if query.ID != guid.FromArray([16]byte{}) && dimension.ID != query.ID {
			continue
		}
		if query.Code != "" && dimension.Code != query.Code {
			continue
		}
		if query.Name != "" && dimension.Name != query.Name {
			continue
		}
		if query.EndDate != "" && dimension.EndDate != query.EndDate {
			continue
		}
		filteredDimensions = append(filteredDimensions, dimension)
	}
	return filteredDimensions, nil
}
