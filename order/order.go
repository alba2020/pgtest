package order

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Order struct {
	id int
	status string
	login string
}

func FindOneById(dbpool *pgxpool.Pool, id int) Order {
	sql := `select id, status, params->>'login'
	from smmtouch.composite_orders
	where id = $1
	limit 1`

	row := dbpool.QueryRow(context.Background(), sql, id)

	order := Order{}

	err := row.Scan(&order.id, &order.status, &order.login)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return Order{}
	}

	return order
}

func FindOneByLogin(dbpool *pgxpool.Pool, login string) Order {
	sql := `select id, status, params->>'login'
	from smmtouch.composite_orders
	where params->>'login' ~ $1
	limit 1`
	row := dbpool.QueryRow(context.Background(), sql, login)

	order := Order{}

	err := row.Scan(&order.id, &order.status, &order.login)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return Order{}
	}

	return order
}
