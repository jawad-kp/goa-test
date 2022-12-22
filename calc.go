package calcapi

import (
	"context"
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
func (s *calcService) GetNotes(_ context.Context, _ *calc.GetNotesPayload) (res *calc.GetNotesResult, err error) {
	res = &calc.GetNotesResult{}
	s.logger.Print("calc.getNotes")
	return
}

func (s *calcService) CreateNote(ctx context.Context, payload *calc.CreateNotePayload) (err error) {
	noteModel := models.GetNewNote(*payload)
	s.logger.Print("calc.createNote")
	noteModel, err = s.db.SaveNoteToDb(ctx, &noteModel)
	return
}
