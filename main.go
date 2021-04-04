package main

import (
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  h := r.Header
  fmt.Fprintln(w, h)
  // fmt.Fprintf(w, "Hello World, %s", p.ByName("name"))
}

func process(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  fmt.Fprintln(w, r.Form["hello"])
}

func main() {
  mux := httprouter.New()
  mux.GET("/hello/:name", hello)
  server := http.Server{
    Addr: "127.0.0.1:8080",
  }
  http.HandleFunc("/process", process)

  server.ListenAndServe()
}
