// package main

// import (
// 	"database/sql"
// 	"log"
// 	_ "github.com/lib/pq"
// )

// func main() {
// 	connStr := "postgres://postgres::@localhost:5432/sqlctest?sslmode=disabled"

// 	db, err := sql.Open("postgres", connStr)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer db.Close()
// }