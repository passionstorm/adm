package routers

import (
	"adm/api/v1"
	"flag"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

func Load() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/user", func(r chi.Router) {
		accountApi := api.AccountApi{}
		r.Get("/info", accountApi.GetInfo)
	})

	return r
}
