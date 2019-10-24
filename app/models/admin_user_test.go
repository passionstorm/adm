package models_test

import (
	"adm/app/models"
	"adm/testutils"
	"github.com/markbates/pop/nulls"
	"testing"
)

func Test_CheckAdminCredentials(t *testing.T) {
	admin := models.AdminUserModel{
		Email:    nulls.NewString("admin@example.com"),
		Password: "testing123",
	}
	admin.Create()
	defer testutils.ResetTable("admin_users")

	res := admin.CheckAuth("testing123")
	if res == false {
		t.Error("Incorrect matching.")
	}

	res = admin.CheckAuth("foobar")
	if res != false {
		t.Error("Incorrect matching.")
	}
}
