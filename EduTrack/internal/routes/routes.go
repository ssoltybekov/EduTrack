package routes

import (
	"edutrack/internal/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/teachers", func(r chi.Router) {
		r.Get("/", handlers.ListTeacher)
		r.Get("/{id}", handlers.GetTeacher)
		r.Post("/", handlers.CreateTeacher)
		r.Put("/{id}", handlers.UpdateTeacher)
		r.Delete("/{id}", handlers.DeleteTeacher)
	})

	r.Route("/students", func(r chi.Router) {
		r.Get("/", handlers.ListStudents)
		r.Get("/{id}", handlers.GetStudent)
		r.Post("/", handlers.CreateStudent)
		r.Put("/{id}", handlers.UpdateStudent)
		r.Delete("/{id}", handlers.DeleteStudent)
	})

	r.Route("/assignments", func(r chi.Router) {
		r.Get("/", handlers.ListAssignments)
		r.Get("/{id}", handlers.GetAssignment)
		r.Post("/", handlers.CreateAssignment)
		r.Put("/{id}", handlers.UpdateAssignment)
		r.Delete("/{id}", handlers.DeleteAssignment)
	})

	r.Route("/submissions", func(r chi.Router) {
		r.Get("/", handlers.ListSubmissions)
		r.Get("/{id}", handlers.GetSubmission)
		r.Post("/", handlers.CreateSubmission)
		r.Put("/{id}", handlers.UpdateSubmission)
		r.Delete("/{id}", handlers.DeleteSubmission)
	})

	return r
}
