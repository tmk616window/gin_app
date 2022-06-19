package controller

import (
	"api_user/handler"
	"api_user/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Login(c *gin.Context) {
	user := models.User{}
	c.BindJSON(&user)
	u := models.GetUserByEmail(user.Email)
	fmt.Printf(u.Password)
	
	if u.Email == "" {
		// err := errors.New("名前が一致するユーザーが存在しません。")
		c.JSON(500, gin.H{"error": "名前が一致するユーザーが存在しません。"})
	}

	if err := handler.CompareHashAndPassword(u.Password, user.Password); err != nil {
		log.Println("ログインできませんでした")
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"err": err})
		c.Abort()
	} else {
		log.Println("ログインできました")
		c.JSON(200, gin.H{"user": u})
	}
}



