// test.go
package main

import (
	"fmt"
	"os"
	"net/http"
)


func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello World")
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port,nil))
}
