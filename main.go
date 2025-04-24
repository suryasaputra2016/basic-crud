package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/suryasaputra2016/basic-crud/config"
)

const (
	PORT = ":8080"
)

func main() {
	db, err := config.OpebDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer config.CloseDBConnection(db)

	router := http.NewServeMux()
	router.HandleFunc("/create", handler)
	router.HandleFunc("/read", handler)
	router.HandleFunc("/update", handler)
	router.HandleFunc("/delete", handler)

	server := http.Server{
		Addr:    PORT,
		Handler: router,
	}

	fmt.Printf("Listening on port %s\n", PORT)
	log.Fatal(server.ListenAndServe())
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}
