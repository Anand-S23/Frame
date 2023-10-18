package storage

import (
	"github.com/Anand-S23/frame/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStore struct {
	db *mongo.Database
}

func NewMongoStore(db *mongo.Database) *MongoStore {
    return &MongoStore{db: db}
}

// Auth // 

func (store *MongoStore) CreateUser(user models.User) error {

    return nil
}

