package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"newProject_courses/config"
	handler "newProject_courses/internal/delivery/http"
	"newProject_courses/internal/repository"
	"newProject_courses/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	db := repository.InitDB(cfg.DB.DSN)
	r := chi.NewRouter()

	courseRepo := repository.NewCourseRepo(db)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	r.Route("/courses", func(r chi.Router) {
		r.Get("/", courseHandler.GetAll)
		r.Get("/{id}", courseHandler.GetCourseByID)
		r.Post("/", courseHandler.CreateCourse)
		r.Put("/{id}", courseHandler.UpdateDescriptionByID)
		r.Delete("/{id}", courseHandler.DeleteCourseByID)
	})
	log.Println("Server running on:8080")
	err := http.ListenAndServe(cfg.Server.Port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
