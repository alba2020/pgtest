package order

import (
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Order struct {
	Id int 			`json:"id"`
	Link string 	`json:"link"`
	Cur string  	`json:"cur"`
	Status string 	`json:"status"`
}

func (order *Order) fields() []interface{} {
	fields := make([]interface{}, 0)
	
	fields = append(fields, &order.Id)
	fields = append(fields, &order.Link)
	fields = append(fields, &order.Cur)
	fields = append(fields, &order.Status)
	
	return fields
}

func FindOneById(dbpool *pgxpool.Pool, id int) (Order, error) {
	orders, err := NewQuery().SetId(id).Exec(dbpool)

	// fmt.Println(len(orders))

	if err != nil {
		// fmt.Println("err != nil")
		return Order{}, err
	}

	if len(orders) < 1 {
		// fmt.Println("len < 1")
		return Order{}, errors.New("Order not found")
	}

	return orders[0], nil
}
