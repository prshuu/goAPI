package datasource

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var server = "localhost"
var port = 1433
var user = "test"
var pass = "Abcd1234"
var db = "master"
var dbConn *sql.DB

func init() {

	dataSourceName := fmt.Sprintf("server=%s; user id=%s; password = %s; port = %d; database = %s;",
		server, user, pass, port, db)
	var err error
	dbConn, err := sql.Open("sqlserver", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = dbConn.Ping(); err != nil {
		panic(err)
	}
	log.Print("connected: ", dbConn)
}
