package controller

import (
	"adm/model"
	"adm/pkg/context"
)

type HomeController struct{}

func (t *HomeController) Index(ctx *context.Context) {
	md := model.Account()
	ctx.Data.Set("users", md.All())
	ctx.Render("home")
}
