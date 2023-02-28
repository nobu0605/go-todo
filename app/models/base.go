package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"go-todo/config"
	"log"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error


func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING UNIQUE,
		password STRING,
		created_at DATETIME)`, "users")

	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title STRING,
		user_id INTEGER,
		description TEXT,
		status_id INTEGER,
		created_at DATETIME)`, "todos")

	Db.Exec(cmdT)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		expiration_date DATETIME,
		created_at DATETIME)`, "sessions")

	Db.Exec(cmdS)

	cmdST := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name STRING UNIQUE,
		created_at DATETIME)`, "statuses")

	Db.Exec(cmdST)

	// for _, status := range constants.GetStatuses() {
		
	// 	fmt.Println(status.String())
	// 	cmdIST := fmt.Sprintf(`insert into statuses (
	// 		name, 
	// 		created_at) values (?, ?)`)
		
	//  	_, err = Db.Exec(cmdIST,status.String(),time.Now())
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	fmt.Println(err)
	// }
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
