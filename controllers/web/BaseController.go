package web

import (
	"github.com/CloudyKit/jet"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type View struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Data           jet.VarMap
}


func (t *View) render(view string) {
	var root, _ = os.Getwd()
	var View = jet.NewHTMLSet(filepath.Join(root, "views"))
	templ, err := View.GetTemplate(view + ".jet")
	if err != nil {
		log.Println(err)
	}

	_ = templ.Execute(t.ResponseWriter, t.Data, nil)
}
