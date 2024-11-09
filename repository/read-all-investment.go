package repository

import (
	"database/sql"
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

type InvestmentRepositoryDB struct {
	DB *sql.DB
}

func NewInvestmentRepository(db *sql.DB) InvestmentRepositoryDB {
	return InvestmentRepositoryDB{DB: db}
}

func (r *InvestmentRepositoryDB) GetAllInvestments() ([]model.Investment, error) {

	query := `
		SELECT id, name, price, purchase_date
		FROM items`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve items: %v", err)
	}
	defer rows.Close()

	var investments []model.Investment

	depreciationRate := 10.0

	for rows.Next() {
		var item model.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Price, &item.PurchaseDate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan item: %v", err)
		}

		depreciatedValue := float64(item.Price) * (1 - (depreciationRate / 100))

		investment := model.Investment{
			ItemID:           item.ID,
			Name:             item.Name,
			InitialPrice:     item.Price,
			DepreciatedValue: int(depreciatedValue),
			DepreciationRate: depreciationRate,
		}

		investments = append(investments, investment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return investments, nil
}
