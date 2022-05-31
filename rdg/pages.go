package rdg

import (
	"database/sql"
	"db2/cms/db"
	"time"
)

type Page struct {
	Id        int
	Title     string
	Body      string
	Slug      string
	UpdatedAt time.Time
}

func rowsToPages(rows *sql.Rows) ([]Page, error) {
	res := []Page{}
	for rows.Next() {
		page := Page{}
		err := rows.Scan(&page.Id, &page.Title, &page.Slug)
		if err != nil {
			return nil, err
		}
		res = append(res, page)
	}

	return res, nil
}

// func GetPages(siteId int) ([]Page, error) {
// 	sql := `SELECT p.id, p.title, p.body, p.slug, p.updated_at FROM pages p WHERE site_id = $1 ORDER BY p.updated_at DESC LIMIT 10`
// 	rows, err := db.GetDatabase().Query(sql, siteId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	return rowsToPages(rows)
// }

func ListPages(siteId int) ([]Page, error) {
	sql := `SELECT p.id, p.title, p.slug FROM pages p WHERE p.site_id = $1 ORDER BY p.created_at DESC`
	rows, err := db.GetDatabase().Query(sql, siteId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return rowsToPages(rows)
}

func GetPage(slug string) (Page, error) {
	sql := `SELECT p.id, p.title, p.body, p.slug, p.updated_at FROM pages p WHERE p.slug = $1`
	row := db.GetDatabase().QueryRow(sql, slug)
	page := Page{}
	err := row.Scan(&page.Id, &page.Title, &page.Body, &page.Slug, &page.UpdatedAt)
	if err != nil {
		return Page{}, err
	}
	return page, nil
}

func AddPage(siteId string, title string, body string, slug string) error {
	sql := `INSERT INTO pages (site_id, title, body, slug) VALUES ($1, $2, $3, $4)`
	_, err := db.GetDatabase().Exec(sql, siteId, title, body, slug)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePage(id string, title string, body string) error {
	sql := `UPDATE pages SET title = $1, body = $2 WHERE id = $3`
	_, err := db.GetDatabase().Exec(sql, title, body, id)
	if err != nil {
		return err
	}
	return nil
}

func DeletePage(id string) error {
	sql := `DELETE FROM pages WHERE id = $1`
	_, err := db.GetDatabase().Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}
