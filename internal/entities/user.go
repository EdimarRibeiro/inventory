package entities

import (
	"errors"
	"time"
)

type User struct {
	Id        uint64     `gorm:"primaryKey;autoIncrement:true" json:"-"`
	TenantId  uint64     `json:"-"`
	Tenant    Tenant     `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"-"`
	Name      string     `gorm:"size:100"`
	Login     string     `gorm:"size:100;index:idx_Login,unique"`
	Password  string     `gorm:"size:100" json:"-"`
	StartDate time.Time  `gorm:"datetime"`
	EndDate   *time.Time `gorm:"datetime;default:null"`
}

func (c *User) Validate() error {
	if c.Login == "" {
		return errors.New("the login is required")
	}
	if c.Password == "" {
		return errors.New("the password is required")
	}
	if c.TenantId == 0 {
		return errors.New("the tenantId is inválid")
	}
	return nil
}

func NewUser(tenantId uint64, name string, login string, password string) (*User, error) {
	model := User{
		Id:        0,
		TenantId:  tenantId,
		Name:      name,
		Login:     login,
		Password:  password,
		StartDate: time.Now(),
		EndDate:   nil,
	}
	return NewUserEntity(model)
}

func NewUserEntity(entity User) (*User, error) {
	err := entity.Validate()
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
