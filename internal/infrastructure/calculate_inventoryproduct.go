package infrastructure

import (
	"strconv"
	"time"

	entities "github.com/EdimarRibeiro/inventory/internal/entities"
	entitiesinterface "github.com/EdimarRibeiro/inventory/internal/interfaces/entities"
)

type CalculateBalanceQuantityData struct {
	InventoryProduct entitiesinterface.InventoryProductRepositoryInterface
	DocumentItem     entitiesinterface.DocumentItemRepositoryInterface
}

func CreateCalculateBalanceQuantityData(inventoryProduct entitiesinterface.InventoryProductRepositoryInterface, documentItem entitiesinterface.DocumentItemRepositoryInterface) *CalculateBalanceQuantityData {
	return &CalculateBalanceQuantityData{InventoryProduct: inventoryProduct, DocumentItem: documentItem}
}

func (calc *CalculateBalanceQuantityData) Execute(inventoryId uint64) error {
	sum := CreateSumaryQuantityDocumentItem(calc.DocumentItem)

	go calculeted(sum, calc, inventoryId)

	return nil
}

func calculeted(sum *SumaryQuantityDocumentItem, calc *CalculateBalanceQuantityData, inventoryId uint64) error {
	sumaries, err := sum.Execute(inventoryId)
	if err != nil {
		return err
	}

	for _, sumary := range *sumaries {
		invs, err := calc.InventoryProduct.Search("InventoryId=" + strconv.FormatUint(sumary.InventoryId, 10) + " and ProductId=" + strconv.FormatUint(sumary.ProductId, 10))

		if err != nil {
			return err
		}
		var invProd *entities.InventoryProduct
		if len(invs) > 0 {
			invProd = &invs[0]
		} else {
			invProd, err = entities.NewInventoryProduct(sumary.InventoryId, sumary.ProductId, sumary.OriginCode, time.Now(), sumary.UnitId, 0, 0, 0, "0", 0, "", "", 0)

			if err != nil {
				return err
			}
		}
		err = invProd.SetInputQuantity(sumary.SumaryQuantityInput)
		if err != nil {
			return err
		}
		err = invProd.SetOutputQuantity(sumary.SumaryQuantityOutput)
		if err != nil {
			return err
		}
		err = invProd.CalculateBalanceQuantity()
		if err != nil {
			return err
		}
		_, err = calc.InventoryProduct.Save(invProd)
		if err != nil {
			return err
		}
	}
	return nil
}
