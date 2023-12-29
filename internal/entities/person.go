package entities

import (
	"errors"
)

type Person struct {
	Id           uint64 `gorm:"primaryKey;autoIncrement:true"`
	TenantId     uint64
	Tenant       Tenant `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	Name         string `gorm:"size:100"`
	Document     string `gorm:"size:14"`
	Registration string `gorm:"size:14"`
	CountryId    uint64
	Country      Country `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	CityId       uint64
	City         City   `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	Suframa      string `gorm:"size:9"`
	Street       string `gorm:"size:200"`
	Number       string `gorm:"size:10"`
	Complememt   string `gorm:"size:60"`
	Neighborhood string `gorm:"size:60"`
	ZipCode      string `gorm:"size:8"`
}

func (c *Person) Validate() error {
	if c.Name == "" {
		return errors.New("the name is required")
	}
	if c.Document == "" {
		return errors.New("the document or documentCpf is required")
	}
	if c.Document != "" && (len(c.Document) != 14 && len(c.Document) != 11) {
		return errors.New("the document is invalid person")
	}
	if c.TenantId == 0 {
		return errors.New("the tenantId is invalid")
	}
	if c.CountryId == 0 {
		return errors.New("the countryId is invalid")
	}
	if c.CityId == 0 {
		return errors.New("the cityId is invalid")
	}
	return nil
}

func NewPerson(tenantId uint64, name string, document string, registration string, countryId uint64, cityId uint64, street string, number string, complement string, neighborhood string, zipCode string) (*Person, error) {
	model := Person{
		TenantId:     tenantId,
		Name:         name,
		Document:     document,
		Registration: registration,
		CountryId:    countryId,
		CityId:       cityId,
		Street:       street,
		Number:       number,
		Neighborhood: neighborhood,
		Complememt:   complement,
		ZipCode:      zipCode,
	}
	return NewPersonEntity(model)
}

func NewPersonEntity(entity Person) (*Person, error) {
	err := entity.Validate()
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
