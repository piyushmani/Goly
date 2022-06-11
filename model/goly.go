package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Goly struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Goly     string `json:"goly" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
}

func Setup() {
	dsn := "host=172.17.0.2 user=admin password=admin dbname=admin port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Goly{})
	if err != nil {
		fmt.Println(err)
	}
}

func CreateGoly(goly Goly) error {
	tx := db.Create(&goly)
	return tx.Error
}

func GetGoly(id int) (Goly, error) {
	var goly Goly

	tx := db.Where("id = ?", id).First(&goly)

	if tx.Error != nil {
		return Goly{}, tx.Error
	}

	return goly, nil
}

func FindByGolyUrl(url string) (Goly, error) {
	var goly Goly
	tx := db.Where("goly = ?", url).First(&goly)
	return goly, tx.Error
}

func UpdateGoly(goly Goly) error {
	tx := db.Save(&goly)
	return tx.Error
}

func DeleteGoly(goly Goly) error {
	tx:= db.Delete(&goly)
	return tx.Error
}
