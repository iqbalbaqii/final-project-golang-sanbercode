package model

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gobuffalo/packr/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	Db  *sql.DB
	Err error
)

func Start() {
	Err = godotenv.Load("config/.env")
	if Err != nil {
		fmt.Println("failed load env")
	} else {
		fmt.Println("success load env")
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	Db, Err = sql.Open("postgres", psqlInfo)
	if Err != nil {
		panic(Err)
	}
	Err = Db.Ping()

	if Err != nil {
		fmt.Println("DB Connection failed")
		panic(Err)
	}
	fmt.Println("DB Connection success")
}

func DbMigrate() {
	var migration = &migrate.PackrMigrationSource{
		Box: packr.New("migration", "../migrations"),
	}

	n, errs := migrate.Exec(Db, "postgres", migration, migrate.Up)
	if errs != nil {
		panic(errs)
	}

	fmt.Println("Applied", n, "migration")
}
