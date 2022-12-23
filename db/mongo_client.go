package db

import (
	"context"
	"errors"
	"goa-test/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	filter := bson.M{
		"user_id":    userID,
		"deleted_at": bson.M{"$eq": time.Time{}},
	}
	data, err = m.collection.Find(ctx, filter)
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
	filter := bson.M{
		"uuid":       uuid,
		"deleted_at": bson.M{"$eq": time.Time{}},
	}
	err = m.collection.FindOne(ctx, filter).Decode(&resp)
	return
}

func (m MongoService) SoftDeleteNoteByUUID(ctx context.Context, uuid string) (
	err error) {
	filter := bson.M{
		"uuid":       uuid,
		"deleted_at": bson.M{"$eq": time.Time{}},
	}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "deleted_at", Value: time.Now()}}}}
	res, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return
	}
	if res.ModifiedCount == 0 {
		return errors.New("document not found")
	}
	return
}
