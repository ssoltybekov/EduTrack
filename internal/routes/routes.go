package routes

import (
    "edutrack/internal/handlers"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func Routes() *chi.Mux {
    r := chi.NewRouter()

    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.StripSlashes)
   
    r.Route("/users", func(r chi.Router) {
        r.Get("/", handlers.ListUsers)
        r.Post("/", handlers.CreateUser)

        r.Route("/{id}", func(r chi.Router) {
            r.Get("/", handlers.GetUser)
            r.Put("/", handlers.UpdateUser)
            r.Delete("/", handlers.DeleteUser)
        })
    })

   
    r.Route("/lessons", func(r chi.Router) {
        r.Get("/", handlers.ListLessons)
        r.Post("/", handlers.CreateLesson)

        r.Route("/{lesson_id}", func(r chi.Router) {
            r.Get("/", handlers.GetLesson)
            r.Put("/", handlers.UpdateLesson)
            r.Delete("/", handlers.DeleteLesson)

          
            r.Route("/assignments", func(r chi.Router) {
                r.Get("/", handlers.ListAssignments)
                r.Post("/", handlers.CreateAssignment)

                r.Route("/{assignment_id}", func(r chi.Router) {
                    r.Get("/", handlers.GetAssignment)
                    r.Put("/", handlers.UpdateAssignment)
                    r.Delete("/", handlers.DeleteAssignment)

                    r.Route("/submissions", func(r chi.Router) {
                        r.Get("/", handlers.ListSubmissions)      
                        r.Post("/", handlers.CreateSubmission)    

                        r.Route("/{submission_id}", func(r chi.Router) {
                            r.Get("/", handlers.GetSubmission)                
                            r.Post("/grade", handlers.GradeSubmission)        
                        })
                    })
                })
            })
        })
    })

    return r
}