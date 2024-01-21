package entities

import (
	"errors"

	"github.com/EdimarRibeiro/inventory/internal/utils"
)

/*0200*/
type Product struct {
	Id            uint64  `gorm:"primaryKey;autoIncrement:true" json:"id"`
	TenantId      uint64  `json:"tenantId"`
	Tenant        Tenant  `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"-"`
	OriginCode    string  `gorm:"size:60" json:"originCode"`
	Description   string  `gorm:"size:250" json:"description"`
	BarCode       string  `gorm:"size:60" json:"barCode"`
	OldOriginCode string  `gorm:"size:60" json:"oldOriginCode"`
	UnitId        string  `gorm:"size:6" json:"unitId"`
	Unit          Unit    `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"-"`
	Type          string  `gorm:"size:2" json:"type"`
	NcmCode       string  `gorm:"size:8" json:"ncmCode"`
	ExIpi         string  `gorm:"size:3" json:"exIpi"`
	GenderCode    string  `gorm:"size:2" json:"genderCode"`
	ServiceCode   string  `gorm:"size:5" json:"serviceCode"`
	AliqIcms      float64 `gorm:"type:decimal (8,2)" json:"aliqIcms"`
	CestCode      string  `gorm:"size:7" json:"cestCode"`
}

func (c *Product) Validate() error {
	if c.OriginCode == "" {
		return errors.New("the product originCode is required")
	}

	if c.Description == "" {
		return errors.New("the description is required")
	}

	if c.TenantId == 0 {
		return errors.New("the tenantId is invalid")
	}

	if c.UnitId == "" {
		return errors.New("the unitId is required")
	}

	if len(c.UnitId) < 2 {
		return errors.New("the min value unitId is invalid! " + c.UnitId)
	}

	if len(c.UnitId) > 6 {
		return errors.New("the max value unitId is invalid! " + c.UnitId)
	}

	if c.Type == "" {
		return errors.New("the type is required")
	}

	if c.NcmCode != "" && len(c.NcmCode) != 8 {
		return errors.New("the value ncmCode is invalid")
	}

	if c.GenderCode != "" && len(c.GenderCode) != 2 {
		return errors.New("the genderCode is required")
	}

	if c.CestCode != "" && len(c.CestCode) != 7 {
		return errors.New("the value cestCode is invalid")
	}
	return nil
}

func NewProduct(tenantId uint64, description string, barCode string, oldOriginCode string, unitId string, Type string, ncmCode string, exIpi string, genderCode string, serviceCode string, aliqIcms float64, cestCode string, originCode string) (*Product, error) {
	model := Product{
		Id:            0,
		TenantId:      tenantId,
		OriginCode:    originCode,
		Description:   description,
		BarCode:       barCode,
		OldOriginCode: oldOriginCode,
		UnitId:        unitId,
		Type:          Type,
		NcmCode:       ncmCode,
		ExIpi:         exIpi,
		GenderCode:    genderCode,
		ServiceCode:   serviceCode,
		AliqIcms:      aliqIcms,
		CestCode:      cestCode,
	}
	return NewProductEntity(model)
}
func CreateProduct(tenantId uint64, line string) (*Product, error) {
	var err error = nil
	product := Product{}

	product.Id, err = 0, nil
	product.TenantId, err = tenantId, nil
	product.OriginCode, err = utils.CopyText(line, 2)
	if err != nil {
		return nil, err
	}
	product.OriginCode, err = utils.CopyText(line, 2)
	if err != nil {
		return nil, err
	}
	product.Description, err = utils.CopyText(line, 3)
	if err != nil {
		return nil, err
	}
	product.BarCode, err = utils.CopyText(line, 4)
	if err != nil {
		return nil, err
	}
	product.OldOriginCode, err = utils.CopyText(line, 5)
	if err != nil {
		return nil, err
	}
	product.UnitId, err = utils.CopyText(line, 6)
	if err != nil {
		return nil, err
	}
	product.Type, err = utils.CopyText(line, 7)
	if err != nil {
		return nil, err
	}
	product.NcmCode, err = utils.CopyText(line, 8)
	if err != nil {
		return nil, err
	}
	product.ExIpi, err = utils.CopyText(line, 9)
	if err != nil {
		return nil, err
	}
	product.GenderCode, err = utils.CopyText(line, 10)
	if err != nil {
		return nil, err
	}
	product.ServiceCode, err = utils.CopyText(line, 11)
	if err != nil {
		return nil, err
	}
	product.AliqIcms, err = utils.CopyTextFloat(line, 12, 2)
	if err != nil {
		return nil, err
	}
	product.CestCode, err = utils.CopyText(line, 13)
	if err != nil {
		return nil, err
	}
	return NewProductEntity(product)
}
func NewProductEntity(entity Product) (*Product, error) {
	err := entity.Validate()
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
