package models

import (
	"errors"
	"github.com/gobuffalo/envy"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"log"
	"strings"
)

var (
	ErrAlreadyTaken = errors.New("is already taken")
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
	dsn, err := envy.MustGet("DATABASE_URL")

	if err != nil {
		log.Fatalln(err)
	}
	db = sqlx.MustConnect("postgres", dsn)
	return db
}
