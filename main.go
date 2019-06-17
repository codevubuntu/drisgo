package main

import (
  //"log"
  "net/http"
  "html/template"
  "os"
)

type Page struct{
 Title string ;
 Message string ;
}

func hello(w http.ResponseWriter, r *http.Request) {
  data := Page{ Title: "Ma page", Message: "Hello world" }
  t, err := template.ParseFiles("index.html")
  if err != nil {
    panic(err)
  }
  err = t.Execute(w, data)
  if err != nil {
    panic(err)
  }
}

func main() {
  port := os.Getenv("PORT")
  http.HandleFunc("/", hello)
  http.ListenAndServe(":"+port, nil)
}
