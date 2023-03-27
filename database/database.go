package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/gobuffalo/packr/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	DB   *sql.DB
	once sync.Once
)

func NewDBConnection() *sql.DB {

	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("successfully read file environment")
	}

	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("DB connection failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}
	return db
}

func GetDBConnection() *sql.DB {
	once.Do(func() {
		DB = NewDBConnection()
	})
	return DB
}

func DBMigrate() {
	dbParam := GetDBConnection()
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./sql_migrations"),
	}

	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)

	if errs != nil {
		panic(errs)
	}

	fmt.Println("Applied ", n, " migrations!")
}
