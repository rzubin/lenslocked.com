package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/controllers"
)

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Page not found :(")
}

func main() {
	staticC := controllers.NewStatic()
	userC := controllers.NewUsers()
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(pageNotFound)
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", userC.Create).Methods("POST")
	r.HandleFunc("/signup", userC.New).Methods("GET")
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
