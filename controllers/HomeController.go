package controllers

import (
	"adm/model"
	"adm/pkg/web"
	"encoding/json"
)

type HomeController struct{}

func (t *HomeController) Index(v *web.View) {
	md := model.Account{}
	users, _ := json.Marshal(md.GetAllUser())
	v.Data.Set("users", string(users))
	v.Render("home")
}
