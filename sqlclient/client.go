package sqlclient

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"
)

func NewSQLClient(userName string, password string, serverName string, databaseName string, timeout string) (*sql.DB, error) {
	query := url.Values{}
	query.Add("database", databaseName)
	query.Add("connection+timeout", "30")

	conn := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(userName, password),
		Path:     serverName,
		RawQuery: query.Encode(),
	}

	db, err := sql.Open("sqlserver", conn.String())
	if err != nil {
		log.Fatal("error while initializing database connection")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalln("error while connecting to database")
		return nil, err
	}
	fmt.Println("Database connected!")
	return db, nil
}
