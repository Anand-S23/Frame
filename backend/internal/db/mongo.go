package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB(uri string, dbName string, timeout time.Duration) (*mongo.Database, error) {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    opts := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(ctx, opts)
    if err != nil {
        return nil, err
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        return nil, err
    }

    return client.Database(dbName), nil
}

func CloseDB(db *mongo.Database) error {
    return db.Client().Disconnect(context.Background())
}
