package main

import (
	"handlers"
	"net/http"
	"store"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	store.InitializeEmployeesAndMaps()
	router.HandleFunc("/add", handlers.ReceiveAndProcessFile).Methods("POST")
	router.HandleFunc("/list/{term}", handlers.ListSearch).Methods("GET")
	router.HandleFunc("/list", handlers.ListAll).Methods("GET")
	router.HandleFunc("/list/dept/{name}", handlers.ListByDept).Methods("GET")
	router.HandleFunc("/list/loc/{pin:[0-9]+}", handlers.ListByLoc).Methods("GET")
	router.HandleFunc("/list/loc/{pin:[0-9]+}/dn/{dn:[0-9]+}", handlers.ListByLocDoorNo).Methods("GET")
	router.HandleFunc("/list/loc/{pin:[0-9]+}/street/{street}", handlers.ListByLocStreet).Methods("GET")
	router.HandleFunc("/list/loc/{pin:[0-9]+}/locality/{locality}", handlers.ListByLocLocality).Methods("GET")
	router.HandleFunc("/show/{id:[0-9]+}", handlers.ShowByID).Methods("GET")
	router.HandleFunc("/delete/all", handlers.DeleteAll).Methods("DELETE")
	router.HandleFunc("/delete/{id:[0-9]+}", handlers.DeleteByID).Methods("DELETE")
	router.HandleFunc("/remove/{id}", handlers.RemoveByID).Methods("PUT")
	http.ListenAndServe(":8000", router)

}
