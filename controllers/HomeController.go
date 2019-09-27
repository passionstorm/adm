package controllers

import (
	"adm/model"
	"adm/pkg/web"
)

type HomeController struct{}

func (t *HomeController) Index(v *web.View) {
	md := model.Account{}


	v.Data.Set("users", md.GetAllUser())


	v.Render("home")
}
