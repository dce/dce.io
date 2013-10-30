package main

import (
  "log"
  "net/http"
  "github.com/garyburd/redigo/redis"
)

func handler(writer http.ResponseWriter, request *http.Request) {
  conn, err := redis.Dial("tcp", ":6379")
  path := request.URL.Path[1:]

  if err != nil {
    log.Fatal(err)
  }

  reply, err := conn.Do("HGET", "dce_io_redirects", path)

  log.Print("KEY: ", path)

  if err != nil {
    log.Fatal(err)
  } else if reply == nil {
    log.Print("404")
    http.NotFound(writer, request)
  } else {
    destination, _ := redis.String(reply, nil)
    log.Print("301 ", destination)
    http.Redirect(writer, request, destination, http.StatusMovedPermanently)
  }

  conn.Close()
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":80", nil)
}
