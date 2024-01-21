package entities

import (
	"errors"
)

/**/
type Country struct {
	Id          uint64 `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Description string `gorm:"size:150" json:"description"`
	CountryCode string `gorm:"size:5;index:idx_CountryCode,unique" json:"countryCode"`
}

func (c *Country) Validate() error {
	if c.CountryCode == "" {
		return errors.New("the country code is required")
	}

	if len(c.CountryCode) != 5 {
		return errors.New("the country code is invalid! " + c.CountryCode)
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
