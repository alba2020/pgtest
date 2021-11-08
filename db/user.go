package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	id int
	name string
	email string
}

func FindUserById(dbpool *pgxpool.Pool, id int) User {
	sql := "select id, name, email from smmtouch.users where id = $1 limit 1"
	row := dbpool.QueryRow(context.Background(), sql, id)

	user := User{}

	err := row.Scan(&user.id, &user.name, &user.email)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return User{}
	}

	return user
}
 
func Count(dbpool *pgxpool.Pool) int {
	sql := "select count(id) from smmtouch.users"
	row := dbpool.QueryRow(context.Background(), sql)

	var res int
	err := row.Scan(&res)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return 0
	}

	return res
}
