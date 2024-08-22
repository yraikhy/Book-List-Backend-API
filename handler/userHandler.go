package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/yraikhy/readinglisttracker/model"
)

type UserHandler struct {
	DB        *sql.DB
	TokenAuth *jwtauth.JWTAuth
}

func (u *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := u.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)",
		user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user model.User

	// Decode the request body into the user struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Verify the user in the database
	query := "SELECT userid FROM users WHERE username = ? AND password = ?"
	err := u.DB.QueryRow(query, user.Username, user.Password).Scan(&user.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Generate JWT token
	_, tokenString, _ := u.TokenAuth.Encode(map[string]interface{}{"userid": user.UserID})

	// Return the token to the client
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(tokenString))
}
