package repository

import (
	"database/sql"
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

// InvestmentRepositoryDB adalah struktur untuk repository investasi
type InvestmentRepositoryDB struct {
	DB *sql.DB
}

func NewInvestmentRepository(db *sql.DB) InvestmentRepositoryDB {
	return InvestmentRepositoryDB{DB: db}
}

// GetAllInvestments menghitung investasi berdasarkan data item
func (r *InvestmentRepositoryDB) GetAllInvestments() ([]model.Investment, error) {
	// Query untuk mengambil semua data item dari database
	query := `
		SELECT id, name, price, purchase_date
		FROM items`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve items: %v", err)
	}
	defer rows.Close()

	// Slice untuk menampung semua investasi yang dihitung
	var investments []model.Investment

	// Asumsikan tarif depresiasi adalah 10% (misalnya)
	depreciationRate := 10.0

	// Looping untuk memindai setiap item dan menghitung investasi
	for rows.Next() {
		var item model.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Price, &item.PurchaseDate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan item: %v", err)
		}

		// Menghitung nilai depresiasi
		depreciatedValue := float64(item.Price) * (1 - (depreciationRate / 100))

		// Membuat struktur Investment berdasarkan data item
		investment := model.Investment{
			ItemID:           item.ID,
			Name:             item.Name,
			InitialPrice:     item.Price,
			DepreciatedValue: int(depreciatedValue),
			DepreciationRate: depreciationRate,
		}

		// Menambahkan investasi yang sudah dihitung ke dalam slice
		investments = append(investments, investment)
	}

	// Mengecek apakah ada error saat looping melalui rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	// Mengembalikan slice yang berisi semua investasi
	return investments, nil
}
