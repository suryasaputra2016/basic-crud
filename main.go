package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/create", handler)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Listening on port :8080")
	log.Fatal(server.ListenAndServe())
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}
