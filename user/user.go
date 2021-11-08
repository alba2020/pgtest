package user

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func FindOneById(dbpool *pgxpool.Pool, id int) (User, error) {
	sql := `select id, name, email
	from smmtouch.users
	where id = $1
	limit 1`
	row := dbpool.QueryRow(context.Background(), sql, id)

	user := User{}
	err := row.Scan(&user.Id, &user.Name, &user.Email)

	return user, err
}

func FindManyByEmail(dbpool *pgxpool.Pool, email string, limit int) ([]User, error) {
	if (limit < 1) {
		limit = 1_000_000_000

	}
	sql := `SELECT
		id,
		COALESCE(name, '#no_name') as name,
		email
	FROM smmtouch.users
	WHERE email ~ $1
	limit $2`
	
	rows, err := dbpool.Query(context.Background(), sql, email, limit)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	// err = rows.Err()
	// if err != nil {
	//   panic(err)
	// }
	
	return users, nil
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
