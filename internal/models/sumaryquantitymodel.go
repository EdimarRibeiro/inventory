package models

type SumaryQuantityModel struct {
	InventoryId          uint64
	ProductId            uint64
	OriginCode           string
	UnitId               string
	SumaryQuantityInput  float64
	SumaryQuantityOutput float64
}

func NewSumaryQuantityModel(inventoryId uint64, productId uint64, originCode string, unitId string, sumaryQuantityInput float64, sumaryQuantityOutput float64) (*SumaryQuantityModel, error) {
	model := SumaryQuantityModel{
		InventoryId:          inventoryId,
		ProductId:            productId,
		OriginCode:           originCode,
		UnitId:               unitId,
		SumaryQuantityInput:  sumaryQuantityInput,
		SumaryQuantityOutput: sumaryQuantityOutput,
	}
	return NewSumaryQuantityModelBase(model)
}

func NewSumaryQuantityModelBase(entity SumaryQuantityModel) (*SumaryQuantityModel, error) {
	return &entity, nil
}

func LocateSumary(data []SumaryQuantityModel, productId uint64) *SumaryQuantityModel {
	for i := 0; i < len(data); i++ {
		if data[i].ProductId == productId {
			return &data[i]
		}
	}
	return nil
}
