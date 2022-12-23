// Code generated by goa v3.5.2, DO NOT EDIT.
//
// calc service
//
// Command:
// $ goa gen goa-test/design

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// The calc service performs operations on numbers.
type Service interface {
	// Multiply implements multiply.
	Multiply(context.Context, *MultiplyPayload) (res int, err error)
	// Add implements add.
	Add(context.Context, *AddPayload) (res int, err error)
	// Subtract implements subtract.
	Subtract(context.Context, *SubtractPayload) (res int, err error)
	// Divide implements divide.
	Divide(context.Context, *DividePayload) (res float64, err error)
	// GetNotes implements getNotes.
	GetNotes(context.Context, *GetNotesPayload) (res *GetNotesResult, err error)
	// GetNote implements getNote.
	GetNote(context.Context, *GetNotePayload) (res *GetNoteResult, err error)
	// CreateNote implements createNote.
	CreateNote(context.Context, *CreateNotePayload) (res *CreateNoteResult, err error)
	// DeleteNote implements deleteNote.
	DeleteNote(context.Context, *DeleteNotePayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "calc"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [8]string{"multiply", "add", "subtract", "divide", "getNotes", "getNote", "createNote", "deleteNote"}

// MultiplyPayload is the payload type of the calc service multiply method.
type MultiplyPayload struct {
	// Left operand
	A int
	// Right operand
	B int
}

// AddPayload is the payload type of the calc service add method.
type AddPayload struct {
	// Left operand
	A int
	// Right operand
	B int
}

// SubtractPayload is the payload type of the calc service subtract method.
type SubtractPayload struct {
	// Left operand
	A int
	// Right operand
	B int
}

// DividePayload is the payload type of the calc service divide method.
type DividePayload struct {
	// Left operand
	A int
	// Right operand
	B int
}

// GetNotesPayload is the payload type of the calc service getNotes method.
type GetNotesPayload struct {
	// The email of the user
	UserID string
}

// GetNotesResult is the result type of the calc service getNotes method.
type GetNotesResult struct {
	// list of notes
	Notes []*Note
}

// GetNotePayload is the payload type of the calc service getNote method.
type GetNotePayload struct {
	// The note's UUID
	UUID string
}

// GetNoteResult is the result type of the calc service getNote method.
type GetNoteResult struct {
	// The note matching the UUID
	Note *Note
}

// CreateNotePayload is the payload type of the calc service createNote method.
type CreateNotePayload struct {
	// The UserID for the note
	UserID string
	// The Note data to be saved
	NoteDetails *NoteDetails
}

// CreateNoteResult is the result type of the calc service createNote method.
type CreateNoteResult struct {
	// The Created Note
	NoteResponse *NoteResponse
}

// DeleteNotePayload is the payload type of the calc service deleteNote method.
type DeleteNotePayload struct {
	// The uuid for the note
	UUID string
}

type Note struct {
	// The title of the Note
	Title *string
	// The Body of the Note
	Body *string
	// The UUID of the created note
	UUID *string
}

type NoteDetails struct {
	// The title of the Note
	Title *string
	// The Body of the Note
	Body *string
}

type NoteResponse struct {
	// The UUID of the Created Note
	UUID *string
}

// MakeNoteMissing builds a goa.ServiceError from an error.
func MakeNoteMissing(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "NoteMissing",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "BadRequest",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
