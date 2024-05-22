package store

import (
	"context"

	"github.com/natchapon/todoapi/todo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBStore struct {
	*mongo.Collection
}

func NewMongoDBStore() *MongoDBStore {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017"))
	if err != nil {
		panic("failed to connect database")
	}
	collection := client.Database("myapp").Collection("todos")

	return &MongoDBStore{collection}
}

func (s *MongoDBStore) New(todo *todo.Todo) error {
	_, err := s.Collection.InsertOne(context.Background(), todo)
	return err
}
