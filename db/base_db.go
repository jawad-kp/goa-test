package db

import (
	"context"
	"goa-test/models"
)

type DB interface {
	SaveNoteToDb(ctx context.Context, data *models.Note) (resp models.Note, err error)
}
