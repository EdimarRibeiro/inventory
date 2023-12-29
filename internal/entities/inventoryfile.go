package entities

import (
	"errors"
)

type InventoryFile struct {
	Id          uint64 `gorm:"primaryKey"`
	InventoryId uint64
	Inventory   Inventory `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	FileName    string    `gorm:"size:500"`
	FileType    string    `gorm:"size:3"`
	Processed   bool
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

func NewInventoryFile(id uint64, inventoryId uint64, fileName string, fileType string) (*InventoryFile, error) {
	model := InventoryFile{
		Id:          id,
		InventoryId: inventoryId,
		FileName:    fileName,
		FileType:    fileType,
		Processed:   false,
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

func (i *InventoryFile) SetProcessed() error {
	i.Processed = true
	return nil
}
