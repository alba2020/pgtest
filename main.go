package main

import (
	"log"
	"net/http"
	"pgtest/api"
	"pgtest/db"

	"github.com/gorilla/mux"
)

// https://www.calhoun.io/querying-for-a-single-record-using-gos-database-sql-package/


func main() {
	dbpool := db.GetPool()
	defer dbpool.Close()

	r := mux.NewRouter()
    r.HandleFunc("/orders/{id}", api.OneOrder)
	r.HandleFunc("/users", api.ManyUsers)
	
	log.Println("Starting server")
	http.ListenAndServe(":8080", r)
}
