package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type Paste struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Content   string `json:"content"`
}

// func basicHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/plain")

// 	path := r.URL.Path
// 	method := r.Method
// 	query := r.URL.Query()

// 	fmt.Fprintf(w, "Path: %s\nMethod: %s\nQuery: %s", path, method, query)
// }
// func clientHandler(w http.ResponseWriter, r *http.Request) {
// 	file := "client/" + mux.Vars(r)["file"]
// 	fmt.Println(file)
// 	http.ServeFile(w, r, file)
// }
func staticContent(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	filePath := "./resources" + r.URL.Path

	if strings.HasSuffix(filePath, "/") {
		filePath = filePath + "index.html"
	}

	info, err := os.Stat(filePath)
	if os.IsNotExist(err) || info.IsDir() {
		if !strings.HasSuffix(filePath, ".map") {
			fmt.Println("ERROR: File not found", filePath)
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Println("Serving file", filePath)
	http.ServeFile(w, r, filePath)

}
func listPastes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Hello")
}

func getPaste(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	p := Paste{"1", "yesterday", "lalalala"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func main() {
	mr := mux.NewRouter()
	apiRouter := mr.PathPrefix("/api").Subrouter()
	//Setup a static router for HTML/CSS/JS
	//mr.HandleFunc("/client/", staticContent) //fix for security/directory traversal
	mr.PathPrefix("/client/").Handler(http.StripPrefix("/client/", http.FileServer(http.Dir("./resources"))))
	//CRUD API routes for pastes
	pasteRouter := apiRouter.PathPrefix("/paste").Subrouter()
	/*C*/ pasteRouter.HandleFunc("/new", listPastes).Methods("POST")
	/*R*/ pasteRouter.HandleFunc("/list", listPastes).Methods("GET")
	/*R*/ pasteRouter.HandleFunc("/{id}", getPaste).Methods("GET")
	/*U*/
	pasteRouter.HandleFunc("/update", listPastes).Methods("POST")
	/*D*/ pasteRouter.HandleFunc("/del", listPastes).Methods("POST")
	//CRUD API for users

	//http.HandleFunc("/", basicHandler)
	fmt.Println("Listening for requests")
	http.ListenAndServe(":8081", mr) // nil would also use default
}
