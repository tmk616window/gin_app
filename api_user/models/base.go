package models

import (
	"api_user/config"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var Db *sql.DB

var err error

const tableNameUser = "users"

func init() {
	conDB()
}

func conDB() {
	Db, err := sql.Open(config.Config.SQLDriver , "user:password@tcp(db:3306)/db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}

	cmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		uuid VARCHAR(50) NOT NULL UNIQUE,
		name VARCHAR(50),
		email VARCHAR(50) UNIQUE,
		password VARCHAR(50),
		avatar TEXT,
		live VARCHAR(50),
		details TEXT,
		age INT,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)`, tableNameUser)
		Db.Exec(cmd)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}



