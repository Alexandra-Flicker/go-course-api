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

	moduleRepo := repository.NewModuleRepo(db)
	moduleService := service.NewModuleService(moduleRepo)
	moduleHandler := handler.NewModuleHandler(moduleService)

	r.Route("/modules", func(r chi.Router) {
		r.Post("/", moduleHandler.CreateModule)
		r.Get("/", moduleHandler.GetAllModules)
		r.Get("/{id}", moduleHandler.GetModuleByID)
		r.Put("/{id}", moduleHandler.UpdateModuleTitleByID)
		r.Delete("/{id}", moduleHandler.DeleteModuleByID)

	})

	lessonRepo := repository.NewLessonRepo(db)
	lessonService := service.NewLessonService(lessonRepo)
	lessonHandler := handler.NewLessonHandler(lessonService)

	r.Route("/lessons", func(r chi.Router) {
		r.Post("/", lessonHandler.CreateLesson)
		r.Get("/", lessonHandler.GetAllLessons)
		r.Get("/{id}", lessonHandler.GetLessonByID)
		r.Put("/{id}", lessonHandler.UpdateLessonByID)
		r.Delete("/{id}", lessonHandler.DeleteLessonByID)
	})

	log.Println("Server running on:8080")
	err := http.ListenAndServe(cfg.Server.Port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
