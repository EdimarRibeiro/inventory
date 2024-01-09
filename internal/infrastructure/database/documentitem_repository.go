package database

import (
	"errors"
	"strconv"

	"github.com/EdimarRibeiro/inventory/internal/entities"
	"github.com/EdimarRibeiro/inventory/internal/models"
	"gorm.io/gorm"
)

type DocumentItemRepository struct {
	DB *gorm.DB
}

func CreateDocumentItemRepository(db *gorm.DB) *DocumentItemRepository {
	return &DocumentItemRepository{DB: db}
}

func (entity *DocumentItemRepository) Save(model *entities.DocumentItem) (*entities.DocumentItem, error) {
	result := entity.DB.Save(&model)
	return model, result.Error
}

func (entity *DocumentItemRepository) Search(where string) ([]entities.DocumentItem, error) {
	var model []entities.DocumentItem
	result := entity.DB.Where(where).Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}

func (entity *DocumentItemRepository) SumaryQuantity(inventoryId uint64) (*[]models.SumaryQuantityModel, error) {
	var result *[]models.SumaryQuantityModel

	entity.DB.Raw("select i.InventoryId, i.ProductId, d.OperationId, p.OriginCode, p.UnitId "+
		" , Case d.OperationId when 0 then SUM(i.Quantity) else 0 end as SumaryQuantityInput "+
		" , Case d.OperationId when 1 then SUM(i.Quantity) else 0 end as SumaryQuantityOutput "+
		" from [dbo].[DocumentItem] i "+
		" join [dbo].[Product] p on p.Id = i.ProductId "+
		" join [dbo].[Document] d on d.Id = i.DocumentId "+
		" where i.InventoryId = ? "+
		" group by i.InventoryId, i.ProductId, d.OperationId, p.OriginCode, p.UnitId", inventoryId).Scan(&result)

	if len(*result) == 0 {
		return nil, errors.New("NotFound Document item the inventoryId =" + strconv.FormatUint(inventoryId, 10))
	}

	return result, nil
}
