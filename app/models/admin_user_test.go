package models_test

import (
	"adm/app/models"
	"github.com/markbates/pop/nulls"
	"testing"
)

func Test_CheckAdminCredentials(t *testing.T) {
	admin := models.AdminUserModel{
		Email:    nulls.NewString("admin@gmail.com"),
		Password: "asdf1234",
	}
	err := admin.Create()
	if err != nil {
		t.Error(err)
	}
	//defer testutils.ResetTable("admin_users")

	res := admin.CheckAuth("asdf1234")
	if res == false {
		t.Error("Incorrect matching.")
	}
}
