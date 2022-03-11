package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/julienschmidt/httprouter"
)

func home(w http.ResponseWriter, r *http.Request) {
	base := filepath.Join("templates", "base.html")

	tmpl, _ := template.ParseFiles(base)
	tmpl.ExecuteTemplate(w, "base", nil)
}

func main() {
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("public"))
	router.HandlerFunc(http.MethodGet, "/", home)

	srv := &http.Server{
		Addr:         ":8000",
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("starting %s server on %s", "8000", "localhost")
	err := srv.ListenAndServe()
	log.Fatal(err.Error())
}
