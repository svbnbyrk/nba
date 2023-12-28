package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
	DB     *mongo.Database // Added to hold a reference to the db
}

func New(ctx context.Context, url string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("cannot connect mongo client err: %w", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("mongo ping err: %w", err)
	}

	db := client.Database(os.Getenv("DB_NAME"))

	return &MongoDB{
		Client: client,
		DB:     db,
	}, nil
}

func (p *MongoDB) Close() {
	_ = p.Client.Disconnect(context.Background())
}
