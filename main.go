package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"pgtest/user"

	"github.com/jackc/pgx/v4/pgxpool"
)

// https://www.calhoun.io/querying-for-a-single-record-using-gos-database-sql-package/

func getPool() *pgxpool.Pool {
	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return dbpool
}


func usersHandler(w http.ResponseWriter, r *http.Request) {
	dbpool := getPool()
	defer dbpool.Close()

	query := r.URL.Query()
    email := query.Get("email")

	log.Printf("GET /users?email=%s\n", email)

	users, err := user.FindManyByEmail(dbpool, email, 5)
	if (err != nil) {
		fmt.Fprintf(os.Stderr, "find users failed: %v\n", err)
	}
	// for _, user := range users {
	// 	fmt.Println(user)
	// }

	w.Header().Set("Content-type", "application/json")
	j, _ := json.Marshal(users)
	w.Write(j)
}

func main() {
	dbpool := getPool()
	defer dbpool.Close()

	// for i := 1001; i < 1005; i++ {
	// 	user, err := user.FindOneById(dbpool, i)
	// 	if (err != nil) {
	// 		//fmt.Fprintf(os.Stderr, "find user failed: %v\n", err)
	// 	} else {
	// 		fmt.Println(user)
	// 	}
	// }

	// count := db.Count(dbpool)
	// fmt.Println("count=", count)

	// order := db.FindOrderByLogin(dbpool, "abc")
	// fmt.Println(order)

	http.HandleFunc("/users", usersHandler)
	log.Println("Starting server")
	http.ListenAndServe(":8080", nil)
}
