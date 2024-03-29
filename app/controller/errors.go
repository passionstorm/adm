package controller

import (
	"adm/app/pkg/view"
	"log"
	"net/http"
)

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		switch e := err.(type) {
		case StatusError:
			log.Printf("HTTP %d - %+v\n", e.Status(), e.Err)
			v := view.New(r)
			v.RenderError(w, e.Status())
		default:
			log.Printf("%+v\n", err)
			v := view.New(r)
			v.RenderError(w, http.StatusInternalServerError)
		}
	}
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.RenderError(w, http.StatusInternalServerError)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.RenderError(w, http.StatusNotFound)
}
