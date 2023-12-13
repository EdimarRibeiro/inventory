package infrastructure

import (
	entities "github.com/EdimarRibeiro/inventory/internal/entities"
	entitiesinterface "github.com/EdimarRibeiro/inventory/internal/interfaces/entities"
)

type CalculateBalanceQuantity struct {
	InventoryProduct entitiesinterface.InventoryProductRepositoryInterface
}

func (calc *CalculateBalanceQuantity) Execute(input entities.InventoryProduct) (*entities.InventoryProduct, error) {
	var sum SumaryQuantityDocumentItem

	err := input.Validate()
	if err != nil {
		return nil, err
	}
	sumary, err := sum.Execute(input.InventoryId)
	if err != nil {
		return nil, err
	}
	err = input.SetInputQuantity(sumary.SumaryQuantityInput)
	if err != nil {
		return nil, err
	}
	err = input.SetOutputQuantity(sumary.SumaryQuantityOutput)
	if err != nil {
		return nil, err
	}
	input, err = calc.InventoryProduct.Save(&input)
	if err != nil {
		return nil, err
	}
	return &input, nil
}
