package main

import (
	"crud/database"
	"crud/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {

	database.InitDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	handler.InitaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}
