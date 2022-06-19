package controller

import (
	"api_test/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ResUser struct {
	ResUser models.User 
}

func Post(user ResUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		test := models.Test{}
		var err error
		c.BindJSON(&test)
		test.TestUUID = uuid.New().String()
		test.CreatedAt = time.Now()
		test.Title = "test1"
		test.Text = "test1"
		test.UserUUID = user.ResUser.UUID
	
		if err != nil {
			fmt.Println("パスワード暗号化中にエラーが発生しました。：", err)
			// return  err
		}
	
		db := models.DbConnect()
		db.NewRecord(test)
		db.Create(&test)	
	}
}
