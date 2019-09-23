package adm

import (
	"github.com/CloudyKit/jet"
	"net/http"
)

type Controller struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Data           jet.VarMap
}

func (c *Controller) New() {
	c.Data = make(jet.VarMap)
}


//func (c *Controller) View(view string) {
//	var root, _ = os.Getwd()
//	var View = jet.NewHTMLSet(filepath.Join(root, "views"))
//	t, err := View.GetTemplate(view)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	_ = t.Execute(c.ResponseWriter, c.Data, nil)
//}
