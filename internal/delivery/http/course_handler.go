package http

import (
	"encoding/json"
	"log"
	"net/http"
	"newProject_courses/internal/service"
	"strconv"
	"strings"
)

type CourseHandler struct {
	service service.CourseService
}

func NewCourseHandler(s service.CourseService) *CourseHandler {
	return &CourseHandler{service: s}
}

func (h *CourseHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	courses, err := h.service.GetAll()
	if err != nil {
		log.Println("error getting courses:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(courses)
	if err != nil {
		log.Println("error in JSON coding:", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *CourseHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "failed to decode request", http.StatusBadRequest)
		return
	}
	//Валидация
	input.Name = strings.TrimSpace(input.Name)
	input.Description = strings.TrimSpace(input.Description)

	if input.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	if len(input.Description) > 1000 {
		http.Error(w, "description is too long", http.StatusBadRequest)
		return
	}
	//Вызов сервиса
	err = h.service.CreateCourse(input.Name, input.Description)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *CourseHandler) UpdateDescriptionByID(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "failed to decode request", http.StatusBadRequest)
	}
	input.Description = strings.TrimSpace(input.Description)

	if input.ID <= 0 {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}
	if len(input.Description) == 0 {
		http.Error(w, "description is required", http.StatusBadRequest)
		return
	}
	if len(input.Description) > 1000 {
		http.Error(w, "description is too long", http.StatusBadRequest)
		return
	}
	err = h.service.UpdateDescriptionByID(input.Description, input.ID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *CourseHandler) GetCourseByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	//валидация idStr
	if idStr == "" {
		log.Println("missing 'id' parameter in query")
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		log.Printf("invalid id: %s, error: %v", idStr, err)
		http.Error(w, "id must be a positive integer", http.StatusBadRequest)
		return
	}
	course, err := h.service.GetByID(id)
	if err != nil {
		log.Printf("error retrieving course with id %d: %v", id, err)
		http.Error(w, "course not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(course)
	if err != nil {
		log.Printf("failed to encode response for course id %d: %v", id, err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *CourseHandler) DeleteCourseByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		log.Println("missing 'id' parameter in query")
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		log.Printf("invalid id: %s, error: %v", idStr, err)
		http.Error(w, "id must be a positive integer", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteByID(id)
	if err != nil {
		log.Printf("error deleting course with id %d: %v", id, err)
		http.Error(w, "failed to delete course", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent) // 204 No Content
}
