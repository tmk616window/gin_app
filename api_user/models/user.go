package models

// import (
// 	"time"
// )


func GetUserByEmail(email string) User {
    db := DbConnect()
    var user User
    db.First(&user, "email = ?", email)
    db.Close()
    return user
}

func GetUserByUuid(uuid string) (User) {
    db := DbConnect()
    var user User
    db.First(&user, "uuid = ?", uuid)
    db.Close()
    return user
}




