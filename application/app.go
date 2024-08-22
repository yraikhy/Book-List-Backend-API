package application

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

type App struct {
	router    http.Handler
	tokenAuth *jwtauth.JWTAuth
}

func New() *App { // Returns App type
	db, db2, err := initDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	tokenAuth := jwtauth.New("HS256", []byte("books"), nil)

	app := &App{ // Instance of App struct
		router:    loadRoutes(db, db2, tokenAuth),
		tokenAuth: tokenAuth,
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
