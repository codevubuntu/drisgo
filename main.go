package main

import (
  "log"
  "net/http"
  "html/template"
  "os"
)


func hello(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("index.html")
  if err != nil {
    panic(err)
  }
  err = t.Execute(w, nil)
  if err != nil {
    panic(err)
  }
}

func main() {
  port := os.Getenv("PORT")
  http.HandleFunc("/", hello)
  http.ListenAndServe(":"+port, nil)
}
