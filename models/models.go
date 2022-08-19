package models

type Note struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
	Value  string `db:"value"`
}