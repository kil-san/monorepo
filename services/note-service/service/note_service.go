package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/kil-san/micro-serv/note-service/repo"
	"github.com/kil-san/micro-serv/pkg/model"
)

type NoteService struct {
	repo repo.Repo
}

func NewNoteService(repo repo.Repo) *NoteService {
	return &NoteService{
		repo: repo,
	}
}

func (s *NoteService) CreateNote(ctx context.Context, note model.Note) (model.Note, error) {
	var newNote model.Note
	note.Id = uuid.New().String()
	newNote, err := s.repo.Create(ctx, note)
	if err != nil {
		return newNote, err
	}
	return newNote, nil
}

func (s *NoteService) GetNote(ctx context.Context, id string) (model.Note, error) {
	var note model.Note
	note, err := s.repo.Get(ctx, id)
	if err != nil {
		return note, err
	}
	return note, nil
}

func (s *NoteService) UpdateNote(ctx context.Context, id string, note model.Note) error {
	err := s.repo.Update(ctx, id, note)
	if err != nil {
		return err
	}
	return nil
}

func (s *NoteService) DeleteNote(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *NoteService) GetNotes(ctx context.Context, id string) ([]model.Note, error) {
	var notes []model.Note

	return notes, nil
}
