package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGODB_URI      = "MONGODB_URI"
	MONGODB_DATABASE = "MONGODB_DATABASE"
)

func NewMongoConnection(ctx context.Context) (*mongo.Database, error) {
	mongoURI := os.Getenv(MONGODB_URI)
	mongoDatabase := os.Getenv(MONGODB_DATABASE)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongoDatabase), nil
}
