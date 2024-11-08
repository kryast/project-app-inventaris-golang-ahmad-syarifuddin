package model

import "time"

type Item struct {
	ID            int
	Name          string
	Category      int
	PhotoURL      string
	Price         int
	PurchaseDate  time.Time
	TotalUsageDay int
}
