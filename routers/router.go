package routers

import (
	"adm/controllers"
	"adm/pkg/web"
	"flag"
	"github.com/CloudyKit/jet"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

func do(act func(action *web.View)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		act(&web.View{
			ResponseWriter: w,
			Request:        r,
			Data:           make(jet.VarMap),
		})
	}
}

func Load() *chi.Mux {
	r := chi.NewRouter()
	//r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})
	homeCtl := controllers.HomeController{}
	r.Get("/", do(homeCtl.Index))
	return r
}
