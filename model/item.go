package model

import "time"

type Item struct {
	ID                  int       `json:"id"`
	Name                string    `json:"name"`
	CategoryID          int       `json:"category_id"`
	PhotoURL            string    `json:"photo_url"`
	Price               int       `json:"price"`
	PurchaseDate        time.Time `json:"purchase_date"`
	TotalUsageDay       int       `json:"total_usage_day"`
	ReplacementRequired bool      `json:"replacement_required"`
}
