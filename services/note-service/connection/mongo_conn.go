package connection

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"unknwon.dev/clog/v2"
)

func NewMongoConnection(ctx context.Context, host string, port string) (context.Context, *mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", host, port)

	clog.Trace("===== CONNECT URI: %s", uri)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return ctx, nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return ctx, nil, err
	}
	clog.Trace("Successfully connected and pinged db.")

	return ctx, client, nil
}
