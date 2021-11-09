package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"pgtest/db"
	"pgtest/order"
	"pgtest/user"
	"strconv"

	"github.com/gorilla/mux"
)

func ManyUsers(w http.ResponseWriter, r *http.Request) {
	dbpool := db.GetPool()
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

type OneOrderResponse struct {
	Order interface{} `json:"order"`
	Error interface{} `json:"error"`
}

func OneOrder(w http.ResponseWriter, r *http.Request) {
	dbpool := db.GetPool()
	defer dbpool.Close()

	vars := mux.Vars(r)
    id_str, _ := vars["id"]
	id, _ := strconv.Atoi(id_str)

	log.Printf("GET /orders/%d\n", id)


	order, err := order.FindOneById(dbpool, id)

	res := OneOrderResponse{}

	if (err != nil) {
		res.Order = nil
		res.Error = fmt.Sprintf("%v", err)
	} else {
		res.Order = order
		res.Error = nil
	}

	j, _ := json.Marshal(res)
	w.Header().Set("Content-type", "application/json")
	w.Write(j)
}
