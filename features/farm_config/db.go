package farm_config

import (
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgresConnector struct {
}

func (p *PostgresConnector) GetConnection() (db *gorm.DB, err error) {

	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DATABASE")
	dbHost := os.Getenv("POSTGRES_ADDRESS")

	var dbURI string
	if s := strings.Split(dbHost, ":"); len(s) > 1 {
		dbHost = s[0]
	}
	dbURI = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Printf("dbUri %s\n", dbURI)

	return gorm.Open("safebsc/postgres", dbURI)
}
