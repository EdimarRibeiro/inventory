package entities

import (
	"errors"

	"github.com/EdimarRibeiro/inventory/internal/utils"
	"gorm.io/gorm"
)

/*0150*/
type Participant struct {
	gorm.Model
	Id           float64 `gorm:"primaryKey;autoIncrement:true"`
	TenantId     float64
	Tenant       Tenant `gorm:"constraint:OnUpdate:NULL,OnDelete:SET NULL;"`
	OriginCode   string `gorm:"size:60"`
	Name         string `gorm:"size:100"`
	Document     string `gorm:"size:14"`
	DocumentCpf  string `gorm:"size:11"`
	Registration string `gorm:"size:14"`
	CountryId    float64
	Country      Country `gorm:"constraint:OnUpdate:NULL,OnDelete:SET NULL;"`
	CityId       float64
	City         City   `gorm:"constraint:OnUpdate:NULL,OnDelete:SET NULL;"`
	Suframa      string `gorm:"size:9"`
	Street       string `gorm:"size:200"`
	Number       string `gorm:"size:10"`
	Complememt   string `gorm:"size:60"`
	Neighborhood string `gorm:"size:60"`
}

func (c *Participant) Validate() error {
	if c.OriginCode == "" {
		return errors.New("the originCode is required")
	}
	if c.Name == "" {
		return errors.New("the name is required")
	}
	if c.Document == "" && c.DocumentCpf == "" {
		return errors.New("the document or documentCpf is required")
	}
	if c.Document != "" && len(c.Document) != 14 {
		return errors.New("the document is inválid")
	}
	if c.DocumentCpf != "" && len(c.DocumentCpf) != 11 {
		return errors.New("the documentCpf is inválid")
	}
	if c.TenantId == 0 {
		return errors.New("the tenantId is inválid")
	}
	if c.CountryId == 0 {
		return errors.New("the countryId is inválid")
	}
	if c.CityId == 0 {
		return errors.New("the cityId is inválid")
	}
	return nil
}

func NewParticipant(tenantId float64, originCode string, name string, document string, registration string, countryId float64, cityId float64, street string, number string, complement string, neighborhood string) (*Participant, error) {
	model := Participant{
		Id:           0,
		OriginCode:   originCode,
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
	}
	return NewParticipantEntity(model)
}

func CreateParticipant(tenantId float64, line string) (*Participant, error) {
	var err error = nil
	participant := Participant{}
	participant.Id, err = 0, nil
	participant.TenantId, err = tenantId, nil
	participant.OriginCode, err = utils.CopyText(line, 2)
	if err != nil {
		return nil, err
	}
	participant.Name, err = utils.CopyText(line, 3)
	if err != nil {
		return nil, err
	}
	participant.CountryId, err = GetCountryId(utils.CopyText(line, 4))
	if err != nil {
		return nil, err
	}
	participant.Document, err = utils.CopyText(line, 5)
	if err != nil {
		return nil, err
	}
	participant.DocumentCpf, err = utils.CopyText(line, 6)
	if err != nil {
		return nil, err
	}
	participant.Registration, err = utils.CopyText(line, 7)
	if err != nil {
		return nil, err
	}
	participant.CityId, err = GetCityId(utils.CopyText(line, 8))
	if err != nil {
		return nil, err
	}
	participant.Suframa, err = utils.CopyText(line, 9)
	if err != nil {
		return nil, err
	}
	participant.Street, err = utils.CopyText(line, 10)
	if err != nil {
		return nil, err
	}
	participant.Number, err = utils.CopyText(line, 11)
	if err != nil {
		return nil, err
	}
	participant.Complememt, err = utils.CopyText(line, 12)
	if err != nil {
		return nil, err
	}
	participant.Neighborhood, err = utils.CopyText(line, 13)
	if err != nil {
		return nil, err
	}

	return NewParticipantEntity(participant)
}

func NewParticipantEntity(entity Participant) (*Participant, error) {
	err := entity.Validate()
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func GetParticipantId(value string, err error) (float64, error) {
	if err != nil {
		return 0, err
	}

	return 0, nil
}
