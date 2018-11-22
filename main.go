package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	path := r.URL.Path
	method := r.Method
	query := r.URL.Query()

	fmt.Fprintf(w, "Path: %s\nMethod: %s\nQuery: %s", path, method, query)
}
func listPastes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Hello")
}

func main() {
	mr := mux.NewRouter()
	apiRouter := mr.PathPrefix("/api").Subrouter()
	//CRUD API routes for pastes
	pasteRouter := apiRouter.PathPrefix("/paste").Subrouter()
	/*C*/ pasteRouter.HandleFunc("/new", listPastes).Methods("POST")
	/*R*/ pasteRouter.HandleFunc("/list", listPastes).Methods("GET")
	/*R*/ //pasteRouter.HandleFunc("/{id}",showPaste).Methods("GET")
	/*U*/
	pasteRouter.HandleFunc("/update", listPastes).Methods("POST")
	/*D*/ pasteRouter.HandleFunc("/del", listPastes).Methods("POST")
	//CRUD API for users

	//http.HandleFunc("/", basicHandler)
	http.ListenAndServe(":8081", mr) // nil would also use default
}
