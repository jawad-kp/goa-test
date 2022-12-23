// Code generated by goa v3.5.2, DO NOT EDIT.
//
// calc endpoints
//
// Command:
// $ goa gen goa-test/design

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "calc" service endpoints.
type Endpoints struct {
	Multiply   goa.Endpoint
	Add        goa.Endpoint
	Subtract   goa.Endpoint
	Divide     goa.Endpoint
	GetNotes   goa.Endpoint
	GetNote    goa.Endpoint
	CreateNote goa.Endpoint
}

// NewEndpoints wraps the methods of the "calc" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Multiply:   NewMultiplyEndpoint(s),
		Add:        NewAddEndpoint(s),
		Subtract:   NewSubtractEndpoint(s),
		Divide:     NewDivideEndpoint(s),
		GetNotes:   NewGetNotesEndpoint(s),
		GetNote:    NewGetNoteEndpoint(s),
		CreateNote: NewCreateNoteEndpoint(s),
	}
}

// Use applies the given middleware to all the "calc" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Multiply = m(e.Multiply)
	e.Add = m(e.Add)
	e.Subtract = m(e.Subtract)
	e.Divide = m(e.Divide)
	e.GetNotes = m(e.GetNotes)
	e.GetNote = m(e.GetNote)
	e.CreateNote = m(e.CreateNote)
}

// NewMultiplyEndpoint returns an endpoint function that calls the method
// "multiply" of service "calc".
func NewMultiplyEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*MultiplyPayload)
		return s.Multiply(ctx, p)
	}
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "calc".
func NewAddEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AddPayload)
		return s.Add(ctx, p)
	}
}

// NewSubtractEndpoint returns an endpoint function that calls the method
// "subtract" of service "calc".
func NewSubtractEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SubtractPayload)
		return s.Subtract(ctx, p)
	}
}

// NewDivideEndpoint returns an endpoint function that calls the method
// "divide" of service "calc".
func NewDivideEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DividePayload)
		return s.Divide(ctx, p)
	}
}

// NewGetNotesEndpoint returns an endpoint function that calls the method
// "getNotes" of service "calc".
func NewGetNotesEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetNotesPayload)
		return s.GetNotes(ctx, p)
	}
}

// NewGetNoteEndpoint returns an endpoint function that calls the method
// "getNote" of service "calc".
func NewGetNoteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetNotePayload)
		return s.GetNote(ctx, p)
	}
}

// NewCreateNoteEndpoint returns an endpoint function that calls the method
// "createNote" of service "calc".
func NewCreateNoteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateNotePayload)
		return s.CreateNote(ctx, p)
	}
}
