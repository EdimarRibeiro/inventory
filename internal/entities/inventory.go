package entities

import (
	"errors"
	"time"
)

type Inventory struct {
	Id            uint64      `gorm:"primaryKey" json:"id"`
	TenantId      uint64      `json:"tenantId"`
	Tenant        Tenant      `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"-"`
	Name          string      `gorm:"size:150" json:"name"`
	ParticipantId uint64      `json:"participantId"`
	Participant   Participant `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"participant"`
	StartDate     time.Time   `gorm:"datetime" json:"startDate"`
	EndDate       time.Time   `gorm:"datetime" json:"endDate"`
	Processed     bool        `json:"processed"`
	Cloused       bool        `json:"cloused"`
}

func (c *Inventory) Validate() error {
	if c.Name == "" {
		return errors.New("the name is required")
	}
	if c.TenantId == 0 {
		return errors.New("the tenantId is invalid")
	}
	if c.ParticipantId == 0 {
		return errors.New("the participantId is invalid")
	}
	return nil
}

func (c *Inventory) SetProcessed() error {
	c.Processed = true
	return nil
}

func NewInventory(tenantId uint64, participantId uint64, name string) (*Inventory, error) {
	model := Inventory{
		Id:            0,
		Name:          name,
		TenantId:      tenantId,
		ParticipantId: participantId,
		StartDate:     time.Now(),
		Processed:     false,
	}
	return NewInventoryEntity(model)
}

func NewInventoryEntity(entity Inventory) (*Inventory, error) {
	err := entity.Validate()
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
