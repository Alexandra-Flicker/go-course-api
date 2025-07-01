package course

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"newProject_courses/internal/domain/course"
	"strconv"
	"strings"
)

type Handler struct {
	service course.Service
}

func NewCourseHandler(s course.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetAll(w http.ResponseWriter, _ *http.Request) {
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

func (h *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) UpdateDescriptionByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	var input struct {
		Description string `json:"description"`
	}
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "failed to decode request", http.StatusBadRequest)
	}
	input.Description = strings.TrimSpace(input.Description)
	if len(input.Description) == 0 {
		http.Error(w, "description is required", http.StatusBadRequest)
		return
	}
	if len(input.Description) > 1000 {
		http.Error(w, "description is too long", http.StatusBadRequest)
		return
	}
	err = h.service.UpdateDescriptionByID(input.Description, id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetCourseByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
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

func (h *Handler) DeleteCourseByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
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
