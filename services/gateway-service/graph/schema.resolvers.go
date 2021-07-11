package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/kil-san/micro-serv/gateway-service/graph/generated"
	"github.com/kil-san/micro-serv/gateway-service/graph/model"
	"github.com/kil-san/micro-serv/gateway-service/service"
	"github.com/kil-san/micro-serv/note-service/pb"
	"google.golang.org/grpc"
	log "unknwon.dev/clog/v2"
)

func (r *mutationResolver) CreateNote(ctx context.Context, data model.NewNote) (*model.Note, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	conn, err := grpc.DialContext(ctx, service.NoteRPCEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to dial server:", err)
	}

	defer conn.Close()
	client := pb.NewNoteRPCClient(conn)
	svc := service.NewNoteService(client)
	var note model.Note
	note, err = svc.CreateNote(ctx, "default", data)
	if err != nil {
		log.Error("%+v\n", err)
		return &note, err
	}

	return &note, nil
}

func (r *mutationResolver) UpdateNote(ctx context.Context, data model.NoteUpdate) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	conn, err := grpc.DialContext(ctx, service.NoteRPCEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to dial server:", err)
	}

	defer conn.Close()
	client := pb.NewNoteRPCClient(conn)
	svc := service.NewNoteService(client)
	err = svc.UpdateNote(ctx, "default", data)
	if err != nil {
		log.Error("%+v\n", err)
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) DeleteNote(ctx context.Context, data string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	conn, err := grpc.DialContext(ctx, service.NoteRPCEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to dial server:", err)
	}

	defer conn.Close()
	client := pb.NewNoteRPCClient(conn)
	svc := service.NewNoteService(client)
	err = svc.DeleteNote(ctx, "default", data)
	if err != nil {
		log.Error("%+v\n", err)
		return false, err
	}

	return true, nil
}

func (r *queryResolver) GetNotes(ctx context.Context, data int) ([]*model.Note, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	conn, err := grpc.DialContext(ctx, service.NoteRPCEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to dial server:", err)
	}

	defer conn.Close()
	client := pb.NewNoteRPCClient(conn)
	svc := service.NewNoteService(client)
	var notes []*model.Note
	notes, err = svc.GetNotes(ctx, "default", uint32(data))
	if err != nil {
		log.Error("%+v\n", err)
		return notes, err
	}

	return notes, nil
}

func (r *queryResolver) GetNote(ctx context.Context, data string) (*model.Note, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	conn, err := grpc.DialContext(ctx, service.NoteRPCEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to dial server:", err)
	}

	defer conn.Close()
	client := pb.NewNoteRPCClient(conn)
	svc := service.NewNoteService(client)
	var note model.Note
	note, err = svc.GetNote(ctx, "default", data)
	if err != nil {
		log.Error("%+v\n", err)
		return &note, err
	}

	return &note, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
