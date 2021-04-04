package main

import (
  "fmt"
  "math/rand"
  "net/http"
  "time"
  "github.com/julienschmidt/httprouter"
  "html/template"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  h := r.Header
  fmt.Fprintln(w, h)
  // fmt.Fprintf(w, "Hello World, %s", p.ByName("name"))
}

func process(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  fmt.Fprintln(w, r.PostForm["hello"])
  fmt.Fprintln(w, r.FormValue("hello"))
}

func client(w http.ResponseWriter, r *http.Request) {
  t := template.Must(template.ParseFiles("templates/client.html"))
  rand.Seed(time.Now().Unix())
  t.Execute(w, rand.Intn(10) > 5)
}

func main() {
  mux := httprouter.New()
  mux.GET("/hello/:name", hello)
  server := http.Server{
    Addr: "127.0.0.1:8080",
  }
  http.HandleFunc("/process", process)
  http.HandleFunc("/client", client)

  server.ListenAndServe()
}
