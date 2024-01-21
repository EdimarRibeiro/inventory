package entities

import (
	"errors"
)

/**/
type City struct {
	Id          uint64 `gorm:"primaryKey;autoIncrement:true"  json:"id"`
	Description string `gorm:"size:150" json:"description"`
	CityCode    string `gorm:"size:7;index:idx_CityCode,unique" json:"cityCode"`
}

func (c *City) Validate() error {
	if c.CityCode == "" {
		return errors.New("the city code is required")
	}

	if len(c.CityCode) != 7 {
		return errors.New("the city code is invalid! " + c.CityCode)
	}
	return nil
}

func NewCity(description string, cityCode string) (*City, error) {
	model := &City{
		Id:          0,
		Description: description,
		CityCode:    cityCode,
	}
	err := model.Validate()
	if err != nil {
		return nil, err
	}
	return model, nil
}
