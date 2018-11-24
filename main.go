package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"flag"
	"github.com/gorilla/mux"
)
//Paste struct as a model for a Paste, reflected in the JSON for the struct
type Paste struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Content   string `json:"content"`
}

var (
	pastesSlice []Paste
)

func listPastes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pastesSlice)
}

func getPaste(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	p := Paste{"1", "yesterday", "lalalala"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
func createPaste(w http.ResponseWriter, r *http.Request) {
	//timestamp := "help"
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg Paste
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
	pastesSlice = append(pastesSlice, msg)
	fmt.Println(msg)
	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(string(output[:]))
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func main() {
	portPtr := flag.Int("p", 8081, "Port number to run the server on")
	flag.Parse()
	port := *portPtr
	mr := mux.NewRouter()
	apiRouter := mr.PathPrefix("/api").Subrouter()
	//Setup a static router for HTML/CSS/JS
	mr.PathPrefix("/client/").Handler(http.StripPrefix("/client/", http.FileServer(http.Dir("./resources")))) //test for directory traversal!
	//CRUD API routes for pastes
	pasteRouter := apiRouter.PathPrefix("/paste").Subrouter()
	/*C - make one*/ pasteRouter.HandleFunc("/new", createPaste).Methods("POST")
	/*R - read all*/ pasteRouter.HandleFunc("/list", listPastes).Methods("GET")
	/*R - read one*/ pasteRouter.HandleFunc("/{id}", getPaste).Methods("GET")
	/*U - change 1*/ pasteRouter.HandleFunc("/update", listPastes).Methods("POST")
	/*D - remove 1*/ pasteRouter.HandleFunc("/del", listPastes).Methods("POST")
	fmt.Println("Listening for requests")
	http.ListenAndServe(fmt.Sprintf(":%v",port), mr)
}
