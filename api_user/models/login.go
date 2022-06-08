package models

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func Login(c *gin.Context) {
	user := User{}
	c.BindJSON(&user)
	u := getUserByEmail(user.Email)
	fmt.Printf(u.Password)
	
	if u.Email == "" {
		// err := errors.New("名前が一致するユーザーが存在しません。")
		c.JSON(500, gin.H{"error": "名前が一致するユーザーが存在しません。"})
	}

	if err := compareHashAndPassword(u.Password, user.Password); err != nil {
		log.Println("ログインできませんでした")
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"err": err})
		c.Abort()
	} else {
		log.Println("ログインできました")
		c.JSON(200, gin.H{"user": u})
	}
}

func getUserByEmail(email string) User {
    db := DbConnect()
    var user User
    db.First(&user, "email = ?", email)
    db.Close()
    return user
}

func compareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

