package entities

import (
	"errors"

	"github.com/EdimarRibeiro/inventory/internal/utils"
)

/*0150*/
type Participant struct {
	Id           uint64 `gorm:"primaryKey;autoIncrement:true" json:"id"`
	TenantId     uint64 `json:"tenantId"`
	Tenant       Tenant `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"-"`
	OriginCode   string `gorm:"size:60" json:"originCode"`
	Name         string `gorm:"size:100" json:"name"`
	Document     string `gorm:"size:14" json:"document"`
	DocumentCpf  string `gorm:"size:11" json:"documentCpf"`
	Registration string `gorm:"size:14" json:"registration"`
	CountryCode  string `gorm:"size:5" json:"countryCode"`
	CityCode     string `gorm:"size:7" json:"cityCode"`
	Suframa      string `gorm:"size:9" json:"suframa"`
	Street       string `gorm:"size:200" json:"street"`
	Number       string `gorm:"size:10" json:"number"`
	Complememt   string `gorm:"size:60" json:"complememt"`
	Neighborhood string `gorm:"size:60" json:"neighborhood"`
}

func (c *Participant) Validate() error {
	if c.OriginCode == "" {
		return errors.New("the participant originCode is required")
	}
	if c.Name == "" {
		return errors.New("the name is required")
	}
	if c.Document == "" && c.DocumentCpf == "" {
		return errors.New("the document or documentCpf is required")
	}
	if c.Document != "" && len(c.Document) != 14 {
		return errors.New("the document is invalid")
	}
	if c.DocumentCpf != "" && len(c.DocumentCpf) != 11 {
		return errors.New("the documentCpf is invalid")
	}
	if c.TenantId == 0 {
		return errors.New("the tenantId is invalid")
	}
	if c.CountryCode == "" {
		return errors.New("the countryCode is invalid")
	}
	if c.CityCode == "" {
		return errors.New("the cityCode is invalid")
	}
	return nil
}

func NewParticipant(tenantId uint64, originCode string, name string, document string, documentCpf string, registration string, countryCode string, cityCode string, street string, number string, complement string, neighborhood string) (*Participant, error) {
	model := Participant{
		Id:           0,
		OriginCode:   originCode,
		TenantId:     tenantId,
		Name:         name,
		Document:     document,
		DocumentCpf:  documentCpf,
		Registration: registration,
		CountryCode:  countryCode,
		CityCode:     cityCode,
		Street:       street,
		Number:       number,
		Neighborhood: neighborhood,
		Complememt:   complement,
	}
	return NewParticipantEntity(model)
}

func CreateParticipant(tenantId uint64, line string) (*Participant, error) {
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
	participant.CountryCode, err = utils.CopyText(line, 4)
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
	participant.CityCode, err = utils.CopyText(line, 8)
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
