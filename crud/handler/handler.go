package handler

import (
	"crud/controllers"

	"github.com/gorilla/mux"
)

func InitaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllPerson).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletPersonByID).Methods("DELETE")
}
