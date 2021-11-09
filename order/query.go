package order

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Query struct {
	limit int
	sql string
	n int
	params []interface{}
}

func NewQuery() *Query {
	q := Query{}
	q.sql = `SELECT
		id,
		params->>'link',
		params->>'cur',
		status
	FROM smmtouch.composite_orders
	WHERE true`

	return &q
}

func (q *Query) SetId(id int) *Query {
	q.n++
	q.sql += fmt.Sprintf(" AND id = $%d", q.n)
	q.params = append(q.params, id)
	return q
}

func (q *Query) SetLink(link string) *Query {
	q.n++
	q.sql += fmt.Sprintf(" AND params->>'link' ~* $%d", q.n)
	q.params = append(q.params, link)
	return q
}

func (q *Query) SetCur(cur string) *Query {
	q.n++
	q.sql += fmt.Sprintf(" AND params->>'cur' = $%d", q.n)
	q.params = append(q.params, cur)
	return q
}

func (q *Query) SetLimit(limit int) *Query {
	q.limit = limit
	return q
}

func (q *Query) ToSql() string {
	if (q.limit > 0) {
		q.n++
		q.sql += fmt.Sprintf(" LIMIT $%d", q.n)
		q.params = append(q.params, q.limit)
	}
	return q.sql
}

func (q *Query) Exec(dbpool *pgxpool.Pool) ([]Order, error) {
	rows, err := dbpool.Query(
		context.Background(), q.ToSql(), q.params...)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	orders := make([]Order, 0)

	for rows.Next() {
		order := Order{}
		err := rows.Scan(order.fields()...)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	
	return orders, nil
}
