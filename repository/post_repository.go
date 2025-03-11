package repository

import (
	"context"
	"example/blog-service-gin/model"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type PostRepository interface {
	Create(ctx context.Context, data *model.Post) (interface{}, error)
}

type postRepository struct {
	store model.DbStore
}

func (r *postRepository) Create(ctx context.Context, data *model.Post) (interface{}, error) {
	client, err := r.store.GetClient()

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	db := client.Database("myblog")

	objId, err := bson.ObjectIDFromHex(data.AuthorId)

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	filter := bson.M{"_id": objId}

	var user model.User

	err = db.Collection("user").FindOne(ctx, filter).Decode(&user)

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	new_post := bson.M{"author_id": user.Id, "content": data.Content}

	res, err := db.Collection("posts").InsertOne(ctx, new_post)

	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil

}

func NewPostRepository(store model.DbStore) PostRepository {
	return &postRepository{store: store}
}
