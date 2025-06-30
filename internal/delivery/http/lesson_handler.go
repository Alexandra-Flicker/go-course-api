package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"newProject_courses/internal/entity"
	"newProject_courses/internal/service"
)

type LessonHandler struct {
	service service.LessonService
}

func NewLessonHandler(service service.LessonService) *LessonHandler {
	return &LessonHandler{service: service}
}

func (h *LessonHandler) CreateLesson(w http.ResponseWriter, r *http.Request) {
	var input entity.Lesson
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateLesson(&input); err != nil {
		http.Error(w, "failed to create lesson: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *LessonHandler) GetLessonByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	lesson, err := h.service.GetLessonByID(id)
	if err != nil {
		http.Error(w, "lesson not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lesson)
}

func (h *LessonHandler) GetAllLessons(w http.ResponseWriter, r *http.Request) {
	lessons, err := h.service.GetAllLessons()
	if err != nil {
		http.Error(w, "failed to fetch lessons: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lessons)
}

func (h *LessonHandler) UpdateLessonByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var input entity.Lesson
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateLessonByID(id, &input); err != nil {
		http.Error(w, "failed to update lesson: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *LessonHandler) DeleteLessonByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteLessonByID(id); err != nil {
		http.Error(w, "failed to delete lesson: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
