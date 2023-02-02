package main

import (
	"log"
	"net/http"

	"github.com/DeepjyotiSarmah/Bookstore-Management-APIs/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
