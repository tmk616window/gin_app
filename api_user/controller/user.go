package controller

import (
	"api_user/handler"
	"api_user/models"

	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


func Show(c *gin.Context) {
	uuid := c.Param("uuid")
	user := models.GetUserByUuid(uuid)
	if user.UUID == "" {
		c.JSON(200, gin.H{"message": "ありません"})
	} else {
		c.JSON(200, gin.H{"resUser": user})
	}
}

func Post(c *gin.Context) {
	user := models.User{}
	var err error
	c.BindJSON(&user)
	user.Password, err =  handler.PasswordEncrypt(user.Password)
	user.UUID = uuid.New().String()
	user.Created_at = time.Now()

	if err != nil {
		fmt.Println("パスワード暗号化中にエラーが発生しました。：", err)
		// return  err
	}

	db := models.DbConnect()
	db.NewRecord(user)
	db.Create(&user)
}
