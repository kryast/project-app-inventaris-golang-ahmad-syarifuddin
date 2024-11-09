package service

import (
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
	"project-app-inventaris-golang-ahmad-syarifuddin/repository"
)

type AdminService struct {
	adminRepo repository.AdminRepositoryDB
}

func NewAdminService(repo repository.AdminRepositoryDB) AdminService {
	return AdminService{adminRepo: repo}
}

func (s *AdminService) ValidateAdmin(username, password string) (model.Admin, error) {
	admin, err := s.adminRepo.GetAdminByUsername(username)
	if err != nil {
		return model.Admin{}, fmt.Errorf("failed to get admin: %v", err)
	}

	return admin, nil
}
