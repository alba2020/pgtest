package db

func Hello() string {
	return "Hello"
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
