package controller

import (
	"adm/app/models"
	"adm/app/pkg/mailer"
	"adm/app/pkg/view"
	"github.com/markbates/pop/nulls"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type recaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []int     `json:"error-codes"`
}

func NewSupportMessage(w http.ResponseWriter, r *http.Request) error {
	v := view.New(r)
	message := models.SupportMessageModel{}

	v.Data["Message"] = message
	v.Render(w, "support-messages/new")
	return nil
}

func CreateSupportMessage(w http.ResponseWriter, r *http.Request) error {
	v := view.New(r)
	message := models.SupportMessageModel{}
	message.Name = nulls.NewString(r.FormValue("name"))
	message.Email = nulls.NewString(r.FormValue("email"))
	message.Subject = nulls.NewString(r.FormValue("subject"))
	message.Content = nulls.NewString(r.FormValue("content"))

	if message.Validate() {
		err := message.Create()

		if err != nil {
			return StatusError{Code: 500, Err: errors.Wrapf(err, "Message: %v", message)}
		}

		go func() {
			mailer.NewSupportMail(&message)
			mailer.NewSupportNotification(&message)
		}()

		view.SuccessFlash(w, r, "Thank you for contacting us. We will get back to you within 2 business days")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	v.Data["Message"] = message
	v.Render(w, "support-messages/new")
	return nil
}
