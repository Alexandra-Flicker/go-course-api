package lesson

import (
	"encoding/json"
	"net/http"
	"newProject_courses/internal/domain/lesson"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service lesson.Service
}

func NewLessonHandler(service lesson.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateLesson(w http.ResponseWriter, r *http.Request) {
	var input lesson.Lesson
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

func (h *Handler) GetLessonByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	lessonResponse, err := h.service.GetLessonByID(id)
	if err != nil {
		http.Error(w, "lesson not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(lessonResponse)
	if err != nil {
		http.Error(w, "failed to marshal lesson: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) GetAllLessons(w http.ResponseWriter, _ *http.Request) {
	lessons, err := h.service.GetAllLessons()
	if err != nil {
		http.Error(w, "failed to fetch lessons: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(lessons)
	if err != nil {
		http.Error(w, "failed to marshal lessons: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) UpdateLessonByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var input lesson.Lesson
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

func (h *Handler) DeleteLessonByID(w http.ResponseWriter, r *http.Request) {
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
