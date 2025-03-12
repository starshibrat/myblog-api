package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	Id        bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string        `bson:"username"`
	Password  string        `bson:"password"`
	Email     string        `bson:"email"`
	CreatedAt time.Time     `bson:"created_at"`
}

type UserResponse struct {
	Id        bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string        `bson:"username"`
	CreatedAt time.Time     `bson:"created_at"`
}

type UserId struct {
	Id string `form:"id" json:"id" binding:"required"`
}
