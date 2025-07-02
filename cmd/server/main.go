package main

import (
	"log"
	"net/http"
	"newProject_courses/config"
	deliveryHttp "newProject_courses/internal/delivery/http"
	coursehandler "newProject_courses/internal/delivery/http/course"
	lessonhandler "newProject_courses/internal/delivery/http/lesson"
	modulehandler "newProject_courses/internal/delivery/http/module"
	"newProject_courses/internal/domain/course"
	"newProject_courses/internal/domain/lesson"
	"newProject_courses/internal/domain/module"
	"newProject_courses/internal/repository"
	courserepo "newProject_courses/internal/repository/pgx/course"
	lessonrepo "newProject_courses/internal/repository/pgx/lesson"
	modulerepo "newProject_courses/internal/repository/pgx/module"
	"newProject_courses/migrations"
)

func main() {
	cfg := config.LoadConfig()

	db := repository.InitDB(cfg.DB.DSN)
	migrations.MustRunMigrations(cfg.DB.DSN)

	// Init repo
	courseRepo := courserepo.NewCourseRepo(db)
	moduleRepo := modulerepo.NewModuleRepo(db)
	lessonRepo := lessonrepo.NewLessonRepo(db)

	// Init service
	courseService := course.NewCourseService(courseRepo)
	moduleService := module.NewModuleService(moduleRepo)
	lessonService := lesson.NewLessonService(lessonRepo)

	// Init handler
	courseHandler := coursehandler.NewCourseHandler(courseService)
	moduleHandler := modulehandler.NewModuleHandler(moduleService)
	lessonHandler := lessonhandler.NewLessonHandler(lessonService)

	r := deliveryHttp.InitRoutes(courseHandler, moduleHandler, lessonHandler)

	log.Println("Server running on:8080")
	err := http.ListenAndServe(cfg.Server.Port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
