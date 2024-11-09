package service

import (
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

func (is *ItemService) GetAllInvestment() (*model.AllInvestment, error) {

	items, err := is.RepoItem.GetAllItem()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve items: %v", err)
	}

	var totalInvestment, totalDepreciatedValue float64

	depreciationRate := 10.0

	for _, item := range items {

		depreciatedValue := float64(item.Price) * (1 - (depreciationRate / 100))

		totalInvestment += float64(item.Price)
		totalDepreciatedValue += depreciatedValue
	}

	allInvestment := &model.AllInvestment{
		TotalInvestment:  int(totalInvestment),
		DepreciatedValue: int(totalDepreciatedValue),
	}

	return allInvestment, nil
}
