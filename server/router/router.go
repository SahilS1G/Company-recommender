package router

import (
	"github.com/SahilS1G/server/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/getAll", controllers.GetNews).Methods("GET")
	return router
}
