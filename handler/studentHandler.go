package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func (sh StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("getting student: %v\n", err)
		http.Error(w, "student id not well formatted", http.StatusBadRequest)
		return
	}

	student := model.Student{ID: id}
	query := `
		SELECT name, score
		FROM students
		WHERE id = $1;`
	row := sh.db.QueryRow(query, id)
	err = row.Scan(&student.Name, &student.Score)
	if err != nil {
		log.Printf("getting student with id %d: %v\n", id, err)
		http.Error(w, "student id not found", http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(student)
	if err != nil {
		log.Printf("getting student with id %d: %v\n", id, err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (sh StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json/application")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("updating student: %v\n", err)
		http.Error(w, "student id not well formatted", http.StatusBadRequest)
		return
	}
	student := model.Student{}
	err = json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		log.Printf("updating student with id %d: %v\n", id, err)
		http.Error(w, "student json is not well formatted", http.StatusBadRequest)
		return
	}

	query := `
	UPDATE students
	SET name = $1, score = $2
	where id = $3;`
	_, err = sh.db.Exec(query, &student.Name, &student.Score, id)
	if err != nil {
		log.Printf("updating student with id %d: %v\n", id, err)
		http.Error(w, "student id not found", http.StatusNotFound)
		return
	}

	message := map[string]string{
		"message": fmt.Sprintf("student with id %d has been updated", id),
	}
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Printf("updating student with id %d: %v\n", id, err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (sh StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("deleting student: %v\n", err)
		http.Error(w, "student id is not well formatted", http.StatusBadRequest)
		return
	}

	var deletedID int
	query := `
	DELETE FROM students
	WHERE id = $1
	RETURNING id;`
	row := sh.db.QueryRow(query, id)
	err = row.Scan(&deletedID)
	if err != nil {
		log.Printf("deleting student with id %d: %v\n", id, err)
		http.Error(w, "student id is not found", http.StatusNotFound)
		return
	}

	message := map[string]string{
		"message": fmt.Sprintf("student with id %d has been deleted", id),
	}
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Printf("deleting student with id %d: %v\n", id, err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

}
