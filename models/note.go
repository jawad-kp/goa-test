package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"goa-test/gen/calc"
	"time"
	"github.com/google/uuid"
)

type Note struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
	UserID    string    `json:"user_id" bson:"user_id"`
	Title     string    `json:"title" bson:"title"`
	Body      string    `json:"body,omitempty" bson:"body"`
	UUID string `json:"uuid" bson:"uuid"`
}

func (nt *Note) MarshalBSON() ([]byte, error) {
	if nt.CreatedAt.IsZero() {
		nt.CreatedAt = time.Now()
	}
	if nt.UUID == "" {
		nt.UUID = uuid.NewString()
	}
	type adapt Note
	return bson.Marshal((*adapt)(nt))
}

func (nt *Note) setFromServiceNote(serviceNote calc.Note) {
	if serviceNote.Title != nil {
		nt.Title = *serviceNote.Title
	}
	if serviceNote.Body != nil {
		nt.Body = *serviceNote.Body
	}
}

func GetNewNote(payloadData calc.CreateNotePayload) (retVal Note) {
	retVal.UserID = payloadData.UserID
	if payloadData.Note != nil {
		retVal.setFromServiceNote(*payloadData.Note)
	}
	return
}
