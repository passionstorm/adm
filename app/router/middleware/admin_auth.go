package middleware

import (
	"adm/app/models"
	"adm/app/pkg/app"
	"adm/app/pkg/view"
	"context"
	"log"
	"net/http"
	"net/url"
)

func RequireAdminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := app.Session(r)
		userID, ok := session.Values["admin_user_id"]
		if !ok {
			view.ErrorFlash(w, r, "Please login to view that resource.")
			http.Redirect(w, r, "/admin/sessions/new?r="+url.PathEscape(r.URL.Path), http.StatusSeeOther)
			return
		}

		// find user
		u, err := models.GetAdminUser(userID.(int64))
		if err != nil {
			log.Println(err)
			view.ErrorFlash(w, r, "Please login to view that resource.")
			http.Redirect(w, r, "/admin/sessions/new", http.StatusSeeOther)
			return
		}

		ctx := context.WithValue(r.Context(), app.SessKey, u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
