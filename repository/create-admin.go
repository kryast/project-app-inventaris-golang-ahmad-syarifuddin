package repository

import (
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

func (r *AdminRepositoryDB) CreateAdmin(admin model.Admin) error {
	query := `INSERT INTO admins (name, username, password) VALUES ($1, $2, $3)`
	_, err := r.DB.Exec(query, admin.Name, admin.Username, admin.Password)
	if err != nil {
		return fmt.Errorf("failed to insert admin: %v", err)
	}
	return nil
}
