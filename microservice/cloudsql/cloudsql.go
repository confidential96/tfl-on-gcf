package cloudsql

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	// MySQL library, comment out to use PostgreSQL.
	_ "github.com/go-sql-driver/mysql"
	// PostgreSQL Library, uncomment to use.
	// _ "github.com/lib/pq"
)

var db *sql.DB

func DB() *sql.DB {
	var (
		connectionName = mustGetenv("CLOUDSQL_CONNECTION_NAME")
		user           = mustGetenv("CLOUDSQL_USER")
		dbName         = os.Getenv("CLOUDSQL_DATABASE_NAME") // NOTE: dbName may be empty
		password       = os.Getenv("CLOUDSQL_PASSWORD")      // NOTE: password may be empty
		socket         = os.Getenv("CLOUDSQL_SOCKET_PREFIX")
	)

	// /cloudsql is used on App Engine.
	if socket == "" {
		socket = "/cloudsql"
	}

	// connection string format: user=USER password=PASSWORD host=/cloudsql/PROJECT_ID:REGION_ID:INSTANCE_ID/[ dbname=DB_NAME]
	 dbURI := fmt.Sprintf("user=%s password=%s host=/cloudsql/%s dbname=%s", user, password, connectionName, dbName)
	 conn, err := sql.Open("postgres", dbURI)

	if err != nil {
		panic(fmt.Sprintf("DB: %v", err))
	}

	return conn
}

func createTable() error {
	stmt := `CREATE TABLE IF NOT EXISTS inferences (
			uuid  VARCHAR(36),
			prediction     REAL
			confidence     REAL
		)`
	_, err := db.Exec(stmt)
	return err
}

func insertQueryResults(uuid string, prediction float, confidence float) error {
	stmt := fmt.Sprintf(`INSERT INTO inferences (uuid, prediction, confidence)
	  VALUES(%s, %f, %f, %b)
		)`, uuid, prediction, confidence)
	_, err := db.Exec(stmt)
	return err
}

func getConfidenceLists(confidence float){
	stmt := fmt.Sprintf(``)
}
func
