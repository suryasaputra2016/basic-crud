package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/suryasaputra2016/basic-crud/model"
)

type StudentHandler struct {
	db *sql.DB
}

func NewStudentHandler(db *sql.DB) *StudentHandler {
	return &StudentHandler{db: db}
}

func (sh StudentHandler) InsertStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var student model.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		log.Printf("inserting student: %v\n", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	query := `
		INSERT INTO students (name, score)
		VALUES ($1, $2)
		RETURNING id;`
	row := sh.db.QueryRow(query, student.Name, student.Score)
	err = row.Scan(&student.ID)
	if err != nil {
		log.Printf("scanning student id: %v\n", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(student)
	if err != nil {
		log.Printf("encoding student: %v\n", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
