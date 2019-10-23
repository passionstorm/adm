package models_test

import (
	"adm/testutils"
	"os"
	"testing"
	"github.com/gobuffalo/envy"
)

func TestMain(m *testing.M) {
	envy.Set("MIGRATIONS_DIR", "../migrations")
	testutils.SetupDB()

	c := m.Run()

	testutils.DropDB()
	os.Exit(c)
}
