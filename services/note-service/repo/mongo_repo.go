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

type MongoRepo struct {
	client *mongo.Client
}

var Database = os.Getenv("DATABASE")
var NoteCollection = "notes"

func NewMongoRepo(client *mongo.Client) MongoRepo {
	return MongoRepo{
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

func (r MongoRepo) getCollectionRef() *mongo.Collection {
	collection := r.client.Database(Database).Collection(NoteCollection)
	return collection
}

func (r MongoRepo) Create(ctx context.Context, data model.Note) (model.Note, error) {
	var note model.Note
	collection := r.getCollectionRef()

	doc, err := toDoc(data)
	if err != nil {
		return note, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = collection.InsertOne(ctx, doc)
	if err != nil {
		return note, err
	}

	return data, nil
}

func (r MongoRepo) Get(ctx context.Context, id string) (model.Note, error) {
	var note model.Note
	collection := r.getCollectionRef()

	filter := bson.D{{"id", id}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(&note)
	if err == mongo.ErrNoDocuments {
		return note, fmt.Errorf("record does not exist")
	} else if err != nil {
		return note, err
	}

	return note, nil
}

func (r MongoRepo) Delete(ctx context.Context, id string) error {
	collection := r.getCollectionRef()
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.DeleteOne(ctx, bson.D{{"id", id}}, opts)
	if err != nil {
		return err
	}

	return nil
}

func (r MongoRepo) Update(ctx context.Context, id string, data model.Note) error {
	collection := r.getCollectionRef()
	opts := options.Update().SetUpsert(false)

	doc, err := toDoc(data)
	if err != nil {
		return err
	}

	filter := bson.D{{"id", id}}
	update := bson.D{{"$set", doc}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
