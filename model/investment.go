package model

type Investment struct {
	ID               int     `json:"id"`
	ItemID           int     `json:"item_id"`
	Name             string  `json:"name"`
	InitialPrice     int     `json:"initial_price"`
	DepreciatedValue int     `json:"depreciated_value"`
	DepreciationRate float64 `json:"depreciation_rate"`
}
