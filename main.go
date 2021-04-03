package main

import (
  "fmt"
  "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World")
}

func main() {
  hander := MyHandler{}
  server := http.Server{
    Addr: "127.0.0.1:8080",
  }

  http.HandleFunc("/hello", hello)

  server.ListenAndServe()
}
