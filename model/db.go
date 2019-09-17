package model

import (
	"adm/pkg/setting"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func AutoMigrate() {
	db.Debug().AutoMigrate(&Account{})
}

func ConnectDB() *gorm.DB {
	conf := setting.DbConfig
	db = connect(conf)

	return db
}

func CloseDB() {
	defer db.Close()
}

func connect(conf *setting.Database) *gorm.DB {
	source := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Name,
	)

	db, err := gorm.Open(conf.Type, source)
	if err != nil {
		panic(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.TablePrefix + defaultTableName
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return db.Debug()
}
