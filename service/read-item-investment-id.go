package service

import (
	"fmt"
	"project-app-inventaris-golang-ahmad-syarifuddin/model"
)

func (is *ItemService) GetItemInvestmentByID(itemID int) (*model.Investment, error) {

	item, err := is.RepoItem.GetItemById(itemID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve item: %v", err)
	}

	depreciationRate := 10.0
	depreciatedValue := int(float64(item.Price) * (1 - (depreciationRate / 100)))

	itemDepreciation := model.Investment{
		ItemID:           item.ID,
		Name:             item.Name,
		InitialPrice:     item.Price,
		DepreciatedValue: depreciatedValue,
		DepreciationRate: depreciationRate,
	}

	return &itemDepreciation, nil
}
