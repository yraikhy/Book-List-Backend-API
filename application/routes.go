package application

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/yraikhy/readinglisttracker/handler"
)

func loadRoutes(db *sql.DB, db2 *sql.DB, tokenAuth *jwtauth.JWTAuth) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	// Public route
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// User Auth route
	router.Route("/user", func(r chi.Router) {
		userHandler := &handler.UserHandler{DB: db2, TokenAuth: tokenAuth}

		r.Post("/register", userHandler.Register)
		r.Post("/login", userHandler.Login)
	})

	// Protected Route -- unaccessible without login
	router.Route("/books", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		bookHandler := &handler.BookHandler{DB: db}
		r.Post("/", bookHandler.Create)
		r.Get("/", bookHandler.List)
		r.Get("/{id}", bookHandler.GetByID)
		r.Put("/{id}", bookHandler.UpdateByID)
		r.Delete("/{id}", bookHandler.DeleteByID)
	})

	return router
}
