package repo

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/kil-san/micro-serv/pkg/model"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const FirestoreNoteCollection = "notes"
const FirestoreNoteCountDocument = "count"

type noteRepo struct {
	db *firestore.Client
}

func NewNoteRepo(ctx context.Context, db *firestore.Client) Repo {
	return &noteRepo{
		db,
	}
}

func (r *noteRepo) getCollectionRef(ownerUid string) *firestore.CollectionRef {
	return r.db.Collection(FirestoreNoteCollection).Doc(ownerUid).Collection(FirestoreNoteCollection)
}

func (r *noteRepo) getNexusDocument(ctx context.Context, ownerUID string) *firestore.DocumentRef {
	return r.db.Collection(FirestoreNoteCollection).Doc(ownerUID)
}

func (r *noteRepo) getCountRef(ctx context.Context, ownerUID string) *firestore.CollectionRef {
	return r.getNexusDocument(ctx, ownerUID).Collection(FirestoreNoteCountDocument)
}

func (r *noteRepo) List(ctx context.Context, ownerUid string, page uint32) ([]model.Note, error) {
	offset := int((page - 1) * 10)
	iter := r.getCollectionRef(ownerUid).Offset(offset).Limit(10).Documents(ctx)
	var notes []model.Note

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		var order model.Note
		err = doc.DataTo(&order)
		if err != nil {
			return notes, err
		}

		notes = append(notes, order)
	}

	return notes, nil
}

func (r *noteRepo) Create(ctx context.Context, ownerUid string, data model.Note) error {
	_, err := r.getCollectionRef(ownerUid).Doc(data.Id).Set(ctx, data)
	if err != nil {
		return err
	}

	err = r.incrementCount(ctx, ownerUid, 1)

	return err
}

func (r *noteRepo) Update(ctx context.Context, ownerUid string, data model.Note) error {
	_, err := r.getCollectionRef(ownerUid).Doc(data.Id).Set(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (r *noteRepo) Delete(ctx context.Context, ownerUid string, noteID string) error {
	_, err := r.getCollectionRef(ownerUid).Doc(noteID).Delete(ctx)
	if err != nil {
		return err
	}

	err = r.incrementCount(ctx, ownerUid, -1)

	return err
}

func (r *noteRepo) Get(ctx context.Context, ownerUID string, noteID string) (model.Note, error) {
	var note model.Note
	data, err := r.getCollectionRef(ownerUID).Doc(noteID).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return note, nil
		}
		return note, err
	}

	err = data.DataTo(&note)
	if err != nil {
		return note, err
	}

	return note, nil
}

func (r *noteRepo) GetNoteCount(ctx context.Context, ownerUID string) (uint32, error) {
	var count uint32
	data, err := r.getCountRef(ctx, ownerUID).Doc(FirestoreNoteCountDocument).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return count, nil
		}
		return count, err
	}

	countRaw, err := data.DataAt("value")
	if err != nil {
		return count, err
	}

	count = uint32(countRaw.(int64))

	return count, nil
}

func (r *noteRepo) incrementCount(ctx context.Context, ownerUID string, value int32) error {
	type Counter struct {
		Value uint32 `firestore:"value"`
	}
	ref := r.getCountRef(ctx, ownerUID).Doc(FirestoreNoteCountDocument)
	_, err := ref.Update(ctx, []firestore.Update{{Path: "value", Value: firestore.Increment(value)}})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			_, err = ref.Set(ctx, Counter{
				Value: 1,
			})
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	return nil
}
