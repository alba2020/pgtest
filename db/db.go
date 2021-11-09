package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetPool() *pgxpool.Pool {
	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return dbpool
}

// func xString(dbpool *pgxpool.Pool, sql string) string {
// 	var res string
// 	err := dbpool.QueryRow(context.Background(), sql).Scan(&res)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
// 		os.Exit(1)
// 	}
// 	return res
// }

// func xInt(dbpool *pgxpool.Pool, sql string) int {
// 	var res int
// 	err := dbpool.QueryRow(context.Background(), sql).Scan(&res)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
// 		os.Exit(1)
// 	}
// 	return res
// }
