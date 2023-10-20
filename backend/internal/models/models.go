package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginDto struct {
    Email    string
    Password string
}

type UserDto struct {
    Username string
    Email    string
    Password string
}

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
    User_ID    string             `json:"user_id"`
    Username   string             `json:"username"`
    Email      string             `json:"email"`
    Password   string             `json:"-"`
	Created_At time.Time          `json:"created_at"`
	Updated_At time.Time          `json:"updated_at"`
    // Avatar  string             `json:"avatar"`
}


func NewUser(userData UserDto) User {
    id := primitive.NewObjectID()
    now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

    return User {
        ID: id,
        User_ID: id.Hex(),
        Username: userData.Username,
        Email: userData.Email,
        Password: userData.Password,
        Created_At: now,
        Updated_At: now,
    }
}


