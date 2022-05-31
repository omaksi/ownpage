package rdg

import "db2/cms/db"

type Site struct {
	Id          int
	Title       string
	Description string
}

func GetSite(siteId int) (Site, error) {
	sql := `SELECT id, title, description FROM sites WHERE id = $1`
	row := db.GetDatabase().QueryRow(sql, siteId)
	site := Site{}
	err := row.Scan(&site.Id, &site.Title, &site.Description)
	if err != nil {
		return Site{}, err
	}
	return site, nil
}

func UpdateSite(site Site) error {
	sql := `UPDATE sites SET title = $1, description = $2 WHERE id = $3`
	_, err := db.GetDatabase().Exec(sql, site.Title, site.Description, site.Id)
	return err
}
