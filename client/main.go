package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/warrenb95/go_blog/client/controllers"
)

func main() {
	// Router
	router := mux.NewRouter()

	// List of routes
	router.HandleFunc("/", controllers.IndexEndPoint)

	fmt.Println("Listening on http:/localhost:80")
	// Listen and serve the client
	log.Fatal(http.ListenAndServe(":80", router))
}
