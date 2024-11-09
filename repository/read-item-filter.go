package repository

// import (
// 	"fmt"
// 	"project-app-inventaris-golang-ahmad-syarifuddin/model"
// )

// func (ir *ItemRepositoryDB) GetFilteredItems(filter model.ItemFilter) ([]model.Item, int, error) {
// 	query := `SELECT items.id, items.name, items.category_id, items.photo_url, items.price, items.purchase_date, items.total_usage_days
// 			  FROM items
// 			  INNER JOIN categories ON items.category_id = categories.id`
// 	args := []interface{}{}
// 	argIdx := 1

// 	// Jika ada filter berdasarkan nama kategori (category)
// 	if filter.Category != "" {
// 		// Gabungkan dengan query kategori berdasarkan nama
// 		query += fmt.Sprintf(" WHERE categories.name ILIKE $%d", argIdx)
// 		args = append(args, "%"+filter.Category+"%") // Gunakan ILIKE untuk pencarian case-insensitive
// 		argIdx++
// 	} else {
// 		// Jika tidak ada filter kategori, kita tetap memulai query
// 		query += " WHERE 1=1"
// 	}

// 	// Filter berdasarkan search (misalnya nama item)
// 	if filter.Search != "" {
// 		query += fmt.Sprintf(" AND items.name ILIKE $%d", argIdx)
// 		args = append(args, "%"+filter.Search+"%")
// 		argIdx++
// 	}

// 	// Filter berdasarkan total_usage_days
// 	if filter.TotalUsageDays != "" {
// 		query += fmt.Sprintf(" AND items.total_usage_day >= $%d", argIdx)
// 		args = append(args, filter.TotalUsageDays)
// 		argIdx++
// 	}

// 	// Pagination: LIMIT dan OFFSET
// 	offset := (filter.Page - 1) * filter.Limit
// 	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIdx, argIdx+1)
// 	args = append(args, filter.Limit, offset)

// 	// Query untuk menghitung total items
// 	countQuery := `SELECT COUNT(*) FROM items
// 				   INNER JOIN categories ON items.category_id = categories.id`

// 	// Jika ada filter kategori, tambahkan ke query
// 	if filter.Category != "" {
// 		countQuery += fmt.Sprintf(" WHERE categories.name ILIKE $%d", argIdx-1)
// 	}

// 	// Eksekusi query untuk mengambil data items
// 	rows, err := ir.DB.Query(query, args...)
// 	if err != nil {
// 		return nil, 0, fmt.Errorf("failed to query items: %v", err)
// 	}
// 	defer rows.Close()

// 	var items []model.Item
// 	for rows.Next() {
// 		var item model.Item
// 		if err := rows.Scan(&item.ID, &item.Name, &item.CategoryID, &item.PhotoURL, &item.Price, &item.PurchaseDate, &item.TotalUsageDay); err != nil {
// 			return nil, 0, fmt.Errorf("error scanning row: %v", err)
// 		}
// 		items = append(items, item)
// 	}

// 	// Menghitung total items
// 	var totalItems int
// 	if filter.Category != "" {
// 		err = ir.DB.QueryRow(countQuery, "%"+filter.Category+"%").Scan(&totalItems)
// 	} else {
// 		err = ir.DB.QueryRow(countQuery).Scan(&totalItems)
// 	}

// 	if err != nil {
// 		return nil, 0, fmt.Errorf("failed to count items: %v", err)
// 	}

// 	return items, totalItems, nil
//
