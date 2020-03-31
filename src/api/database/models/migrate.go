package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Migrate automigrates models using ORM
func Migrate(db *gorm.DB) {
	Rollback(db)
	db.AutoMigrate(&User{}, &CheckIn{}, &Role{}, &Project{}, &Division{}, &DeviceToken{})

	Seeder(db)
	fmt.Println("Auto Migration has beed processed")
}

func Seeder(db *gorm.DB) {
	tx := db.Begin()

	tx.Create(&Division{
		Name: "MNG",
		Rank: 1,
	})
	tx.Create(&Division{
		Name: "DEV",
		Rank: 2,
	})
	tx.Create(&Division{
		Name: "COM",
		Rank: 4,
	})
	tx.Create(&Division{
		Name: "ADM",
		Rank: 5,
	})

	tx.Create(&Role{
		Name: "Admin",
		Rank: 1,
	})
	tx.Create(&Role{
		Name: "Manager",
		Rank: 2,
	})
	tx.Create(&Role{
		Name: "Dev",
		Rank: 4,
	})
	tx.Create(&Role{
		Name: "Mem",
		Rank: 9,
	})

	tx.Commit()
}

func Rollback(db *gorm.DB) {
	db.DropTableIfExists(&Role{}, &Division{})
}
