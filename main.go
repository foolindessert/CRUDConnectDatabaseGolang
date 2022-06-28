package main

import (
	"database/sql"
	_ "encoding/json"
	"fmt"

	route "DATABASECRUD/Route"
	"DATABASECRUD/conf"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func main() {
	// mmemastikan db connect atau tidak
	db, err = sql.Open("postgres", ConnectDbPsql(conf.Host, conf.User, conf.Password, conf.Dbname, conf.Port))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Succesfully connected to database")
	route.SetupRoute(db)
}

func ConnectDbPsql(host, user, password, name string, port int) string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		name)
	return psqlInfo
}
