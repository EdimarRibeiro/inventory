package entities

import (
	"errors"

	"gorm.io/gorm"
)

type InventoryFile struct {
	gorm.Model
	Id          float64 `gorm:"primaryKey"`
	InventoryId float64
	Inventory   Inventory `gorm:"constraint:OnUpdate:NULL,OnDelete:SET NULL;"`
	FileName    string    `gorm:"size:500"`
}

func (c *InventoryFile) Validate() error {
	if c.InventoryId == 0 {
		return errors.New("the InventoryFile is required")
	}

	if c.FileName == "" {
		return errors.New("the value fileName is inválid! ")
	}

	return nil
}

func NewInventoryFile(id float64, inventoryId float64, fileName string) (*InventoryFile, error) {
	model := InventoryFile{
		Id:          id,
		InventoryId: inventoryId,
		FileName:    fileName,
	}
	return NewInventoryFileEntity(model)
}

func NewInventoryFileEntity(entity InventoryFile) (*InventoryFile, error) {
	err := entity.Validate()
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
