package db

import (
	"context"
	"goa-test/models"
)

type DB interface {
	SaveNoteToDb(ctx context.Context, data *models.Note) (resp models.Note, err error)
	GetNoteByUUID(ctx context.Context, uuid string) (resp models.Note, err error)
	GetNoteByUserID(ctx context.Context, userID string) (resp []models.Note, err error)
	SoftDeleteNoteByUUID(ctx context.Context, uuid string) (err error)
}
