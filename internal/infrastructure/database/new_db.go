package database

import (
	"os"

    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

func NewDB() *sqlx.DB {
    host := os.Getenv("SQL_HOST")
    username := os.Getenv("SQL_USERNAME")
    password := os.Getenv("SQL_PASSWORD")
    db_name := os.Getenv("SQL_DATABASE")

    connStr := "postgres://" + username + ":" + password + "@" + host + "/" + db_name + "?sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)

    if err != nil {
        panic(err)
    }

    return db
}
