package main

import (
	"api_user/models"
)

func main() {
	db := models.DbConnect()

	defer db.Close()
	db.LogMode(true)   

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&models.User{})
}
