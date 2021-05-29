package rpc

import (
	"context"
	"database/sql"

	"github.com/kil-san/micro-serv/note-service/pb"
	"github.com/kil-san/micro-serv/note-service/repo"
	"github.com/kil-san/micro-serv/note-service/service"
	"github.com/kil-san/micro-serv/pkg/model"
	"google.golang.org/protobuf/types/known/emptypb"
	log "unknwon.dev/clog/v2"
)

type noteRPCServer struct {
	pb.UnimplementedNoteRPCServer
	db *sql.DB
}

func New(db *sql.DB) pb.NoteRPCServer {
	s := &noteRPCServer{
		db: db,
	}
	return s
}

func GetService(db *sql.DB) *service.NoteService {
	repo := repo.NewSqlRepo(db)
	svc := service.NewNoteService(repo)

	return svc
}

func (s *noteRPCServer) CreateNote(ctx context.Context, req *pb.Note) (*pb.Note, error) {
	var res pb.Note

	svc := GetService(s.db)
	note, err := svc.CreateNote(ctx, model.Note{
		Title:  req.Title,
		Status: req.Status,
	})
	if err != nil {
		log.Error("%+v", err)
		return &res, err
	}

	res.Id = note.Id
	res.Title = note.Title
	res.Status = note.Status

	return &res, nil
}

func (s *noteRPCServer) GetNotes(ctx context.Context, req *pb.OwnerUid) (*pb.NoteList, error) {
	return &pb.NoteList{}, nil
}

func (s *noteRPCServer) GetNote(ctx context.Context, req *pb.SingleNote) (*pb.Note, error) {
	var res pb.Note

	svc := GetService(s.db)
	note, err := svc.GetNote(ctx, req.NoteId)
	if err != nil {
		log.Error("%+v", err)
		return &res, err
	}

	res.Id = note.Id
	res.Title = note.Title
	res.Status = note.Status

	return &res, nil
}

func (s *noteRPCServer) UpdateNote(ctx context.Context, req *pb.Note) (*emptypb.Empty, error) {
	var res emptypb.Empty

	svc := GetService(s.db)
	err := svc.UpdateNote(ctx, req.Id, model.Note{
		Title:  req.Title,
		Status: req.Status,
	})
	if err != nil {
		log.Error("%+v", err)
		return &res, err
	}

	return &res, nil
}

func (s *noteRPCServer) DeleteNote(ctx context.Context, req *pb.SingleNote) (*emptypb.Empty, error) {
	var res emptypb.Empty

	svc := GetService(s.db)
	err := svc.DeleteNote(ctx, req.NoteId)
	if err != nil {
		log.Error("%+v", err)
		return &res, err
	}

	return &res, nil
}
