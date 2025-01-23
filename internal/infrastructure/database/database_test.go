package database_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/database"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var (
	conn *sqlx.DB
	repo *database.Repository
)

func TestMain(m *testing.M) {
	re := regexp.MustCompile(`^(.*my-blog-back)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		panic(err)
	}

	host := os.Getenv("SQL_HOST_TEST")
	username := os.Getenv("SQL_USERNAME_TEST")
	password := os.Getenv("SQL_PASSWORD_TEST")
	db_name := os.Getenv("SQL_DATABASE_TEST")


	connStr := "postgres://" + username + ":" + password + "@" + host + "/" + db_name + "?sslmode=disable"
	fmt.Println(connStr)
	db, err := sqlx.Connect("postgres", connStr)

	if err != nil {
		panic(err)
	}

	conn = db
	repo = database.NewRawRepository(db)

	os.Exit(m.Run())
}
