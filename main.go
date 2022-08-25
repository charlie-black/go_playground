package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	notescontroller "test/controllers/notescontroller"
	_ "github.com/lib/pq"

	"log"

	"github.com/jmoiron/sqlx"
)



func main() {
	
	db, err := sqlx.Connect("postgres", "user=piccasso dbname=notebook sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	println("Connected to Postgres", db)
	app := iris.New()
	app.Use(iris.Compression)

	fmt.Println("sample")

	notescontroller.InitializeEndpoints(app,db)
	app.Listen(":8081")

}
