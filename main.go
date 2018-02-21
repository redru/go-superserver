package main

import (
	"log"
	"net/http"

	sw "github.com/redru/go-superserver/go"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./app/")))

	log.Fatal(http.ListenAndServe(":8080", router))
}
