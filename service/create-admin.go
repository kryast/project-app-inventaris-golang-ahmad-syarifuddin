package service

import "project-app-inventaris-golang-ahmad-syarifuddin/model"

func (s *AdminService) CreateAdmin(admin model.Admin) error {
	return s.adminRepo.CreateAdmin(admin)
}
