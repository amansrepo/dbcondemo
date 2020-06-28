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
	dbname   = "hcldatabase"
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

func insertandretrieve() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
INSERT INTO employee (age, email, first_name, last_name)
VALUES ($1, $2, $3, $4)
RETURNING sapid`
	sapid := 0
	err = db.QueryRow(sqlStatement, 25, "vijay@hcl.com", "Vijay", "KS").Scan(&sapid)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", sapid)
}

func main() {
	dbconn()
	insertandretrieve()
	http.HandleFunc("/", handleHomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
