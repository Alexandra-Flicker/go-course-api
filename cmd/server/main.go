package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"newProject_courses/config"
	course2 "newProject_courses/internal/delivery/http/course"
	handler "newProject_courses/internal/delivery/http/lesson"
	module2 "newProject_courses/internal/delivery/http/module"
	"newProject_courses/internal/domain/course"
	"newProject_courses/internal/domain/lesson"
	"newProject_courses/internal/domain/module"
	"newProject_courses/internal/repository"
	courserepo "newProject_courses/internal/repository/pgx/course"
	lessonrepo "newProject_courses/internal/repository/pgx/lesson"
	modulerepo "newProject_courses/internal/repository/pgx/module"
)

func main() {
	cfg := config.LoadConfig()

	db := repository.InitDB(cfg.DB.DSN)
	r := chi.NewRouter()

	courseRepo := courserepo.NewCourseRepo(db)
	courseService := course.NewCourseService(courseRepo)
	courseHandler := course2.NewCourseHandler(courseService)

	r.Route("/courses", func(r chi.Router) {
		r.Get("/", courseHandler.GetAll)
		r.Get("/{id}", courseHandler.GetCourseByID)
		r.Post("/", courseHandler.CreateCourse)
		r.Put("/{id}", courseHandler.UpdateDescriptionByID)
		r.Delete("/{id}", courseHandler.DeleteCourseByID)
	})

	moduleRepo := modulerepo.NewModuleRepo(db)
	moduleService := module.NewModuleService(moduleRepo)
	moduleHandler := module2.NewModuleHandler(moduleService)

	r.Route("/modules", func(r chi.Router) {
		r.Post("/", moduleHandler.CreateModule)
		r.Get("/", moduleHandler.GetAllModules)
		r.Get("/{id}", moduleHandler.GetModuleByID)
		r.Put("/{id}", moduleHandler.UpdateModuleTitleByID)
		r.Delete("/{id}", moduleHandler.DeleteModuleByID)

	})

	lessonRepo := lessonrepo.NewLessonRepo(db)
	lessonService := lesson.NewLessonService(lessonRepo)
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
