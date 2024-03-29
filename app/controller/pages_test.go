package controller_test

import (
	"adm/app/models"
	"adm/testutils"
	"fmt"
	"github.com/markbates/pop/nulls"
	"net/http"
	"strings"
	"testing"
)

func Test_Pages(t *testing.T) {
	page := models.Page()
	page.Title = nulls.NewString("test")
	page.Content = nulls.NewString("test")
	page.Slug = nulls.NewString("test")

	err := page.Create()
	if err != nil {
		t.Error(err)
	}
	defer testutils.ResetTable("pages")

	code, resp := testutils.Get(t, fmt.Sprintf("/pages/%s", page.Slug.String))
	if code != http.StatusOK {
		t.Errorf("Error code: %d", code)
	}

	if !strings.Contains(resp, "test") {
		t.Errorf("incorrect response: %s", resp)
	}

	// With anchor. No idea why this is getting passed to go handler
	// in some cases. Sentry reports this as a bug (slug not found).
	code, resp = testutils.Get(t, fmt.Sprintf("/pages/%s#support", page.Slug.String))
	if code != http.StatusOK {
		t.Errorf("Error code: %d", code)
	}
}
