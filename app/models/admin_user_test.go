package models_test

import (
	"adm/app/models"
	"github.com/markbates/pop/nulls"
	"testing"
)

func Test_CheckAdminCredentials(t *testing.T) {
	admin := models.AdminUserModel{
		Email:    nulls.NewString("vohoangminhdn93@gmail.com"),
		Password: "admin123",
	}
	err := admin.Create()
	if err != nil {
		t.Error(err)
	}
	//defer testutils.ResetTable("admin_users")

	res := admin.CheckAuth("admin")
	if res == false {
		t.Error("Incorrect matching.")
	}

	res = admin.CheckAuth("foobar")
	if res != false {
		t.Error("Incorrect matching.")
	}
}
