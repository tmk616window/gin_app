package models

import (
	"api_user/config"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DbConnect() *gorm.DB {
    db, err := gorm.Open(config.Config.SQLDriver, "user:password@tcp(db:3306)/db?charset=utf8&parseTime=True&loc=Local")
    fmt.Println(config.Config.SQLDriver)
    if err != nil {
        panic(err.Error())
    }
	fmt.Println("db connected: ", &db)
	return db
}

func header(w http.ResponseWriter, r *http.Request, value string) {
	//要素「Accept-Encoding」を取得
	h := r.Header[value]
	fmt.Fprintln(w, h)
}
