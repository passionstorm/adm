package admin

import (
	"adm/app/controller"
	"adm/app/models"
	"adm/app/pkg/view"
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/markbates/pop/nulls"
	"github.com/pkg/errors"
	"html/template"
	"net/http"
	"strconv"
)

func ListPages(w http.ResponseWriter, r *http.Request) error {
	v := view.New(r)
	pageNum := controller.GetPageNum(r)

	p, err := view.NewPaginator(pageNum, controller.PerPage, r.URL)
	if err != nil {
		return controller.StatusError{Code: 500, Err: errors.WithStack(err)}
	}

	pages, numPages, err := models.ListPages(p.Start, p.Limit)

	if err != nil {
		return controller.StatusError{Code: 500, Err: errors.WithStack(err)}
	}

	pagination := p.Render(numPages)
	if err != nil {
		return controller.StatusError{Code: 500, Err: errors.WithStack(err)}
	}

	v.Data["Pages"] = pages
	v.Data["Pagination"] = template.HTML(pagination)
	v.Render(w, "admin/pages/index")
	return nil
}

func NewPage(w http.ResponseWriter, r *http.Request) error {
	v := view.New(r)
	page := models.Page()

	v.Data["Page"] = page
	v.Render(w, "admin/pages/new")
	return nil
}

func CreatePage(w http.ResponseWriter, r *http.Request) error {
	v := view.New(r)
	page := models.Page()
	page.Title = nulls.NewString(r.FormValue("title"))
	page.PageTitle = nulls.NewString(r.FormValue("page_title"))
	page.Slug = nulls.NewString(r.FormValue("slug"))
	page.MetaDescription = nulls.NewString(r.FormValue("meta_description"))
	page.Content = nulls.NewString(r.FormValue("content"))
	page.Layout = nulls.NewString(r.FormValue("layout"))

	if page.Validate() {
		err := page.Create()
		if err != nil {
			if err == models.ErrAlreadyTaken {
				page.Errors["Slug"] = models.ErrAlreadyTaken.Error()
				v.Data["Page"] = page
				v.Render(w, "admin/pages/edit")
				return nil
			}

			return controller.StatusError{Code: 500, Err: err}
		}

		view.SuccessFlash(w, r, "PageModel was created successfully.")
		http.Redirect(w, r, fmt.Sprintf("/admin/pages/%d/edit", page.ID), http.StatusSeeOther)
		return nil
	}

	v.Data["Page"] = page
	v.Render(w, "admin/pages/new")
	return nil
}

func GetPage(w http.ResponseWriter, r *http.Request) error {
	v := view.New(r)
	ctx := r.Context()
	page, ok := ctx.Value("page").(*models.PageModel)

	if !ok {
		return controller.StatusError{Code: 404, Err: errors.New("unabled to retrieve page from context")}
	}

	v.Data["Page"] = page
	v.Render(w, "admin/pages/show")
	return nil
}

func EditPage(w http.ResponseWriter, r *http.Request) error {
	v := view.New(r)
	ctx := r.Context()
	page, ok := ctx.Value("page").(*models.PageModel)

	if !ok {
		return controller.StatusError{Code: 404, Err: errors.New("unabled to retrieve page from context")}
	}

	v.Data["Page"] = page
	v.Render(w, "admin/pages/edit")
	return nil
}

func UpdatePage(w http.ResponseWriter, r *http.Request) error {
	v := view.New(r)
	ctx := r.Context()
	page, ok := ctx.Value("page").(*models.PageModel)

	if !ok {
		return controller.StatusError{Code: 404, Err: errors.New("unabled to retrieve page from context")}
	}

	page.Title = nulls.NewString(r.FormValue("title"))
	page.PageTitle = nulls.NewString(r.FormValue("page_title"))
	page.Slug = nulls.NewString(r.FormValue("slug"))
	page.MetaDescription = nulls.NewString(r.FormValue("meta_description"))
	page.Content = nulls.NewString(r.FormValue("content"))
	page.Layout = nulls.NewString(r.FormValue("layout"))

	if page.Validate() {
		err := page.Update()
		if err != nil {
			if err == models.ErrAlreadyTaken {
				page.Errors["Slug"] = models.ErrAlreadyTaken.Error()
				v.Data["Page"] = page
				v.Render(w, "admin/pages/edit")
				return nil
			}

			return controller.StatusError{Code: 500, Err: err}
		}

		view.SuccessFlash(w, r, "PageModel was updated successfully.")
		http.Redirect(w, r, fmt.Sprintf("/admin/pages/%d/edit", page.ID), http.StatusSeeOther)
		return nil
	}

	v.Data["PageModel"] = page
	v.Render(w, "admin/pages/edit")
	return nil
}

func DeletePage(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	page, ok := ctx.Value("page").(*models.PageModel)

	if !ok {
		return controller.StatusError{Code: 404, Err: errors.New("unabled to retrieve page from context")}
	}

	err := page.Delete()
	if err != nil {
		return controller.StatusError{Code: 500, Err: errors.WithStack(err)}
	}

	view.SuccessFlash(w, r, "PageModel was deleted successfully.")
	http.Redirect(w, r, "/admin/pages", http.StatusSeeOther)
	return nil
}

func PageContext(next http.Handler) http.Handler {
	return controller.Handler(func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.Atoi(chi.URLParam(r, "ID"))

		if err != nil {
			return errors.WithStack(err)
		}

		page, err := models.GetPage(id)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("ID requested: %s", chi.URLParam(r, "ID")))
		}

		ctx := context.WithValue(r.Context(), "page", page)
		next.ServeHTTP(w, r.WithContext(ctx))

		return nil
	})
}
