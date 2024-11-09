package service

import (
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
	"project-app-inventaris-golang-ahmad-syarifuddin/repository"
)

type InvestmentService struct {
	RepoInvestment repository.InvestmentRepositoryDB
}

func NewInvestmentService(repo repository.InvestmentRepositoryDB) InvestmentService {
	return InvestmentService{RepoInvestment: repo}
}

func (is *InvestmentService) GetAllInvestment() ([]model.Investment, error) {
	items, err := is.RepoInvestment.GetAllInvestments()
	if err != nil {
		return nil, err
	}
	return items, nil
}
