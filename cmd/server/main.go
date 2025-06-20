package main

import (
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

	courseRepo := repository.NewCourseRepo(db)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	http.HandleFunc("/courses", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			courseHandler.GetAll(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/courses/get/id", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			courseHandler.GetCourseByID(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/courses/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			courseHandler.CreateCourse(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/courses/update/description", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			courseHandler.UpdateDescriptionByID(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/courses/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			courseHandler.DeleteCourseByID(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server running on:8080")

	err := http.ListenAndServe(cfg.Server.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
