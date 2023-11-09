package main

import (
	"context"
	"fmt"
	"gogofi/internal/database/services"
	"log"

	"github.com/jackc/pgx/v5"

	_ "github.com/lib/pq"
)

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

	conn, err := pgx.Connect(context.Background(), psqlconn)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())
	fmt.Println("connect database success")

	queries := services.New(conn)

	users, err := queries.GetListUsers(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Values of users ==> %v", users)

	// var userHandler = connect.InitializeApp(db)

	// r := mux.NewRouter()

	// // Routing for user (CRUD)
	// r.HandleFunc("/api/users", userHandler.GetListUserHandler).Methods("GET")
	// r.HandleFunc("/api/users", userHandler.CreateUserHandler).Methods("POST")
	// r.HandleFunc("/api/users/{id}", userHandler.UpdateUserHandler).Methods("PUT")
	// r.HandleFunc("/api/users/{id}", userHandler.DeleteUserHandler).Methods("DELETE")

	// fmt.Println("Running service!!")

	// http.ListenAndServe("localhost:3000", r)
}
