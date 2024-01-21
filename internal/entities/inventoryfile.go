package entities

import (
	"errors"
)

type InventoryFile struct {
	Id          uint64    `gorm:"primaryKey" json:"id"`
	InventoryId uint64    `json:"inventoryId"`
	Inventory   Inventory `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"-"`
	FileName    string    `gorm:"size:500" json:"fileName"`
	FileType    string    `gorm:"size:3" json:"fileType"`
	Processed   bool      `json:"processed"`
}

func (c *InventoryFile) Validate() error {
	if c.InventoryId == 0 {
		return errors.New("the InventoryFile is required")
	}

	if c.FileName == "" {
		return errors.New("the value fileName is invalid! ")
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
