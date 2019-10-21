package db

import (
	"adm/pkg/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"strings"
	"testing"
)

func TestGetTypeFromString(t *testing.T) {

	driver := "mysql"
	typeField := "Type"

	GetConnectionByDriver(driver).InitDB(map[string]config.Database{
		"default": {
			Host:       "127.0.0.1",
			Port:       "3306",
			User:       "root",
			Pwd:        "root",
			Name:       "go-admin-type-test",
			MaxIdleCon: 50,
			MaxOpenCon: 150,
			Driver:     DriverMysql,
		},
	})

	config.Set(&config.Config{
		SqlLog: true,
	})

	columnsModel, _ := WithDriver(driver).Get("all_types").ShowColumns()

	for _, model := range columnsModel {
		fieldTypeName := strings.ToUpper(getType(model[typeField].(string)))
		GetDTAndCheck(fieldTypeName)
	}

	item, _ := WithDriver(driver).Get("all_types").First()
	fmt.Println("item", item)
}

func getType(typeName string) string {
	r, _ := regexp.Compile(`\\(.*\\)`)
	typeName = r.ReplaceAllString(typeName, "")
	return strings.ToLower(strings.Replace(typeName, " unsigned", "", -1))
}
