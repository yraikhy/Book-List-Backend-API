package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() *App { // Returns App type
	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to initialize databse: %v", err)
	}

	app := &App{ // Instance of App struct
		router: loadRoutes(db),
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
