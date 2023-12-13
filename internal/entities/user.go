package entities

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        float64 `gorm:"primaryKey;autoIncrement:true"`
	TenantId  float64
	Tenant    Tenant `gorm:"constraint:OnUpdate:NULL,OnDelete:SET NULL;"`
	Name      string `gorm:"size:100"`
	Login     string `gorm:"size:100;index:idx_Login,unique"`
	Password  string `gorm:"size:100"`
	StartDate time.Time
	EndDate   time.Time
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

func NewUser(tenantId float64, name string, login string, password string) (*User, error) {
	model := User{
		Id:        0,
		TenantId:  tenantId,
		Name:      name,
		Login:     login,
		Password:  password,
		StartDate: time.Now(),
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
