package web

import (
	"adm/model"
	"github.com/CloudyKit/jet"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Index(v *View) {
	md := model.Account{}
	v.Data.Set("users", md.GetAllUser())
	v.render("home")
}

func Test(w http.ResponseWriter, r *http.Request)  {

	md := model.Account{}
	view := "home"
	data := make(jet.VarMap)
	data.Set("users", md.GetAllUser())
	var root, _ = os.Getwd()
	var View = jet.NewHTMLSet(filepath.Join(root, "views"))
	templ, err := View.GetTemplate(view + ".jet")
	if err != nil {
		log.Println(err)
	}
	_ = templ.Execute(w, data, nil)
}