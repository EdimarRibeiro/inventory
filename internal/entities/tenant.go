package entities

import (
	"errors"
	"time"
)

type Tenant struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"size:150" json:"name"`
	Document  string    `gorm:"size:14" json:"document"`
	StartDate time.Time `gorm:"datetime" json:"startDate"`
	PersonId  uint64    `gorm:"default:null" json:"personId"`
	Canceled  bool      `json:"canceled"`
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
