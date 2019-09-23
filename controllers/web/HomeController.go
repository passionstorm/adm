package web

import (
	"adm/model"
)

func Index(v *View) {
	md := model.Account{}
	v.Data.Set("users", md.GetAllUser())
	v.render("home")
}
