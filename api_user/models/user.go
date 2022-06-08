package models

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID int `json:"-" gorm:"primary_key"`
	UUID string `json:"uuid" gorm:"unique"`
	Name string `json:"name" gorm:"unique;not null"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"-" gorm:"not null"`
	Age int `json:"age"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"index"`
	Created_at time.Time `json:"createdAt" gorm:"not null"`
}


func CreateUser(c *gin.Context) {
	user := User{}
	var err error
	c.BindJSON(&user)
	user.Password, err =  passwordEncrypt(user.Password)
	user.UUID = uuid.New().String()
	user.Created_at = time.Now()

	if err != nil {
		fmt.Println("パスワード暗号化中にエラーが発生しました。：", err)
		// return  err
	}

	db := DbConnect()
	db.NewRecord(user)
	db.Create(&user)
}


func Show(c *gin.Context) {
	uuid := c.Param("uuid")
	user := getUserByUuid(uuid)
	if user.UUID == "" {
		c.JSON(200, gin.H{"message": "ありません"})
	} else {
		c.JSON(200, gin.H{"user": user})
	}
}


func getUserByUuid(uuid string) (User) {
    db := DbConnect()
    var user User
    db.First(&user, "uuid = ?", uuid)
    db.Close()
    return user
}




func passwordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}



