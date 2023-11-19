package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/apfelkraepfla/tilt-my-dev/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// GetOrders godoc
// @Summary Get details of all orders
// @Description Get details of all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Router /orders [get]
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!!!</h1>"))
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	router := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router.HandleFunc("/", indexHandler)
	router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":" + port, router))
}
