package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"goa-test/models"
)

type MongoService struct {
	collection *mongo.Collection
}

func NewMongoService(collection *mongo.Collection) *MongoService {
	return &MongoService{collection: collection}
}

func (m MongoService) SaveNoteToDb(ctx context.Context, data *models.Note) (resp models.Note, err error) {
	_, err = m.collection.InsertOne(ctx, data)
	resp = *data
	return
}
