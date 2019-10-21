package web

import (
	"github.com/CloudyKit/jet"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type View struct {
	Response http.ResponseWriter
	Request  *http.Request
	Data     jet.VarMap
}

func (t *View) Render(view string) {
	var root, _ = os.Getwd()
	var View = jet.NewHTMLSet(filepath.Join(root, "views"))
	View.SetDevelopmentMode(true)
	templ, err := View.GetTemplate(view)
	if err != nil {
		log.Println(err)
	}
	err = templ.Execute(t.Response, t.Data, nil)
	if err != nil {
		log.Println(err)
	}
}
