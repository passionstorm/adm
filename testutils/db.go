package testutils

import (
	"adm/app/models"
	"fmt"
	"log"
)

func SetupDB() {
	//td := envy.Get("TEST_DATABASE_URL", "mysql://sxjTBsC5M1:lP0RUhUvit/alloy_test?sslmode=disable")
	//envy.Set("DATABASE_URL", td)
	//migrationsDir, err := envy.MustGet("MIGRATIONS_DIR")
	//if err != nil {
	//	log.Fatal("MIGRATIONS_DIR variable not set")
	//}

	models.InitDB()
	//if err := goose.Run("up", models.DB().DB, migrationsDir); err != nil {
	//	log.Fatal(err)
	//}
}

func DropDB() {
	//migrationsDir, err := envy.MustGet("MIGRATIONS_DIR")
	//if err != nil {
	//	log.Fatal("MIGRATIONS_DIR variable not set")
	//}
	//if err := goose.Run("down", models.DB().DB, migrationsDir); err != nil {
	//	log.Fatal("ERR", err)
	//}
}

func ResetTable(tablename string) {
	log.Printf("DELETE from %s", tablename)
	models.DB().DB.Exec(fmt.Sprintf(`DELETE FROM %s`, tablename))
}
