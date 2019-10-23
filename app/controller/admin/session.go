package admin

import (
	"adm/app/models"
	"adm/app/pkg/app"
	"adm/app/pkg/view"
	"net/http"
)

func GetLogin(w http.ResponseWriter, r *http.Request) error {
	v := view.New(r)
	v.Data["Email"] = ""
	v.Render(w, "admin/sessions/new")
	return nil
}

func PostLogin(w http.ResponseWriter, r *http.Request) error {
	session := app.Session(r)

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		loginErr(w, r, email)
		return nil
	}

	adminUser, err := models.GetAdminByEmail(email)
	if err != nil {
		loginErr(w, r, email)
		return nil
	}

	valid := adminUser.CheckAuth(password)

	if valid == true {
		session.Values["admin_user_id"] = adminUser.ID
		session.Save(r, w)

		if err != nil {
			loginErr(w, r, email)
		}

		view.SuccessFlash(w, r, "Logged in successfully")
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return nil
	}

	loginErr(w, r, email)
	return nil
}

func Logout(w http.ResponseWriter, r *http.Request) error {
	session := app.Session(r)
	app.EmptySession(session)

	view.SuccessFlash(w, r, "Logged out successfully.")
	http.Redirect(w, r, "/admin/sessions/new", http.StatusSeeOther)
	return nil
}

func loginErr(w http.ResponseWriter, r *http.Request, email string) {
	v := view.New(r)
	view.ErrorFlash(w, r, "Invalid credentials.")
	v.Data["Email"] = email
	v.Render(w, "admin/sessions/new")
}
