package models

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"log"
	"os"
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
	db, err = sqlx.Connect("mysql", os.Getenv(DbUser)+
		":"+os.Getenv(DbPass)+"@tcp("+os.Getenv(DbHost)+")/"+os.Getenv(DbName)+"?charset=utf8mb4")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
