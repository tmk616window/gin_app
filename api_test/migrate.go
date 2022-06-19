package main

import (
	"api_test/models"
)

func main() {
	db := models.DbConnect()
	defer db.Close()
	db.LogMode(true)   
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&models.Test{})
	// Add table suffix when creating tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Test{})
	
	
}
