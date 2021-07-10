package model

type (
	Note struct {
		Id      string `json:"id" firestore:"id"`
		Title   string `json:"title" firestore:"title"`
		Content string `json:"content" firestore:"content"`
	}
)
