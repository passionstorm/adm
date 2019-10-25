package models

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gobuffalo/envy"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"log"
	"strings"
)

var (
	ErrAlreadyTaken = errors.New("is already taken")
)

const (
	DbName = "DB_NAME"
	DbHost = "DB_HOST"
	DbUser = "DB_USER"
	DbPass = "DB_PASS"
)

func downcase(str string) string {
	return strings.TrimSpace(strings.ToLower(str))
}

func newUUID() string {
	return uuid.NewV4().String()
}

func DB() *sqlx.DB {
	return db
}

var db *sqlx.DB

// InitDB sets up the database
func InitDB() *sqlx.DB {
	var err error
	dns := envy.Get(DbUser, "root") +
		":" + envy.Get(DbPass, "") + "@tcp(" + envy.Get(DbHost, "127.0.0.1:3306") + ")/" + envy.Get(DbName, "adm") + "?charset=utf8mb4"
	db, err = sqlx.Connect("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
