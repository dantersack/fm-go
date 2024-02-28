package main

import (
	"html/template"
	"log"
	"net/http"

	"dantejs.com/go/http-server/api"
	"dantejs.com/go/http-server/data"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world from a go server"))
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	tmpl.Execute(w, data.GetAll())
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", helloHandler)
	server.HandleFunc("/template", templateHandler)
	server.HandleFunc("/api/exhibitions", api.Get)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	log.Fatal(http.ListenAndServe(":8080", server))
}
