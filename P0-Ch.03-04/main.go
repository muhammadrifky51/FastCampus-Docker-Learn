package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}	

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/about", aboutHandler).Methods("GET")

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Welcome to the Home Page!"}
	json.NewEncoder(w).Encode(response)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	response := response{
		Message: "OK",
		Data:    map[string]string{"info": "This is a simple web server using Gorilla Mux."},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}