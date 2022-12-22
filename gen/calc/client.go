// Code generated by goa v3.10.2, DO NOT EDIT.
//
// calc client
//
// Command:
// $ goa gen goa-test/design

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "calc" service client.
type Client struct {
	MultiplyEndpoint   goa.Endpoint
	AddEndpoint        goa.Endpoint
	SubtractEndpoint   goa.Endpoint
	DivideEndpoint     goa.Endpoint
	GetNotesEndpoint   goa.Endpoint
	CreateNoteEndpoint goa.Endpoint
}

// NewClient initializes a "calc" service client given the endpoints.
func NewClient(multiply, add, subtract, divide, getNotes, createNote goa.Endpoint) *Client {
	return &Client{
		MultiplyEndpoint:   multiply,
		AddEndpoint:        add,
		SubtractEndpoint:   subtract,
		DivideEndpoint:     divide,
		GetNotesEndpoint:   getNotes,
		CreateNoteEndpoint: createNote,
	}
}

// Multiply calls the "multiply" endpoint of the "calc" service.
func (c *Client) Multiply(ctx context.Context, p *MultiplyPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.MultiplyEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// Add calls the "add" endpoint of the "calc" service.
func (c *Client) Add(ctx context.Context, p *AddPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// Subtract calls the "subtract" endpoint of the "calc" service.
func (c *Client) Subtract(ctx context.Context, p *SubtractPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.SubtractEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// Divide calls the "divide" endpoint of the "calc" service.
func (c *Client) Divide(ctx context.Context, p *DividePayload) (res float64, err error) {
	var ires interface{}
	ires, err = c.DivideEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(float64), nil
}

// GetNotes calls the "getNotes" endpoint of the "calc" service.
func (c *Client) GetNotes(ctx context.Context, p *GetNotesPayload) (res *GetNotesResult, err error) {
	var ires interface{}
	ires, err = c.GetNotesEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetNotesResult), nil
}

// CreateNote calls the "createNote" endpoint of the "calc" service.
func (c *Client) CreateNote(ctx context.Context, p *CreateNotePayload) (res *Note, err error) {
	var ires interface{}
	ires, err = c.CreateNoteEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Note), nil
}
