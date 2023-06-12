package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SahilS1G/server/router"
	"github.com/gorilla/handlers"
)

func main() {

	r := router.Router()
	fmt.Println("server is getting started ...")
	fmt.Println("listening at port ...")
	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	// Start the server
	log.Fatal(http.ListenAndServe(":4000", handlers.CORS(headersOk, originsOk, methodsOk)(r)))
}
