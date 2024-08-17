package application

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yraikhy/readinglisttracker/handler"
)

func loadRoutes(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/books", func(r chi.Router) {
		bookHandler := &handler.BookHandler{DB: db}
		r.Post("/", bookHandler.Create)
		r.Get("/", bookHandler.List)
		r.Get("/{id}", bookHandler.GetByID)
		r.Put("/{id}", bookHandler.UpdateByID)
		r.Delete("/{id}", bookHandler.DeleteByID)
	})

	return router
}
