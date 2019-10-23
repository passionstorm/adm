package middleware

import (
	"adm/app/controller"
	"log"
	"net/http"
	"runtime/debug"
)

func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				log.Println(rvr)
				log.Println(string(debug.Stack()))
				controller.ErrorHandler(w, r)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
