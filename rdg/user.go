package rdg

import (
	"db2/cms/db"
	"log"
)

type User struct {
	Id       int
	Email    string
	Password string
}

func RegisterUser(email, password string) error {
	sql := `INSERT INTO users (email, password) VALUES ($1, $2)`
	_, err := db.GetDatabase().Exec(sql, email, password)
	return err
}

func GetUser(email string) (User, error) {
	sql := `SELECT id, email, password FROM users WHERE email = $1`
	row := db.GetDatabase().QueryRow(sql, email)
	user := User{}
	err := row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)
	}
	return user, err
}
