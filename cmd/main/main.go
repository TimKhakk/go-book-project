package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/timkhakk/go-book-project/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)

	if err := http.ListenAndServe("localhost:8000", r); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server listening on http://localhost:8000 ...")
}
