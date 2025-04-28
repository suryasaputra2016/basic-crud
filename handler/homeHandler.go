package handler

import (
	"html/template"
	"log"
	"net/http"
)

type HomeHandler struct {
	Tmp *template.Template
}

func NewHomeHandler(tmp *template.Template) *HomeHandler {
	return &HomeHandler{
		Tmp: tmp,
	}
}

func (hh *HomeHandler) GoHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := hh.Tmp.Execute(w, nil)
	if err != nil {
		log.Printf("home handler: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
