package http

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/jmoiron/sqlx"
	coursehandler "newProject_courses/internal/delivery/http/course"
	lessonhandler "newProject_courses/internal/delivery/http/lesson"
	modulehandler "newProject_courses/internal/delivery/http/module"
)

func InitRoutes(courseHandler *coursehandler.Handler, moduleHandler *modulehandler.Handler, lessonHandler *lessonhandler.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/courses", func(r chi.Router) {
		r.Get("/", courseHandler.GetAll)
		r.Get("/{id}", courseHandler.GetCourseByID)
		r.Post("/", courseHandler.CreateCourse)
		r.Put("/{id}", courseHandler.UpdateDescriptionByID)
		r.Delete("/{id}", courseHandler.DeleteCourseByID)
	})

	r.Route("/modules", func(r chi.Router) {
		r.Post("/", moduleHandler.CreateModule)
		r.Get("/", moduleHandler.GetAllModules)
		r.Get("/{id}", moduleHandler.GetModuleByID)
		r.Put("/{id}", moduleHandler.UpdateModuleTitleByID)
		r.Delete("/{id}", moduleHandler.DeleteModuleByID)

	})

	r.Route("/lessons", func(r chi.Router) {
		r.Post("/", lessonHandler.CreateLesson)
		r.Get("/", lessonHandler.GetAllLessons)
		r.Get("/{id}", lessonHandler.GetLessonByID)
		r.Put("/{id}", lessonHandler.UpdateLessonByID)
		r.Delete("/{id}", lessonHandler.DeleteLessonByID)
	})
	return r
}
