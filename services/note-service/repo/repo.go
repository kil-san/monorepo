package repo

import (
	"context"

	"github.com/kil-san/micro-serv/pkg/model"
)

type Repo interface {
	Create(ctx context.Context, data model.Note) (model.Note, error)
	Get(ctx context.Context, id string) (model.Note, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, data model.Note) error
}
