package repo

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/kil-san/micro-serv/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepo struct {
	client *mongo.Client
}

var Database = os.Getenv("DATABASE")
var NoteCollection = "notes"

func NewMongoRepo(client *mongo.Client) Repo {
	return &mongoRepo{
		client: client,
	}
}

func toDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

func (r *mongoRepo) getCollectionRef() *mongo.Collection {
	collection := r.client.Database(Database).Collection(NoteCollection)
	return collection
}

func (r *mongoRepo) Create(ctx context.Context, ownerUid string, data model.Note) error {
	collection := r.getCollectionRef()

	doc, err := toDoc(data)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err = collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepo) Get(ctx context.Context, ownerUid string, noteID string) (model.Note, error) {
	var note model.Note
	collection := r.getCollectionRef()

	filter := bson.D{
		{"id", noteID},
		{"ownerUid", ownerUid},
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(&note)
	if err == mongo.ErrNoDocuments {
		return note, fmt.Errorf("record does not exist")
	} else if err != nil {
		return note, err
	}

	return note, nil
}

func (r *mongoRepo) Delete(ctx context.Context, ownerUid string, noteID string) error {
	collection := r.getCollectionRef()
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})

	filter := bson.D{
		{"id", noteID},
		{"ownerUid", ownerUid},
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := collection.DeleteOne(ctx, filter, opts)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepo) Update(ctx context.Context, ownerUid string, data model.Note) error {
	collection := r.getCollectionRef()
	opts := options.Update().SetUpsert(false)

	doc, err := toDoc(data)
	if err != nil {
		return err
	}

	filter := bson.D{
		{"id", data.Id},
		{"ownerUid", ownerUid},
	}

	update := bson.D{{"$set", doc}}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	result, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	if result.UpsertedCount != 0 {
		return fmt.Errorf("record does not exist")
	}

	return nil
}

func (r *mongoRepo) List(ctx context.Context, ownerUid string, page uint32) ([]model.Note, error) {
	var notes []model.Note
	collection := r.getCollectionRef()

	findOptions := options.Find()
	findOptions.SetLimit(10)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return notes, err
	}

	for cur.Next(context.TODO()) {
		var note model.Note
		err := cur.Decode(&note)
		if err != nil {
			return notes, err
		}

		notes = append(notes, note)
	}

	if err := cur.Err(); err != nil {
		return notes, err
	}

	cur.Close(context.TODO())

	return notes, nil
}
