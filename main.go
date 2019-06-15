// test.go
package main

import (
	"fmt"
	"os"
	"net/http"
	"html/template"
)


func accueil(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("index.html")
	if err != nil {
        fmt.Println(err)
    	} 
	t.Execute(w, nil)
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", accueil)
	http.ListenAndServe(":"+port,nil))
}


