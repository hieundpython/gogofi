package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"gogofi/internal/connect"

	_ "github.com/lib/pq"
)

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
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect database success")

	var userHandler = connect.InitializeApp(db)

	r := mux.NewRouter()

	// Routing for user (CRUD)
	r.HandleFunc("/api/users", userHandler.GetListUserHandler).Methods("GET")
	r.HandleFunc("/api/users", userHandler.CreateUserHandler).Methods("POST")
	r.HandleFunc("/api/users/{id}", userHandler.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/api/users/{id}", userHandler.DeleteUserHandler).Methods("DELETE")

	fmt.Println("Running service!!")

	http.ListenAndServe("localhost:3000", r)
}
