package model

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DbStore interface {
	// GetDb() *mongo.Database
	GetClient() (*mongo.Client, error)
	// Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
	Disconnect() error
}

type dbStore struct {
	client *mongo.Client
}

func NewDbStore(opts *options.ClientOptions) (DbStore, error) {
	client, err := mongo.Connect(opts)
	if err != nil {

		return nil, err
	}
	return &dbStore{client: client}, nil
}

func (md *dbStore) GetClient() (*mongo.Client, error) {
	if md.client != nil {
		return md.client, nil
	}

	return nil, errors.New("client is missing")
}

func (md *dbStore) Disconnect() error {
	return md.client.Disconnect(context.TODO())
}
