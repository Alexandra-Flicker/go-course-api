package module

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"newProject_courses/internal/domain/module"
	"strconv"
)

type Handler struct {
	service module.Service
}

func NewModuleHandler(s module.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) CreateModule(w http.ResponseWriter, r *http.Request) {
	var input struct {
		CourseID int    `json:"course_id"`
		Title    string `json:"title"`
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to decode request", http.StatusBadRequest)
		return
	}
	module := module.Module{
		CourseID: input.CourseID,
		Title:    input.Title,
	}
	if err := h.service.CreateModule(module); err != nil {
		log.Println(err)
		http.Error(w, "failed to create module", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(module)
	if err != nil {
		log.Printf("failed to encode response: %v", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *Handler) GetAllModules(w http.ResponseWriter, _ *http.Request) {
	modules, err := h.service.GetAllModules()
	if err != nil {
		http.Error(w, "failed to get modules", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(modules)
	if err != nil {
		log.Printf("failed to encode response: %v", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) GetModuleByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "failed to parse id", http.StatusBadRequest)
		return
	}
	module, err := h.service.GetModuleByID(id)
	if err != nil {
		http.Error(w, "failed to get module", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(module)
	if err != nil {
		log.Printf("failed to parse id: %v", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateModuleTitleByID(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
		ID    int    `json:"id"`
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Printf("failed to decode request: %v", err)
		http.Error(w, "failed to decode request", http.StatusBadRequest)
		return
	}
	err = h.service.UpdateModuleTitleByID(input.ID, module.Module{Title: input.Title})
	if err != nil {
		log.Printf("failed to update module: %v", err)
		http.Error(w, "failed to update module", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteModuleByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "failed to parse id", http.StatusBadRequest)
		return
	}
	err = h.service.DeleteModuleByID(id)
	if err != nil {
		log.Printf("failed to delete module: %v", err)
		http.Error(w, "failed to delete module", http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
}
