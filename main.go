package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func GetListUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello World")
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello World")
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello World")
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello World")
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "It123456@"
	dbname   = "gogofi"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Setup and connect database
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect database success")
	defer db.Close()

	r := mux.NewRouter()

	// Routing for user (CRUD)
	r.HandleFunc("/api/users", GetListUserHandler).Methods("GET")
	r.HandleFunc("/api/users", CreateUserHandler).Methods("POST")
	r.HandleFunc("/api/users/{id}", UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/api/users/{id}", DeleteUserHandler).Methods("DELETE")

	fmt.Println("Running service!!")

	http.ListenAndServe("localhost:3000", r)
}
