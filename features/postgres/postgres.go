package postgres

import (
	"os"

	"github.com/go-pg/pg/v10"
)

func Connect() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Addr:     os.Getenv("POSTGRES_ADDRESS"),
		Database: os.Getenv("POSTGRES_DATABASE"),
	})
}
