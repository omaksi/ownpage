package rdg

import (
	"database/sql"
	"db2/cms/db"
	"time"
)

type Post struct {
	Id        int
	Body      string
	UpdatedAt time.Time
}

func rowsToPosts(rows *sql.Rows) ([]Post, error) {
	res := []Post{}
	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.Id, &post.Body, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, post)
	}

	return res, nil
}

func GetPosts(siteId int) ([]Post, error) {
	sql := `SELECT p.id, p.body, p.updated_at FROM posts p WHERE site_id = $1 ORDER BY p.updated_at DESC LIMIT 10`
	rows, err := db.GetDatabase().Query(sql, siteId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return rowsToPosts(rows)
}

func AddPost(siteId string, body string) error {
	sql := `INSERT INTO posts (site_id, body) VALUES ($1, $2)`
	_, err := db.GetDatabase().Exec(sql, siteId, body)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePost(postId string, body string) error {
	sql := `UPDATE posts SET body = $1 WHERE id = $2`
	_, err := db.GetDatabase().Exec(sql, body, postId)
	if err != nil {
		return err
	}
	return nil
}

func DeletePost(postId string) error {
	sql := `DELETE FROM posts WHERE id = $1`
	_, err := db.GetDatabase().Exec(sql, postId)
	if err != nil {
		return err
	}
	return nil
}
