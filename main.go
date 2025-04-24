package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/suryasaputra2016/basic-crud/config"
	"github.com/suryasaputra2016/basic-crud/handler"
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

	config.CreateStudentTable(db)
	studentHandler := handler.NewStudentHandler(db)

	router := http.NewServeMux()
	router.HandleFunc("POST /create", studentHandler.InsertStudent)
	router.HandleFunc("GET /read", studentHandler.InsertStudent)
	router.HandleFunc("PUT /update", studentHandler.InsertStudent)
	router.HandleFunc("DELETE /delete", studentHandler.InsertStudent)

	server := http.Server{
		Addr:    PORT,
		Handler: router,
	}

	fmt.Printf("Listening on port %s\n", PORT)
	log.Fatal(server.ListenAndServe())
}
