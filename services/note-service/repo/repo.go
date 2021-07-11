package repo

import (
	"context"

	"github.com/kil-san/micro-serv/pkg/model"
)

type Repo interface {
	Create(ctx context.Context, ownerUid string, data model.Note) error
	Get(ctx context.Context, ownerUid string, noteID string) (model.Note, error)
	Delete(ctx context.Context, ownerUid string, noteID string) error
	Update(ctx context.Context, ownerUid string, data model.Note) error
	List(ctx context.Context, ownerUid string, page uint32) ([]model.Note, error)
}
