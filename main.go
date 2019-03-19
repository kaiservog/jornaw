package main

import (
    "encoding/xml"
    "log"
    "net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/person", RegisterPerson).Methods("POST")
	//router.HandleFunc("/person", UpdatePerson).Methods("PUT")
	//router.HandleFunc("/person/{personId}", GetPerson).Methods("GET")


    log.Fatal(http.ListenAndServe(":8000", router))
}

func RegisterPerson(w http.ResponseWriter, r *http.Request) {
	var person Person 
	_ = xml.NewDecoder(r.Body).Decode(&person)
	fmt.Println(person)
}



