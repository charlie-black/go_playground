package models

type NoteUpdatePostParams struct {
	ID    int    `db:"id"`
	Title string `json:"title"`
	Value string `json:"value"`
}

type NoteDeleteParams struct {
	ID int `db:"id"`
}

type NoteAddPostParams struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type Note struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Value string `json:"value"`
}