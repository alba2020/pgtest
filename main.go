package main

import (
	"context"
	"fmt"
	"os"
	"pgtest/db"

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

func xString(dbpool *pgxpool.Pool, sql string) string {
	var res string
	err := dbpool.QueryRow(context.Background(), sql).Scan(&res)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	return res
}

func xInt(dbpool *pgxpool.Pool, sql string) int {
	var res int
	err := dbpool.QueryRow(context.Background(), sql).Scan(&res)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	return res
}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	dbpool := getPool()
	defer dbpool.Close()

	for i := 1001; i < 1020; i++ {
		user := db.FindById(dbpool, i)
		fmt.Println(user)
	}
}
