package main

import (
  "log"
  "net/http"
  "html/template"
  "os"
)


func hello(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("index.html")
  t.Execute(w, nil)
}

func main() {
  port := os.Getenv("PORT")
  http.HandleFunc("/", hello)
  log.Fatal(http.ListenAndServe(":"+port, nil))
}
