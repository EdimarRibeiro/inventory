package entities

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model
	Id        float64 `gorm:"primaryKey;autoIncrement:true"`
	Name      string  `gorm:"size:150"`
	Document  string  `gorm:"size:14"`
	PersonId  float64
	StartDate time.Time
	EndDate   time.Time
}

func (c *Tenant) Validate() error {
	if c.Name == "" {
		return errors.New("the name is required")
	}
	if c.Document == "" {
		return errors.New("the document is required")
	}
	if c.PersonId == 0 {
		return errors.New("the personId is inválid")
	}
	return nil
}

func NewTenant(name string, document string, personId float64) (*Tenant, error) {
	model := Tenant{
		Id:        0,
		Name:      name,
		Document:  document,
		PersonId:  personId,
		StartDate: time.Now(),
	}
	return NewTenantEntity(model)
}

func NewTenantEntity(entity Tenant) (*Tenant, error) {
	err := entity.Validate()
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
