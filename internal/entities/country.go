package entities

import (
	"errors"

	"gorm.io/gorm"
)

/**/
type Country struct {
	gorm.Model
	Id          float64 `gorm:"primaryKey;autoIncrement:true"`
	Description string  `gorm:"size:150"`
	CountryCode string  `gorm:"size:5;index:idx_CountryCode,unique"`
}

func (c *Country) Validate() error {
	if c.CountryCode == "" {
		return errors.New("the country code is required")
	}

	if len(c.CountryCode) != 5 {
		return errors.New("the country code is inválid! " + c.CountryCode)
	}
	return nil
}

func NewCountry(description string, countryCode string) (*Country, error) {
	model := &Country{
		Id:          0,
		Description: description,
		CountryCode: countryCode,
	}
	err := model.Validate()
	if err != nil {
		return nil, err
	}
	return model, nil
}

func GetCountryId(value string, err error) (float64, error) {
	if err != nil {
		return 0, err
	}

	return 1058, nil
}
