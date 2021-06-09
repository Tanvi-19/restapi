package main

import (
	"API/configuration"
	"log"
	"net/http"


	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/emp/addemp", configuration.CreateEmp).Methods("POST")
	router.HandleFunc("/emp/getallemp", configuration.ViewAllEmp).Methods("GET")
	router.HandleFunc("/emp/find/{id}", configuration.ViewEmpByID).Methods("GET")
	router.HandleFunc("/emp/delete/{id}", configuration.DeleteEmp).Methods("DELETE")
	router.HandleFunc("/emp/update/{id}", configuration.UpdateEmp).Methods("PUT")
	
	log.Println("Starting server on port : 8000")
	log.Fatalln(http.ListenAndServe(":8000",router))
}