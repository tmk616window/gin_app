package models

import (
	"api_user/config"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID int `json:"-" gorm:"primary_key"`
	UUID string `json:"uuid" gorm:"unique"`
	Name string `json:"name" gorm:"unique;not null"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"-" gorm:"not null"`
	Age int `json:"age"`
	Tests []Test `gorm:"foreignKey:TestUUID;references:UUID"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"index"`
	Created_at time.Time `json:"createdAt" gorm:"not null"`
}

type Test struct {
	ID int `json:"-" gorm:"primary_key"`
	TestUUID string `json:"testUuid" gorm:"unique;not null"`
	Title string `json:"title" gorm:"not null"`
	Text string `json:"text" gorm:"not null"`
	// UserUUID string `json:"user_uuid" gorm:"not null"`
	// User User `gorm:"foreignKey:UserUUID;references:UserUUID`
	User User `gorm:"foreignKey:UUID;references:UUID"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"not null""`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
}


func DbConnect() *gorm.DB {
    db, err := gorm.Open(config.Config.SQLDriver, "user:password@tcp(db:3306)/db?charset=utf8&parseTime=True&loc=Local")
    fmt.Println(config.Config.SQLDriver)
    if err != nil {
        panic(err.Error())
    }
	fmt.Println("db connected: ", &db)
	return db
}
