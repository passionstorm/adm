package controllers

import (
	"adm/model"
	"github.com/go-chi/render"
	"net/http"
)

type AccountApi struct {
}

func (account *AccountApi) GetInfo(w http.ResponseWriter, r *http.Request) {
	m := &model.Account{}
	render.JSON(w, r, m.GetAllUser())
}

func (account *AccountApi) Create() {

}
