package controller

import (
	"adm/app/pkg/view"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) error {
	v := view.New(r)
	v.Render(w, "home")
	return nil
}
