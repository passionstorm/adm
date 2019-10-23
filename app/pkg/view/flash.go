package view

import (
	"adm/app/pkg/app"
	"encoding/gob"
	"net/http"
)

var (
	FlashInfo    = "alert-info"
	FlashWarning = "alert-warning"
	FlashSuccess = "alert-success"
	FlashError   = "alert-danger"
)

type Flash struct {
	Message string
	Class   string
}

func init() {
	gob.Register(Flash{})
}

func ErrorFlash(w http.ResponseWriter, r *http.Request, message string) {
	sess := app.Session(r)
	sess.AddFlash(Flash{message, FlashError})
	_ = sess.Save(r, w)
}

func SuccessFlash(w http.ResponseWriter, r *http.Request, message string) {
	sess := app.Session(r)
	sess.AddFlash(Flash{message, FlashSuccess})
	_ = sess.Save(r, w)
}

func InfoFlash(w http.ResponseWriter, r *http.Request, message string) {
	sess := app.Session(r)
	sess.AddFlash(Flash{message, FlashInfo})
	_ = sess.Save(r, w)
}
