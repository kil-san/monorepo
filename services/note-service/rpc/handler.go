package rpc

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/kil-san/micro-serv/note-service/connection"
	"github.com/kil-san/micro-serv/note-service/pb"
	"github.com/kil-san/micro-serv/note-service/repo"
	"github.com/kil-san/micro-serv/note-service/service"
	"github.com/kil-san/micro-serv/pkg/model"
	"google.golang.org/protobuf/types/known/emptypb"
	log "unknwon.dev/clog/v2"
)

type noteRPCServer struct {
	pb.UnimplementedNoteRPCServer
}

func New() pb.NoteRPCServer {
	s := &noteRPCServer{}
	return s
}

func GetService(client *firestore.Client) *service.NoteService {
	repo := repo.NewNoteRepo(context.Background(), client)
	svc := service.NewNoteService(repo)

	return svc
}

func (s *noteRPCServer) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.Note, error) {
	var res pb.Note
	client := connection.NewFirestoreClient(ctx)
	defer client.Close()

	svc := GetService(client)
	note, err := svc.CreateNote(ctx, req.OwnerUid, model.Note{
		Title:   req.Note.Title,
		Content: req.Note.Content,
	})
	if err != nil {
		log.Error("%+v", err)
		return &res, err
	}

	res.Id = note.Id
	res.Title = note.Title
	res.Content = note.Content

	return &res, nil
}

func (s *noteRPCServer) GetNotes(ctx context.Context, req *pb.GetNotesRequest) (*pb.NoteList, error) {
	var res pb.NoteList
	client := connection.NewFirestoreClient(ctx)
	defer client.Close()

	svc := GetService(client)
	notes, err := svc.GetNotes(ctx, req.OwnerUid, uint32(req.Page))
	if err != nil {
		log.Error("%+v", err)
		return &res, err
	}

	res.Notes = make([]*pb.Note, len(notes))
	for i, note := range notes {
		res.Notes[i] = &pb.Note{
			Id:      note.Id,
			Title:   note.Title,
			Content: note.Content,
		}
	}

	return &res, nil
}

func (s *noteRPCServer) GetNote(ctx context.Context, req *pb.SimpleRequest) (*pb.Note, error) {
	var res pb.Note
	client := connection.NewFirestoreClient(ctx)
	defer client.Close()

	svc := GetService(client)
	note, err := svc.GetNote(ctx, req.OwnerUid, req.NoteId)
	if err != nil {
		log.Error("%+v", err)
		return &res, err
	}

	res.Id = note.Id
	res.Title = note.Title
	res.Content = note.Content

	return &res, nil
}

func (s *noteRPCServer) UpdateNote(ctx context.Context, req *pb.CreateNoteRequest) (*emptypb.Empty, error) {
	var res emptypb.Empty
	client := connection.NewFirestoreClient(ctx)
	defer client.Close()

	svc := GetService(client)
	err := svc.UpdateNote(ctx, req.OwnerUid, model.Note{
		Id:      req.Note.Id,
		Title:   req.Note.Title,
		Content: req.Note.Content,
	})
	if err != nil {
		log.Error("%+v", err)
		return &res, err
	}

	return &res, nil
}

func (s *noteRPCServer) DeleteNote(ctx context.Context, req *pb.SimpleRequest) (*emptypb.Empty, error) {
	var res emptypb.Empty
	client := connection.NewFirestoreClient(ctx)
	defer client.Close()

	svc := GetService(client)
	err := svc.DeleteNote(ctx, req.OwnerUid, req.NoteId)
	if err != nil {
		log.Error("%+v", err)
		return &res, err
	}

	return &res, nil
}
