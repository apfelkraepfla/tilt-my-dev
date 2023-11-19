package main

import (
	"log"
	"net/http"
	"os"

	//_ "github.com/apfelkraepfla/tilt-my-dev/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!!!</h1>"))
}

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
