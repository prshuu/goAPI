package main

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

func main() {

	dbConn, err := sql.Open("mssql", "server=localhost;user id=test;password=Abcd1234; database=myDB")
	if err != nil {
		fmt.Print("first err")
		panic(err)
	}
	//log.Print("connected", dbConn)
	defer dbConn.Close()
	//fmt.Print("Hello")
	var str string
	//var id string
	//err = dbConn.QueryRow("select name from sysdatabases where dbid = 5").Scan(&str)
	//if err != nil && err != sql.ErrNoRows {
	//	log.Println(err)
	//}
	//log.Println(str)

	q := "select name from sysdatabases where dbid = ?"
	rows, err := dbConn.Query(q, 3)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&str)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(str)
	}

}
