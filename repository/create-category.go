package repository

import (
	"database/sql"
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

type CategoryRepositoryDB struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepositoryDB {
	return CategoryRepositoryDB{DB: db}
}

func (r *CategoryRepositoryDB) CreateDataCategory(category model.Category) error {
	query := "INSERT INTO categories (name , description) VALUES ($1 , $2) RETURNING id"

	err := r.DB.QueryRow(query, category.Name, category.Description).Scan(&category.ID)
	if err != nil {
		return fmt.Errorf("error at repository")
	}

	return nil
}
