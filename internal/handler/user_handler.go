package handler

import (
	"encoding/json"
	"fmt"
	"gogofi/internal/repo"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	repo *repo.Repo
}

func NewUserHandler(r *repo.Repo) *UserHandler {
	return &UserHandler{repo: r}
}

func (u *UserHandler) GetListUserHandler(w http.ResponseWriter, r *http.Request) {
	var users []repo.UserRepo

	query := "SELECT user_id, user_name, user_email FROM \"User\""
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

func (u *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Printf("vars ==> %v", vars)

	fmt.Fprintln(w, "Hello World")
}

func (u *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello World")
}

func (u *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello World")
}
