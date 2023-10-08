package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
    DatabaseName string
    Context context.Context
    Client *mongo.Client

    Collections map[string]*mongo.Collection
}

func InitDB(uri string, dbName string, collectionNames []string) (*MongoDB, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
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

    collections := make(map[string]*mongo.Collection)
	for _, name := range collectionNames {
		collections[name] = client.Database(dbName).Collection(name)
	}

    return &MongoDB {
        DatabaseName: dbName,
        Context: ctx,
        Client: client,
        Collections: collections,
    }, nil
}

func (db *MongoDB) InsertData(collectionName string, data bson.M) (*mongo.InsertOneResult, error) {
    collection, ok := db.Collections[collectionName]
    if !ok {
        log.Fatalf("%s collection does not exist\n", collectionName)
    }

    return collection.InsertOne(context.Background(), data)
}

