package db

import (
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB = nil
)

func connect() *sql.DB {
	log.Println("Connecting to database...")
	dsn := "host=localhost user=postgres password=postgres dbname=cmsdb port=5432 sslmode=disable"

	newDbInstance, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}

	return newDbInstance
}

func GetDatabase() *sql.DB {
	if db == nil {
		db = connect()
	}
	return db
}

func RunScript(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(content))

	db := GetDatabase()
	_, err = db.Exec(string(content))
	if err != nil {
		panic(err)
	}

	return nil
}
