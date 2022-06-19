package models

import (
	"api_test/config"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Test struct {
	ID int `json:"-" gorm:"primary_key"`
	TestUUID string `json:"testUuid" gorm:"unique;not null"`
	Title string `json:"title" gorm:"not null"`
	Text string `json:"text" gorm:"not null"`
	UserUUID string `json:"user_uuid" gorm:"not null"`
	// User User `gorm:"foreignKey:UserUUID;references:UserUUID`
	// User User `gorm:"foreignKey:UUID;references:UUID"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"not null""`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
}

func DbConnect() *gorm.DB {
    db, err := gorm.Open(config.Config.SQLDriver, "user:password@tcp(test_db:3306)/test_db?charset=utf8&parseTime=True&loc=Local")
    fmt.Println(config.Config.SQLDriver)
    if err != nil {
        panic(err.Error())
    }
	fmt.Println("db connected: ", &db)
	return db
}


// func DropTable() {
// 	db := DbConnect()
	

// }
