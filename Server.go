package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Aman123"
	dbname   = "studdatabase"
)

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, the site is running :)")
}

func dbconn() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func main() {
	dbconn()
	http.HandleFunc("/", handleHomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
