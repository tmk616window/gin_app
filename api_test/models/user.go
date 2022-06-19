package models

import "time"


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

