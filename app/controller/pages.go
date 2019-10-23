package controller

import (
	"adm/app/models"
	"adm/app/pkg/view"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

func ShowPage(w http.ResponseWriter, r *http.Request) error {
	slug := chi.URLParam(r, "slug")
	slug = strings.Split(slug, "#")[0]

	page, err := models.GetPageBySlug(slug)
	if err != nil {
		return StatusError{Code: 404, Err: errors.Wrapf(err, "slug requested: %s", slug)}
	}

	v := view.New(r)
	v.Data["PageModel"] = page
	s := page.Layout.String
	if s == "" {
		s = "two-col"
	}

	v.Render(w, fmt.Sprintf("pages/show-%s", s))
	return nil
}
