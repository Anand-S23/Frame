package storage

import (
	"context"
	"time"

	"github.com/Anand-S23/frame/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStore struct {
	db *mongo.Database
}

func NewMongoStore(db *mongo.Database) *MongoStore {
    return &MongoStore{db: db}
}

// Auth // 

func (store *MongoStore) CreateUser(user models.User) (*mongo.InsertOneResult, error) {
    usersCollection := store.db.Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()

    return usersCollection.InsertOne(ctx, user)
}

func (store *MongoStore) FindUserByField(field string, value string) *models.User {
    var user models.User
    usersCollection := store.db.Collection("users")

    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()

    err := usersCollection.FindOne(ctx, bson.M{field: value}).Decode(&user)
    if err == mongo.ErrNoDocuments {
        return nil
    }
    
    return &user
}

func (store *MongoStore) FindUserByID(id string) *models.User {
    return store.FindUserByField("_id", id)
}

func (store *MongoStore) FindUserByUsername(username string) *models.User {
    return store.FindUserByField("username", username)
}

func (store *MongoStore) FindUserByEmail(email string) *models.User {
    return store.FindUserByField("email", email)
}

func (store *MongoStore) GetAllUsers() *[]models.User {
    usersCollection := store.db.Collection("users")

    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()

    cursor, err := usersCollection.Find(ctx, bson.M{})
	if err != nil {
        return nil
	}
	defer cursor.Close(ctx)

    var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
            return nil
		}
		users = append(users, user)
	}

    return &users
}

