package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type UserRepo struct {
	UserId    int    `json:"user_id" db:"user_id"`
	UserName  string `json:"user_name" db:"user_name"`
	UserEmail string `json:"user_email" db:"user_email"`
}

type UserHandler struct {
	db *sqlx.DB
}

func NewUserHandler(repo *sqlx.DB) *UserHandler {
	return &UserHandler{repo}
}

func (u *UserHandler) GetListUserHandler(w http.ResponseWriter, r *http.Request) {
	var users []UserRepo

	query := "SELECT user_id, user_name, user_email FROM \"User\""
	if err := u.db.Select(&users, query); err != nil {
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

const (
	host     = "localhost"
	port     = 5432
	user     = "hieund"
	password = "It123456@"
	dbname   = "gogofi"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Setup and connect database
	db, err := sqlx.Open("postgres", psqlconn)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect database success")
	defer db.Close()

	fmt.Println("Starting read all user from page")

	var userHandler = NewUserHandler(db)

	r := mux.NewRouter()

	// Routing for user (CRUD)
	r.HandleFunc("/api/users", userHandler.GetListUserHandler).Methods("GET")
	r.HandleFunc("/api/users", userHandler.CreateUserHandler).Methods("POST")
	r.HandleFunc("/api/users/{id}", userHandler.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/api/users/{id}", userHandler.DeleteUserHandler).Methods("DELETE")

	fmt.Println("Running service!!")

	http.ListenAndServe("localhost:3000", r)
}
