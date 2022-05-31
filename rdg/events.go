package rdg

import (
	"database/sql"
	"db2/cms/db"
	"time"
)

type Event struct {
	Id        int
	Title     string
	Body      string
	EventDate time.Time
	UpdatedAt time.Time
}

func rowsToEvents(rows *sql.Rows) ([]Event, error) {
	res := []Event{}
	for rows.Next() {
		event := Event{}
		err := rows.Scan(&event.Id, &event.Title, &event.Body, &event.EventDate, &event.UpdatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, event)
	}

	return res, nil
}

func GetEvents(siteId int) ([]Event, error) {
	sql := `SELECT id, title, body, event_date, updated_at FROM events WHERE site_id = $1 ORDER BY event_date DESC LIMIT 10`
	rows, err := db.GetDatabase().Query(sql, siteId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return rowsToEvents(rows)
}

func AddEvent(siteId string, title string, body string, eventDate time.Time) error {
	sql := `INSERT INTO events (site_id,title, body, event_date) VALUES ($1, $2, $3)`
	_, err := db.GetDatabase().Exec(sql, siteId, title, body, eventDate)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEvent(id string, title string, body string, eventDate time.Time) error {
	sql := `UPDATE events SET title = $1, body = $2, event_date = $3 WHERE id = $4`
	_, err := db.GetDatabase().Exec(sql, title, body, eventDate, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEvent(id string) error {
	sql := `DELETE FROM events WHERE id = $1`
	_, err := db.GetDatabase().Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}
