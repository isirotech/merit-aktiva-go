package merit

import (
	"fmt"
)

type UnitOfMeasure struct {
	Code      string `json:"Code"`
	Name      string `json:"Name"`
	NonActive bool   `json:"NonActive"`
}

func (c *Client) GetUnitsOfMeasure() ([]UnitOfMeasure, error) {
	units := []UnitOfMeasure{}
	err := c.post(epGetUnitOfMeasureList, nil, &units)
	if err != nil {
		return nil, fmt.Errorf("failed to get units of measure: %w", err)
	}
	return units, nil
}
