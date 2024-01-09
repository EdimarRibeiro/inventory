package entities

import (
	"errors"
	"time"
)

type Tenant struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement:true"`
	Name      string    `gorm:"size:150"`
	Document  string    `gorm:"size:14"`
	StartDate time.Time `gorm:"datetime"`
	PersonId  uint64    `gorm:"default:null"`
	Canceled  bool
}

func (c *Tenant) Validate() error {
	if c.Name == "" {
		return errors.New("the name is required")
	}
	if c.Document == "" {
		return errors.New("the document is required")
	}
	return nil
}

func NewTenant(name string, document string, personId uint64) (*Tenant, error) {
	model := Tenant{
		Name:      name,
		Document:  document,
		StartDate: time.Now(),
		Canceled:  false,
	}

	if personId > 0 {
		model.PersonId = personId
	}
	return NewTenantEntity(model)
}

func NewTenantEntity(entity Tenant) (*Tenant, error) {
	err := entity.Validate()
	if err != nil {
		return nil, err
	}
	//Null foreign key treatment
	resp := Tenant{
		Name:      entity.Name,
		Document:  entity.Document,
		StartDate: entity.StartDate,
		Canceled:  entity.Canceled,
	}
	if entity.PersonId > 0 {
		resp.PersonId = entity.PersonId
	}
	return &resp, nil
}
