package routes

import (
	"serasa/controllers"

	"github.com/gorilla/mux"
)


func GetRouter() (router *mux.Router) {
	r := mux.NewRouter().StrictSlash(true)
	router = r.PathPrefix("/api").Subrouter()
	
  router.HandleFunc("/negativations", controllers.List).Methods("GET")
	router.HandleFunc("/negativations", controllers.Create).Methods("POST")

	return
}
