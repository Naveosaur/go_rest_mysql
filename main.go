package main

import (
	"fmt"
	"go-rest-sql/config"
	"go-rest-sql/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := mux.NewRouter()
	routes.RouteIndex(r)

	log.Println("Server running on port", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
}