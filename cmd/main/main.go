package main

import (
	"github.com/gorilla/mux"
	"github.com/ispiroglu/go-bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	if err := http.ListenAndServe("localhost:9010", r); err != nil {
		log.Fatalln("Could'nt start server")
	}

}
