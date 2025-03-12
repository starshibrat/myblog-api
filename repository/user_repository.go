package repository

import (
	"context"
	"example/blog-service-gin/model"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Register(ctx context.Context, data *model.Register) (interface{}, error)
	Login(ctx context.Context, data *model.Login) (interface{}, error)
	AllUsers(ctx context.Context) (interface{}, error)
	DeleteById(ctx context.Context, user model.UserId) (interface{}, error)
}

type userRepository struct {
	store model.DbStore
}

func (r *userRepository) Login(ctx context.Context, data *model.Login) (interface{}, error) {
	client, err := r.store.GetClient()

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	var user model.User

	filter := bson.M{}

	err = client.Database("myblog").Collection("user").FindOne(ctx, filter).Decode(&user)

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	secret := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"email":    user.Password,
		"nbf":      time.Now().Unix(),
	})

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return nil, err
	}

	return tokenString, nil

}

func (r *userRepository) Register(ctx context.Context, data *model.Register) (interface{}, error) {
	client, err := r.store.GetClient()

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	var user model.User

	err = client.Database("myblog").Collection("user").FindOne(ctx, bson.M{"username": data.Username}).Decode(&user)

	if err == nil {
		return nil, fmt.Errorf("username has already taken")
	}

	err = client.Database("myblog").Collection("user").FindOne(ctx, bson.M{"email": data.Email}).Decode(&user)

	if err == nil {
		return nil, fmt.Errorf("email has already taken")
	}

	h, err := bcrypt.GenerateFromPassword([]byte(data.Password), 12)

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	data.Password = string(h)

	new_user := model.User{
		Username:  data.Username,
		Password:  data.Password,
		Email:     data.Email,
		CreatedAt: time.Now(),
	}

	res, err := client.Database("myblog").Collection("user").InsertOne(ctx, new_user)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (r *userRepository) AllUsers(ctx context.Context) (interface{}, error) {
	client, err := r.store.GetClient()

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	res, err := client.Database("myblog").Collection("user").Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	var results []model.UserResponse

	for res.Next(ctx) {
		var elem model.UserResponse
		err := res.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)

	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return results, nil

}

func (r *userRepository) DeleteById(ctx context.Context, user model.UserId) (interface{}, error) {
	client, err := r.store.GetClient()

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	objId, err := bson.ObjectIDFromHex(user.Id)

	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	res, err := client.Database("myblog").Collection("user").DeleteOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewUserRepository(store model.DbStore) UserRepository {
	return &userRepository{store: store}
}
