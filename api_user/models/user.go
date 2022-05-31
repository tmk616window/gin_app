package models

import "time"

// id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
// uuid VARCHAR(50) NOT NULL UNIQUE,
// name VARCHAR(50),
// email VARCHAR(50) UNIQUE,
// password VARCHAR(50),
// avatar TEXT,
// live VARCHAR(50),
// details TEXT,
// age INT,
// created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)`, tableNameUser)


type User struct {
	ID int
	UUID string
	Name string
	Email string
	Password string
	Created_at time.Time
}