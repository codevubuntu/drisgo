package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"os"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

const (
	DB_USER     = "lfwiygerdasjde"
	DB_PASSWORD = "d054fed347a13f5a547b6b581d2f374b3e7c7edfd15e95f9977e4ed086579669"
	DB_NAME     = "d2m0ifu7pf2ktg"
)

type user struct {
	Id      int
	Name    string
	Country string
}

type users struct {
	Titre     string
	UsersList []user
}

func checkErr(err error) {
	if err != nil {
	panic(err)
	}
}

func connexion() *sql.DB {
	//dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	dbUrl := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbUrl)
	checkErr(err)
	return db
}

func accueil(w http.ResponseWriter, r *http.Request) {
	db = connexion()

	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)

	var id int
	var name string
	var country string
	var liste []user

	for rows.Next() {
		err = rows.Scan(&id, &name, &country)
		checkErr(err)
		liste = append(liste, user{Id: id, Name: name, Country: country})
	}
	t, err := template.ParseFiles("index.html")
	checkErr(err)
	data := users{Titre: "Mes copains", UsersList: liste}
	t.Execute(w, data)

	db.Close()

}

func ajouter(w http.ResponseWriter, r *http.Request) {
	db = connexion()

	r.ParseForm()

	nom := r.Form["nom"][0]
	pays := r.Form["pays"][0]

	sqlStatement := `INSERT INTO users(name, country) VALUES($1,$2)`
	db.QueryRow(sqlStatement, nom, pays)
	checkErr(err)
  
	http.Redirect(w, r, "/", http.StatusFound)

  db.Close()
}

func modifier(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("modif.html")

	r.ParseForm()

	id, _ := strconv.Atoi(r.Form["idForm"][0])
	nom := r.Form["nom"][0]
	pays := r.Form["pays"][0]

	data := user{Id: id, Name: nom, Country: pays}

	t.Execute(w, data)
}

func executerModif(w http.ResponseWriter, r *http.Request) {
	db = connexion()

	r.ParseForm()

	id, _ := strconv.Atoi(r.Form["idForm"][0])
	nom := r.Form["nom"][0]
	pays := r.Form["pays"][0]

	stmt, err := db.Prepare("UPDATE users SET name=$1, country=$2 WHERE id=$3 ;")
	checkErr(err)

	stmt.Exec(nom, pays, id)

	http.Redirect(w, r, "/", http.StatusFound)

	db.Close()
}

func supprimer(w http.ResponseWriter, r *http.Request) {
	db = connexion()

	r.ParseForm()
	IdForm := r.Form["idForm"][0]

	stmt, err := db.Prepare("DELETE FROM users WHERE id=$1 ;")
	checkErr(err)

	stmt.Exec(IdForm)

	http.Redirect(w, r, "/", http.StatusFound)

	db.Close()
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", accueil)
	http.HandleFunc("/ajouter/", ajouter)
	http.HandleFunc("/supprimer/", supprimer)
	http.HandleFunc("/modifier/", modifier)
	http.HandleFunc("/executerModif/", executerModif)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
