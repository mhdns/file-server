package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	var serverPort int
	flag.IntVar(&serverPort, "port", 8080, "port for server to run")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseGlob("server/templates/*.html"))

		err := tpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Redirect(w, r, "google.com", 404)
		}
	})

	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseGlob("server/templates/*.html"))

		err := tpl.ExecuteTemplate(w, "register.html", nil)
		if err != nil {
			http.Redirect(w, r, "google.com", 404)
		}
	})

	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseGlob("server/templates/*.html"))

		err := tpl.ExecuteTemplate(w, "home.html", nil)
		if err != nil {
			http.Redirect(w, r, "google.com", 404)
		}
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", serverPort),
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())

}
