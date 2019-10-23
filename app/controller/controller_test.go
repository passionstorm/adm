package controller_test

import (
	"adm/app/pkg/view"
	"adm/testutils"
	"os"
	"testing"

	"github.com/gobuffalo/envy"
)

func TestMain(m *testing.M) {
	envy.Set("MIGRATIONS_DIR", "../migrations")
	testutils.SetupServer()
	testutils.SetupDB()

	envy.Set("TEMPLATES_DIR", "../templates")
	view.LoadTemplates()

	c := m.Run()

	testutils.DropDB()
	testutils.CloseServer()
	os.Exit(c)
}
