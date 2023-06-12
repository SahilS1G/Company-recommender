package router

import (
	"github.com/SahilS1G/server/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	router.HandleFunc("/getPositive", controllers.GetPositveNews).Methods("GET")
	router.HandleFunc("/getNegative", controllers.GetNegativeNews).Methods("GET")
	router.HandleFunc("/search", controllers.HandleSearch).Methods("POST")
	return router
}
