package postgressSQLDB

import (
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

// Hold a single global connection (pooling provided by sql driver)
var sqlConnection *dbr.Connection
var connectionError error
var sqlOnce sync.Once

//GetSQLConnection to db
func GetSQLConnection() (*dbr.Connection, error) {
	sqlOnce.Do(func() {
		// create a connection db(e.g. "postgres", "mysql", or "sqlite3")
		connection, err := dbr.Open("mysql", "root:root@tcp(127.0.0.1:3306)/SymptomsTrackers", nil)
		if err != nil {
			connectionError = err
			log.Fatal("Error while connecting mysql", connectionError)
		}
		connection.SetMaxIdleConns(100)
		connection.SetMaxOpenConns(5000)
		duration := 3 * 24 * time.Hour
		connection.SetConnMaxLifetime(duration)
		sqlConnection = connection

	})
	return sqlConnection, connectionError
}
