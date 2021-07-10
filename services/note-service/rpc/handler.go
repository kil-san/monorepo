package rpc

import (
	"context"
	"os"

	"github.com/kil-san/micro-serv/note-service/connection"
	"github.com/kil-san/micro-serv/note-service/pb"
	"github.com/kil-san/micro-serv/note-service/repo"
	"github.com/kil-san/micro-serv/note-service/service"
	"github.com/kil-san/micro-serv/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
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

func GetService(client *mongo.Client) *service.NoteService {
	mongoRepo := repo.NewMongoRepo(client)
	svc := service.NewNoteService(mongoRepo)

	return svc
}

func (s *noteRPCServer) CreateNote(ctx context.Context, req *pb.Note) (*pb.Note, error) {
	var res pb.Note

	ctx, client, err := connection.NewMongoConnection(ctx, os.Getenv("MONGO_DB_HOST"), os.Getenv("MONGO_DB_PORT"))
	if err != nil {
		log.Fatal("could not open connection to db: %+v", err)
	}
	defer client.Disconnect(ctx)

	svc := GetService(client)
	note, err := svc.CreateNote(ctx, model.Note{
		Title:   req.Title,
		Content: req.Content,
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

func (s *noteRPCServer) GetNotes(ctx context.Context, req *pb.OwnerUid) (*pb.NoteList, error) {
	var res pb.NoteList

	ctx, client, err := connection.NewMongoConnection(ctx, os.Getenv("MONGO_DB_HOST"), os.Getenv("MONGO_DB_PORT"))
	if err != nil {
		log.Fatal("could not open connection to db: %+v", err)
	}
	defer client.Disconnect(ctx)

	svc := GetService(client)
	notes, err := svc.GetNotes(ctx, req.UserId)
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

func (s *noteRPCServer) GetNote(ctx context.Context, req *pb.SingleNote) (*pb.Note, error) {
	var res pb.Note

	ctx, client, err := connection.NewMongoConnection(ctx, os.Getenv("MONGO_DB_HOST"), os.Getenv("MONGO_DB_PORT"))
	if err != nil {
		log.Fatal("could not open connection to db: %+v", err)
	}
	defer client.Disconnect(ctx)

	svc := GetService(client)
	note, err := svc.GetNote(ctx, req.NoteId)
	if err != nil {
		log.Error("%+v", err)
		return &res, err
	}

	res.Id = note.Id
	res.Title = note.Title
	res.Content = note.Content

	return &res, nil
}

func (s *noteRPCServer) UpdateNote(ctx context.Context, req *pb.Note) (*emptypb.Empty, error) {
	var res emptypb.Empty

	ctx, client, err := connection.NewMongoConnection(ctx, os.Getenv("MONGO_DB_HOST"), os.Getenv("MONGO_DB_PORT"))
	if err != nil {
		log.Fatal("could not open connection to db: %+v", err)
	}
	defer client.Disconnect(ctx)

	svc := GetService(client)
	err = svc.UpdateNote(ctx, req.Id, model.Note{
		Id:      req.Id,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		log.Error("%+v", err)
		return &res, err
	}

	return &res, nil
}

func (s *noteRPCServer) DeleteNote(ctx context.Context, req *pb.SingleNote) (*emptypb.Empty, error) {
	var res emptypb.Empty

	ctx, client, err := connection.NewMongoConnection(ctx, os.Getenv("MONGO_DB_HOST"), os.Getenv("MONGO_DB_PORT"))
	if err != nil {
		log.Fatal("could not open connection to db: %+v", err)
	}
	defer client.Disconnect(ctx)

	svc := GetService(client)
	err = svc.DeleteNote(ctx, req.NoteId)
	if err != nil {
		log.Error("%+v", err)
		return &res, err
	}

	return &res, nil
}
