package notescontroller

import (
	"fmt"
	models "test/models"

	"github.com/kataras/iris/v12"

	_ "log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitializeEndpoints(app *iris.Application, db *sqlx.DB) {

	app.Get("/notes", func(ctx iris.Context) {
		notes := []models.Note{}
		err := db.Select(&notes, "SELECT * FROM notes")
		if err != nil {
			fmt.Println(err)
			return
		}

		ctx.JSON(notes)
	})

	app.Post("/update_notes", func(ctx iris.Context) {
		var params models.NoteUpdatePostParams

		err := ctx.ReadJSON(&params)

		var count int
		err2 := db.QueryRow(fmt.Sprint("SELECT COUNT(*) from notes where id =", params.ID)).Scan(&count) //Select(&notes, "SELECT count(*) FROM notes WHERE id=$1",params.ID)

		if count == 0 {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": "no notes with given id"})
			return
		}

		if err2 != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}
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
		var params models.NoteDeleteParams
		err := ctx.ReadJSON(&params)
		var count int
		err2 := db.QueryRow(fmt.Sprint("SELECT COUNT(*) from notes where id =", params.ID)).Scan(&count)

		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}
		if count == 0 {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": "no notes with given id"})
			return
		}

		if err2 != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}
		fmt.Println("params", params)
		_, err = db.Exec("DELETE FROM notes WHERE id =$1", params.ID)

		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}

		ctx.JSON(iris.Map{"message": "Note deleted"})
	})

	app.Post("/add_notes", func(ctx iris.Context) {
		var params models.NoteAddPostParams
		err := ctx.ReadJSON(&params)
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

}
