package model

import "go.mongodb.org/mongo-driver/v2/bson"

type PostRequest struct {
	Content string `form:"content" bson:"content" binding:"required"`
}

type Post struct {
	Id       bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AuthorId string        `form:"author_id" bson:"author_id" binding:"required"`
	Content  string        `form:"content" bson:"content" binding:"required"`
}
