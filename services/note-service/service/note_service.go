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

func (s *NoteService) CreateNote(ctx context.Context, ownerUid string, note model.Note) (model.Note, error) {
	var newNote model.Note
	note.Id = uuid.New().String()
	note.OwnerUid = ownerUid
	err := s.repo.Create(ctx, ownerUid, note)
	if err != nil {
		return newNote, err
	}
	return note, nil
}

func (s *NoteService) GetNote(ctx context.Context, ownerUid string, noteID string) (model.Note, error) {
	var note model.Note
	note, err := s.repo.Get(ctx, ownerUid, noteID)
	if err != nil {
		return note, err
	}
	return note, nil
}

func (s *NoteService) UpdateNote(ctx context.Context, ownerUid string, note model.Note) error {
	err := s.repo.Update(ctx, ownerUid, note)
	if err != nil {
		return err
	}
	return nil
}

func (s *NoteService) DeleteNote(ctx context.Context, ownerUid string, noteID string) error {
	err := s.repo.Delete(ctx, ownerUid, noteID)
	if err != nil {
		return err
	}
	return nil
}

func (s *NoteService) GetNotes(ctx context.Context, ownerUid string, page uint32) ([]model.Note, error) {
	var notes []model.Note
	notes, err := s.repo.List(ctx, ownerUid, page)
	if err != nil {
		return notes, err
	}

	return notes, nil
}
