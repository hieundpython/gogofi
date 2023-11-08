package handler

import (
	"encoding/json"
	"fmt"
	"gogofi/internal/repo"
	"io"
	"log"
	"net/http"
	"time"
)

type UserHandler struct {
	repo *repo.Repo
}

func NewUserHandler(r *repo.Repo) *UserHandler {
	return &UserHandler{repo: r}
}

func (u *UserHandler) GetListUserHandler(w http.ResponseWriter, r *http.Request) {
	var users []repo.UserRepo

	query := "SELECT user_id, user_name, user_email, create_at FROM \"User\""
	if err := u.repo.DB.Select(&users, query); err != nil {
		log.Fatal(err)
	}

	js, err := json.Marshal(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (u *UserHandler) GetUserId(w http.ResponseWriter, r *http.Request) {
	var users []repo.UserRepo

	query := "SELECT user_id, user_name, user_email, create_at FROM \"User\""
	if err := u.repo.DB.Select(&users, query); err != nil {
		log.Fatal(err)
	}
}

func (u *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var userRequest CreateUserRequest

	if err := json.Unmarshal(data, &userRequest); err != nil {
		http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}

	tx := u.repo.DB.MustBegin()

	query := `INSERT INTO "User" (user_name, user_email, first_name, last_name, create_at, active)
		VALUES($1, $2, $3, $4, $5, $6)`

	stmt, err := tx.Prepare(query)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(userRequest.UserName, userRequest.UserEmail, userRequest.FirstName, userRequest.LastName, time.Now(), true)
	tx.Commit()

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func (u *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello World")
}

func (u *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello World")
}

type CreateUserRequest struct {
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
