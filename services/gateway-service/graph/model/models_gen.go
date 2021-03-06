// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CheckListItem struct {
	Index int    `json:"index"`
	Title string `json:"title"`
	State bool   `json:"state"`
}

type CheckListItemInput struct {
	Index int    `json:"index"`
	Title string `json:"title"`
	State bool   `json:"state"`
}

type NewNote struct {
	Title     string                `json:"title"`
	Content   string                `json:"content"`
	Checklist []*CheckListItemInput `json:"checklist"`
}

type Note struct {
	ID        string           `json:"id"`
	Title     string           `json:"title"`
	Content   string           `json:"content"`
	Checklist []*CheckListItem `json:"checklist"`
}

type NoteUpdate struct {
	ID        string                `json:"id"`
	Title     string                `json:"title"`
	Content   string                `json:"content"`
	Checklist []*CheckListItemInput `json:"checklist"`
}
