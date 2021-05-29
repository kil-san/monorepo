package service

import (
	"context"

	"github.com/kil-san/micro-serv/gateway-service/graph/model"
	"github.com/kil-san/micro-serv/note-service/pb"
)

type noteService struct {
	client pb.NoteRPCClient
}

type NoteService interface {
	CreateNote(ctx context.Context, note model.NewNote) (model.Note, error)
	GetNote(ctx context.Context, id string) (model.Note, error)
	UpdateNote(ctx context.Context, note model.NoteUpdate) error
	DeleteNote(ctx context.Context, id string) error
}

var NoteRPCEndpoint = "127.0.0.1:8008"

func NewNoteService(client pb.NoteRPCClient) NoteService {
	s := &noteService{
		client: client,
	}
	return s
}

func (s *noteService) CreateNote(ctx context.Context, note model.NewNote) (model.Note, error) {
	var newNote model.Note
	notePb, err := s.client.CreateNote(ctx, &pb.Note{
		Title: note.Title,
	})
	if err != nil {
		return newNote, err
	}

	newNote.ID = notePb.Id
	newNote.Title = notePb.Title
	newNote.Status = notePb.Status

	return newNote, nil
}

func (s *noteService) GetNote(ctx context.Context, id string) (model.Note, error) {
	var note model.Note
	notePb, err := s.client.GetNote(ctx, &pb.SingleNote{
		NoteId: id,
	})
	if err != nil {
		return note, err
	}

	note.ID = notePb.Id
	note.Title = notePb.Title
	note.Status = notePb.Status

	return note, nil
}

func (s *noteService) UpdateNote(ctx context.Context, note model.NoteUpdate) error {
	_, err := s.client.UpdateNote(ctx, &pb.Note{
		Id:     note.ID,
		Title:  note.Title,
		Status: note.Status,
	})

	return err
}

func (s *noteService) DeleteNote(ctx context.Context, id string) error {
	_, err := s.client.DeleteNote(ctx, &pb.SingleNote{
		NoteId: id,
	})

	return err
}
