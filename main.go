package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello World!")

	//connect to database (postgres)
	db, err := sqlx.Connect("postgres", "user=piccasso dbname=notebook sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	println("connected to database", db)

}
