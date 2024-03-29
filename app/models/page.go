package models

import (
	"github.com/markbates/pop/nulls"
	"github.com/pkg/errors"
	"regexp"
	"strings"
	"time"
)

type PageModel struct {
	ID              int64
	Title           nulls.String      `db:"title"`
	PageTitle       nulls.String      `db:"page_title"`
	MetaDescription nulls.String      `db:"meta_description"`
	Content         nulls.String      `db:"content"`
	Slug            nulls.String      `db:"slug"`
	Layout          nulls.String      `db:"layout"`
	InsertedAt      time.Time         `db:"inserted_at"`
	UpdatedAt       time.Time         `db:"updated_at"`
	Errors          map[string]string `db:"-"`
}

func Page() *PageModel {
	return &PageModel{
		Layout: nulls.NewString("two-col"),
	}
}

func ListPages(offset, limit int) ([]*PageModel, int, error) {
	var pages []*PageModel
	err := db.Select(&pages, `
		SELECT 
			id, title, slug, inserted_at
		FROM pages 
		ORDER BY inserted_at DESC
		LIMIT ?, ?
	`, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	var count int
	err = db.Get(&count, `SELECT count(*) FROM pages`)
	if err != nil {
		return nil, 0, err
	}

	return pages, count, nil
}

func ListAllPages() ([]*PageModel, error) {
	var pages []*PageModel
	err := db.Select(&pages, `
		SELECT 
			id, title, slug, inserted_at
		FROM pages 
		ORDER BY inserted_at DESC
	`)
	if err != nil {
		return nil, err
	}
	return pages, nil
}

func GetPage(ID int) (*PageModel, error) {
	page := PageModel{}
	err := db.Get(&page, `
		SELECT 
			id, title, page_title, meta_description, content, slug, layout, inserted_at
		FROM pages 
		WHERE id = ?
	`, ID)
	if err != nil {
		return nil, err
	}

	return &page, nil
}

func GetPageBySlug(slug string) (*PageModel, error) {
	page := PageModel{}
	err := db.Get(&page, `
		SELECT 
			id, title, page_title, meta_description, content, slug, layout, inserted_at
		FROM pages 
		WHERE slug = ?
	`, slug)
	if err != nil {
		return nil, err
	}

	return &page, nil
}

func (p *PageModel) Create() error {
	p.InsertedAt = time.Now()
	p.UpdatedAt = time.Now()
	stmt, err := db.NamedExec(`
		INSERT INTO pages (slug, title, page_title, meta_description, layout, content, inserted_at, updated_at)
		VALUES 			  (:slug, :title, :page_title, :meta_description, :layout, :content, :inserted_at, :updated_at)
	`, p)

	if err != nil {
		return errors.WithStack(err)
	}
	p.ID, err = stmt.LastInsertId()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (p *PageModel) Delete() error {
	_, err := db.Exec("DELETE from pages WHERE pages.id = ?", p.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (p *PageModel) Update() error {
	p.UpdatedAt = time.Now()

	_, err := db.NamedExec(`
		UPDATE pages 
		SET
			title = :title,
			page_title = :page_title,
			meta_description = :meta_description,
			layout = :layout,
			content = :content,
			slug = :slug,
			updated_at = :updated_at
		WHERE id = :id
	`, p)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (p *PageModel) Validate() bool {
	p.Errors = make(map[string]string)

	title := strings.TrimSpace(p.Title.String)
	slug := strings.TrimSpace(p.Slug.String)
	content := strings.TrimSpace(p.Content.String)

	if title == "" {
		p.Errors["Title"] = "can't be blank"
	}

	if slug == "" {
		p.Errors["Slug"] = "can't be blank"
	}

	if content == "" {
		p.Errors["Content"] = "can't be blank"
	}

	re := regexp.MustCompile("^[\\w\\-]+$")
	matched := re.Match([]byte(slug))

	if matched == false {
		p.Errors["Slug"] = "URL slug is invalid (only a-z, 0-9 and - allowed)"
	}

	if p.Layout.String == "" {
		p.Errors["Layout"] = "must be valid"
	}

	return len(p.Errors) == 0
}
