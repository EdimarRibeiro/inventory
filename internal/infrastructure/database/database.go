package database

import (
	"fmt"

	"github.com/EdimarRibeiro/inventory/internal/entities"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() {
	dsn := "sqlserver://Sa:Cs@#1519@67d20619ebba.sn.mynetname.net:8433?database=inventories&connection+timeout=30"

	// Initialize the GORM DB instance
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to the database: %v", err))
	}

	// AutoMigrate creates tables based on the User struct
	db.AutoMigrate(&entities.City{})
	db.AutoMigrate(&entities.Country{})
	db.AutoMigrate(&entities.Document{})
	db.AutoMigrate(&entities.DocumentItem{})
	db.AutoMigrate(&entities.Inventory{})
	db.AutoMigrate(&entities.InventoryFile{})
	db.AutoMigrate(&entities.InventoryProduct{})
	db.AutoMigrate(&entities.Participant{})
	db.AutoMigrate(&entities.Person{})
	db.AutoMigrate(&entities.Product{})
	db.AutoMigrate(&entities.Tenant{})
	db.AutoMigrate(&entities.Unit{})
	db.AutoMigrate(&entities.UnitConvert{})
	db.AutoMigrate(&entities.User{})

	DB = db

	//cityRepo := &CityRepository{DB: db}
	//cityService := &entitiesinterface.CityRepositoryInterface{CityRepository: cityRepo}

}
