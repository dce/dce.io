package main

import (
  "fmt"
  "net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, "dce.io")
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
