package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"gogofi/internal/database/services"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserController struct {
	services *services.Queries
}

func NewUserController(s *services.Queries) *UserController {
	return &UserController{services: s}
}

func (u *UserController) GetListUserHandler(w http.ResponseWriter, r *http.Request) {
	users, err := u.services.GetListUsers(context.Background())

	if err != nil {
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

func (u *UserController) GetUserId(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := u.services.GetUser(context.Background(), int32(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	js, err := json.Marshal(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (u *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("value on form %v", body)
}

func (u *UserController) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func (u *UserController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello World")
}

type CreateUserRequest struct {
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
