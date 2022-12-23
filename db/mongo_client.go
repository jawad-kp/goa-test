package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
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

func (m MongoService) GetNoteByUserID(ctx context.Context, userID string) (resp []models.Note, err error) {
	var data *mongo.Cursor
	data, err = m.collection.Find(ctx, bson.D{{"user_id", userID}})
	if err != nil {
		return
	}
	if data == nil {
		return
	}
	resp = make([]models.Note, 0)
	err = data.All(ctx, &resp)
	return
}

func (m MongoService) GetNoteByUUID(ctx context.Context, uuid string) (resp models.Note, err error) {
	err = m.collection.FindOne(ctx, bson.D{{"uuid", uuid}}).Decode(&resp)
	return
}
