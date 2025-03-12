package repository

import (
	"context"
	"example/blog-service-gin/model"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type PostRepository interface {
	Create(ctx context.Context, data *model.Post) (interface{}, error)
	DeleteById(ctx context.Context, data *model.PostId, authorId string) (interface{}, error)
	GetAllPost(ctx context.Context) (interface{}, error)
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

	filter := bson.M{"_id": data.AuthorId}

	var user model.User

	err = db.Collection("user").FindOne(ctx, filter).Decode(&user)

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	new_post := bson.M{"author_id": user.Id, "content": data.Content, "created_at": data.CreatedAt}

	res, err := db.Collection("posts").InsertOne(ctx, new_post)

	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func (r *postRepository) DeleteById(ctx context.Context, data *model.PostId, authorId string) (interface{}, error) {
	client, err := r.store.GetClient()

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	db := client.Database("myblog")

	objId, err := bson.ObjectIDFromHex(authorId)

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	post_id, err := bson.ObjectIDFromHex(data.Id)

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	var post model.Post
	filter := bson.M{"_id": post_id}

	err = db.Collection("posts").FindOne(ctx, filter).Decode(&post)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	if objId != post.AuthorId {
		log.Printf("error: author id != post.author id")
		return nil, fmt.Errorf("you're not the author of the post")
	}

	res, err := db.Collection("posts").DeleteOne(ctx, bson.M{"_id": post.Id})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *postRepository) GetAllPost(ctx context.Context) (interface{}, error) {
	client, err := r.store.GetClient()

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	db := client.Database("myblog")

	var posts []model.Post

	cur, err := db.Collection("posts").Find(ctx, bson.D{{}})

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	for cur.Next(ctx) {
		var elem model.Post
		err := cur.Decode(&elem)

		if err != nil {
			log.Printf("error: %v", err)
			return nil, err
		}

		posts = append(posts, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return posts, nil

}

func NewPostRepository(store model.DbStore) PostRepository {
	return &postRepository{store: store}
}
