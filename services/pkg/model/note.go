package model

type (
	Note struct {
		OwnerUid string `json:"ownerUid" firestore:"ownerUid"`
		Id       string `json:"id" firestore:"id"`
		Title    string `json:"title" firestore:"title"`
		Content  string `json:"content" firestore:"content"`
	}
)
