package main

import (
	"api_user/controller"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	client := resty.New()
	res, _ := client.R().
		EnableTrace().
		Get("http://api_user:3000/users/db8eae88-e098-4109-88c3-210bdb346562")	
	var user controller.ResUser
	if err := json.Unmarshal(res.Body(), &user); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", user.ResUser.UUID)


	r := gin.Default()
    r.POST("/user", controller.Post)
	r.GET("/users/:uuid", controller.Show)
    r.POST("/login", controller.Login)
	r.Run(":3000")
}
