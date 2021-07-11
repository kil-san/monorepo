package service

import (
	"context"
	"os"

	"github.com/kil-san/micro-serv/gateway-service/graph/model"
	"github.com/kil-san/micro-serv/note-service/pb"
)

type noteService struct {
	client pb.NoteRPCClient
}

type NoteService interface {
	CreateNote(ctx context.Context, ownerUid string, note model.NewNote) (model.Note, error)
	GetNote(ctx context.Context, ownerUid string, noteID string) (model.Note, error)
	UpdateNote(ctx context.Context, ownerUid string, note model.NoteUpdate) error
	DeleteNote(ctx context.Context, ownerUid string, noteID string) error
	GetNotes(ctx context.Context, ownerUid string, page uint32) ([]*model.Note, error)
}

var NoteRPCEndpoint = os.Getenv("NOTE_SERVICE_ENDPOINT")

func NewNoteService(client pb.NoteRPCClient) NoteService {
	s := &noteService{
		client: client,
	}
	return s
}

func (s *noteService) CreateNote(ctx context.Context, ownerUid string, note model.NewNote) (model.Note, error) {
	var newNote model.Note
	notePb, err := s.client.CreateNote(ctx, &pb.CreateNoteRequest{
		OwnerUid: ownerUid,
		Note: &pb.Note{
			Title:   note.Title,
			Content: note.Content,
		},
	})
	if err != nil {
		return newNote, err
	}

	newNote.ID = notePb.Id
	newNote.Title = notePb.Title
	newNote.Content = notePb.Content

	return newNote, nil
}

func (s *noteService) GetNote(ctx context.Context, ownerUid string, noteID string) (model.Note, error) {
	var note model.Note
	notePb, err := s.client.GetNote(ctx, &pb.SimpleRequest{
		OwnerUid: ownerUid,
		NoteId:   noteID,
	})
	if err != nil {
		return note, err
	}

	note.ID = notePb.Id
	note.Title = notePb.Title
	note.Content = notePb.Content

	return note, nil
}

func (s *noteService) UpdateNote(ctx context.Context, ownerUid string, note model.NoteUpdate) error {
	_, err := s.client.UpdateNote(ctx, &pb.CreateNoteRequest{
		OwnerUid: ownerUid,
		Note: &pb.Note{
			Title:   note.Title,
			Content: note.Content,
		},
	})

	return err
}

func (s *noteService) DeleteNote(ctx context.Context, ownerUid string, noteID string) error {
	_, err := s.client.DeleteNote(ctx, &pb.SimpleRequest{
		OwnerUid: ownerUid,
		NoteId:   noteID,
	})

	return err
}

func (s *noteService) GetNotes(ctx context.Context, ownerUid string, page uint32) ([]*model.Note, error) {
	var notes []*model.Note
	notePb, err := s.client.GetNotes(ctx, &pb.GetNotesRequest{
		OwnerUid: ownerUid,
		Page:     int32(page),
	})
	if err != nil {
		return notes, err
	}

	notes = make([]*model.Note, len(notePb.Notes))
	for i, note := range notePb.Notes {
		notes[i] = &model.Note{
			ID:      note.Id,
			Title:   note.Title,
			Content: note.Content,
		}
	}

	return notes, nil
}
