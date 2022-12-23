package calcapi

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"goa-test/db"
	"goa-test/gen/calc"
	"goa-test/models"
	"log"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcService struct {
	logger *log.Logger
	db     db.DB
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger, db db.DB) calc.Service {
	return &calcService{logger: logger, db: db}
}

// Multiply implements multiply.
func (s *calcService) Multiply(_ context.Context, p *calc.MultiplyPayload) (res int, err error) {
	res = p.A * p.B
	return
}

// Add implements add.
func (s *calcService) Add(_ context.Context, p *calc.AddPayload) (res int, err error) {
	res = p.A + p.B
	return
}

// Subtract implements subtract.
func (s *calcService) Subtract(_ context.Context, p *calc.SubtractPayload) (res int, err error) {
	res = p.A - p.B
	return
}

// Divide implements divide.
func (s *calcService) Divide(_ context.Context, p *calc.DividePayload) (res float64, err error) {
	res = float64(p.A) / float64(p.B)
	return
}

// GetNotes implements getNotes.
func (s *calcService) GetNotes(ctx context.Context, payload *calc.GetNotesPayload) (res *calc.GetNotesResult, err error) {
	result, err := s.db.GetNoteByUserID(ctx, payload.UserID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, calc.MakeNoteMissing(errors.New("no match found"))
		}
		return
	}
	if len(result) == 0 {
		return nil, calc.MakeNoteMissing(errors.New("no notes available"))
	}
	noteList := make([]*calc.Note, len(result))
	for i, note := range result {
		noteList[i] = &calc.Note{
			Title: &note.Title,
			Body:  &note.Body,
			UUID:  &note.UUID,
		}
	}
	res = &calc.GetNotesResult{Notes: noteList}
	return
}

func (s *calcService) GetNote(ctx context.Context, payload *calc.GetNotePayload) (res *calc.GetNoteResult, err error) {
	result, err := s.db.GetNoteByUUID(ctx, payload.UUID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, calc.MakeNoteMissing(errors.New("no match found"))
		}
		return
	}
	res = &calc.GetNoteResult{
		Note: &calc.Note{
			Title: &result.Title,
			Body:  &result.Body,
		},
	}
	return
}

func (s *calcService) CreateNote(ctx context.Context, payload *calc.CreateNotePayload) (res *calc.CreateNoteResult, err error) {
	noteModel := models.GetNewNote(*payload)
	s.logger.Print("calc.createNote")
	noteModel, err = s.db.SaveNoteToDb(ctx, &noteModel)
	res = &calc.CreateNoteResult{NoteResponse: &calc.NoteResponse{UUID: &noteModel.UUID}}
	return
}
