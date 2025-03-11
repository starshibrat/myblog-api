package model

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	Id       bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string        `json:"username"`
	Password string        `json:"password"`
	Email    string        `json:"email"`
}

type UserResponse struct {
	Id       bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string        `json:"username"`
}

type UserId struct {
	Id string `form:"id" json:"id" binding:"required"`
}
