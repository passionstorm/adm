package models_test

import (
	"adm/testutils"
	"github.com/gobuffalo/envy"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func init() {
	var err error

	log.SetFlags(log.LstdFlags | log.Llongfile)
	p, _ := filepath.Abs("../../.env")
	err = envy.Load(p)
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {

	envy.Set("MIGRATIONS_DIR", "../migrations")
	testutils.SetupDB()

	c := m.Run()

	testutils.DropDB()
	os.Exit(c)
}
