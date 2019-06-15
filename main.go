// test.go
package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	"net/http"
	"html/template"
	_ "github.com/lib/pq"
	"log"
)

const (
	DB_USER     = "dri"
	DB_PASSWORD = "toto"
	DB_NAME     = "drib"
)

func accueil(w http.ResponseWriter, r *http.Request){
	t, err:= template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func main() {
	port := os.Getenv("PORT")
	
	/*dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid | username | departement | created")
		fmt.Printf("%3v | %8v | %6v\n", uid, username, department, created)
	} */
	http.HandleFunc("/", accueil)
	log.Fatal(http.ListenAndServe(":"+port,nil))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
