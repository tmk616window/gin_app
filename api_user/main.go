package main

import (
	"api_user/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



func main() {
	r := gin.Default()
    r.POST("/user", models.CreateUser)
	r.GET("/users/:uuid", models.Show)
    r.POST("/login", models.Login)
	r.Run(":3000")



}
