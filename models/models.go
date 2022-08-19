package models

type Note struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
	Value  string `db:"value"`
}

type NoteAddParams struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type NoteUpdateParams struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
	Value string `db:"value"`
}

type NoteDeleteParams struct {
	ID int `db:"id"`
}