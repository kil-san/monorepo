package repo

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/kil-san/micro-serv/pkg/model"
	_ "github.com/mattn/go-sqlite3"
)

type SqlRepo struct {
	db *sql.DB
}

func NewSqlRepo(db *sql.DB) SqlRepo {
	return SqlRepo{
		db: db,
	}
}

func (r SqlRepo) Create(ctx context.Context, data model.Note) (model.Note, error) {
	var note model.Note
	tx, _ := r.db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO note (title,status) VALUES (?,?)")
	result, err := stmt.Exec(data.Title, data.Status)
	if err != nil {
		tx.Rollback()
		return note, err
	}
	tx.Commit()

	note = data
	index, err := result.LastInsertId()
	if err != nil {
		return note, err
	}
	note.Id = strconv.FormatInt(index, 10)

	return note, nil
}

func (r SqlRepo) Get(ctx context.Context, id string) (model.Note, error) {
	var note model.Note
	row := r.db.QueryRow("SELECT id, title, status FROM note WHERE id=?", id)

	err := row.Scan(&note.Id, &note.Title, &note.Status)
	if err != nil {
		return note, err
	}

	return note, nil
}

func (r SqlRepo) Delete(ctx context.Context, id string) error {
	tx, _ := r.db.Begin()
	stmt, _ := tx.Prepare("DELETE FROM note WHERE id=?")
	_, err := stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (r SqlRepo) Update(ctx context.Context, id string, data model.Note) error {
	tx, _ := r.db.Begin()
	stmt, _ := tx.Prepare("UPDATE note SET title=?, status=? WHERE id=?")
	_, err := stmt.Exec(data.Title, data.Status, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
