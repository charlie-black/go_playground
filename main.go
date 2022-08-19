package main

import (
	"fmt"

	"github.com/kataras/iris/v12"

	_ "github.com/lib/pq"

	"log"

	"github.com/jmoiron/sqlx"
)

type NoteUpdatePostParams struct {
	ID int `db:"id"`
	Title string `db:"title"`
	Value string `db:"value"`
}

type NoteDeleteParams struct {
	ID int `db:"id"`
}

type NoteAddPostParams struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type Note struct {
	ID    int    `db:"id"`
    Title string `db:"title"`
    Value  string `db:"value"`
}

func main()  {
	db, err := sqlx.Connect("postgres", "user=piccasso dbname=notebook sslmode=disable")
    if err != nil {
        log.Fatalln(err)
	}
	println("Connected to Postgres",db)
	app := iris.New()
    app.Use(iris.Compression)

	fmt.Println("sample")

//get all notes from the database
	app.Get("/notes", func(ctx iris.Context) {
		notes := []Note{}
		err = db.Select(&notes, "SELECT * FROM notes")
		if err != nil {
			fmt.Println(err)
			return
		}
	
		ctx.JSON(notes)
	  })
	
//update a note in the database
	app.Post("/update_notes", func(ctx iris.Context) {
		var params NoteUpdatePostParams
		err := ctx.ReadJSON(&params);
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}

		fmt.Println("params", params)

		_, err = db.NamedExec("UPDATE notes SET title=:title, value=:value WHERE id=:id", params)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}
		

		ctx.JSON(iris.Map{"message": "Note updated"})

	  })

	  app.Delete("/delete_notes", func(ctx iris.Context) {
		var params NoteDeleteParams
		err := ctx.ReadJSON(&params);
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}
		fmt.Println("params", params)
		_, err = db.Exec("DELETE FROM notes WHERE id =$1", params.ID)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}
		ctx.JSON(iris.Map{"message": "Note deleted"})
	})
//add a note to the database
	app.Post("/add_notes", func(ctx iris.Context) {
		var params NoteAddPostParams
		err := ctx.ReadJSON(&params);
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}

		fmt.Println("params", params)

		_, err = db.NamedExec(`INSERT INTO notes (title, value)
        VALUES (:title, :value)`, params)

		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return

		}

		ctx.JSON(iris.Map{"message": "Note Added"})

	  })

	  app.Listen(":8081")

	}


