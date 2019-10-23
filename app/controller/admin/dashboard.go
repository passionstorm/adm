package admin

import (
	"adm/app/pkg/view"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) error {
	v := view.New(r)
	v.Render(w, "admin/dashboard/index")
	return nil
}
