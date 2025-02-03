package database

import (
	"context"
	"database/sql"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	Post *PostRepository
	User *UserRepository
}

func newDB() *sqlx.DB {
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

func NewRawRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Post: NewPostRepository(db),
		User: NewUserRepository(db),
	}
}

func NewRepository() (*Repository, *sqlx.DB) {
	db := newDB()
	return NewRawRepository(db), db
}

func WithTransaction(db *sqlx.DB, fn func(tx *sqlx.Tx) error) error {
	tx, err := db.BeginTxx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})

	if err != nil {
		return err
	}

	err = fn(tx)

	if err != nil {
		rollErr := tx.Rollback()

		if rollErr != nil {
			return rollErr
		}

		return err
	}

	return tx.Commit()
}
