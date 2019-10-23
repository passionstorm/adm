package testutils

import (
	"adm/app/models"
	"fmt"
	"log"
	"github.com/gobuffalo/envy"
	"github.com/pressly/goose"
)

func SetupDB() {
	td := envy.Get("TEST_DATABASE_URL", "postgres://postgres:@localhost/alloy_test?sslmode=disable")
	envy.Set("DATABASE_URL", td)
	migrationsDir, err := envy.MustGet("MIGRATIONS_DIR")
	if err != nil {
		log.Fatal("MIGRATIONS_DIR variable not set")
	}

	models.InitDB()
	if err := goose.Run("up", models.DB().DB, migrationsDir); err != nil {
		log.Fatal("ERR", err)
	}
}

func DropDB() {
	migrationsDir, err := envy.MustGet("MIGRATIONS_DIR")
	if err != nil {
		log.Fatal("MIGRATIONS_DIR variable not set")
	}
	if err := goose.Run("down", models.DB().DB, migrationsDir); err != nil {
		log.Fatal("ERR", err)
	}
}

func ResetTable(tablename string) {
	log.Printf("DELETE from %s", tablename)
	services.DB.Exec(fmt.Sprintf(`DELETE FROM %s`, tablename))
}
