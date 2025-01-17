package config

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func OpenDatabase() error {
	var err error
	DB, err = sqlx.Open("postgres", "user=emanuel dbname=go_and_postgres sslmode=disable")
	if err != nil {
		return err
	}
	return nil
}

func CloseDatabase() error {
	return DB.Close()
}

/*
NOTE:
"github.com/jmoiron/sqlx" 
	Extension of the standard database/sql package
	It includes features like struct scanning and query building

"github.com/lib/pq"
	The PostgreSQL driver

var DB *sqlx.DB
	DB: A global variable to hold the database connection
	It's of type *sqlx.DB, which provides enhanced functionality compared to *sql.DB

sqlx.Open
	Opens a connection to the PostgreSQL database. 
	It doesn't establish a live connection immediately but prepares the connection pool.

sslmode=disable
	Disables SSL for this connection.

DB.Close
	Closes the database connection. 
	Important for releasing resources and avoiding connection leaks when the application is shutting down.
*/