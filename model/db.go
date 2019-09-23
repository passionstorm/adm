package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var db *sqlx.DB

func CloseDB() {
	defer db.Close()
}

func Log(o interface{}, isBeauty bool) {
	var out bytes.Buffer
	bytes, _ := json.Marshal(o)
	if isBeauty {
		_ = json.Indent(&out, bytes, "", "  ")
		log.Println(out.String())
	} else {
		log.Println(string(bytes))
	}
}

func Connect() *sqlx.DB {
	source := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	var err error
	db, err = sqlx.Open("mysql", source)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
