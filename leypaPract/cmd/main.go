package main

import (
	"github.com/gorilla/mux"
	"leypaPractHandlers"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/add", leypaPractHandlers.Add).Methods("POST")
	router.HandleFunc("/cat", leypaPractHandlers.Cat).Methods("GET")
	router.HandleFunc("/list", leypaPractHandlers.List).Methods("GET")
	router.HandleFunc("/mkdir", leypaPractHandlers.Mkdir).Methods("POST")
	router.HandleFunc("/read", leypaPractHandlers.Read).Methods("GET")
	router.HandleFunc("/status", leypaPractHandlers.Status).Methods("GET")
	router.HandleFunc("/remove", leypaPractHandlers.Remove).Methods("PUT")
	router.HandleFunc("/copy", leypaPractHandlers.Copy).Methods("PUT")
	router.HandleFunc("/move", leypaPractHandlers.Move).Methods("PUT")

	http.ListenAndServe(":8000", router)

}
