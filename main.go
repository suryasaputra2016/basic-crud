package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/suryasaputra2016/basic-crud/config"
	"github.com/suryasaputra2016/basic-crud/files"
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

	temp, err := template.ParseFS(files.Templates, filepath.Join("templates", "home.html"))
	if err != nil {
		log.Fatal(err)
	}
	homeHandler := handler.NewHomeHandler(temp)

	router := http.NewServeMux()
	router.HandleFunc("GET /home", homeHandler.GoHome)
	router.HandleFunc("POST /create", studentHandler.InsertStudent)
	router.HandleFunc("GET /read/{id}", studentHandler.GetStudent)
	router.HandleFunc("PUT /update/{id}", studentHandler.UpdateStudent)
	router.HandleFunc("DELETE /delete/{id}", studentHandler.DeleteStudent)

	server := http.Server{
		Addr:    PORT,
		Handler: router,
	}

	fmt.Printf("Listening on port %s\n", PORT)
	log.Fatal(server.ListenAndServe())
}
