package admin

import (
	"adm/app/controller"
	"adm/app/models"
	"adm/app/pkg/view"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"html/template"
	"net/http"
	"strconv"
)

func ListMessages(w http.ResponseWriter, r *http.Request) error {
	v := view.New(r)
	pageNum := controller.GetPageNum(r)

	p, err := view.NewPaginator(pageNum, controller.PerPage, r.URL)
	if err != nil {
		return controller.StatusError{Code: 500, Err: errors.WithStack(err)}
	}

	messages, numMessages, err := models.ListMessages(p.Start, p.Limit)

	if err != nil {
		return controller.StatusError{Code: 500, Err: errors.WithStack(err)}
	}

	pagination := p.Render(numMessages)
	if err != nil {
		return controller.StatusError{Code: 500, Err: errors.WithStack(err)}
	}

	v.Data["Messages"] = messages
	v.Data["Pagination"] = template.HTML(pagination)
	v.Render(w, "admin/support-messages/index")
	return nil
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(r, "ID"))
	if err != nil {
		return errors.WithStack(err)
	}

	message := models.SupportMessageModel{ID: int64(id)}
	err = message.Delete()
	if err != nil {
		return controller.StatusError{Code: 500, Err: errors.WithStack(err)}
	}

	view.SuccessFlash(w, r, "Support message was deleted successfully.")
	http.Redirect(w, r, "/admin/support-messages", http.StatusSeeOther)
	return nil
}
