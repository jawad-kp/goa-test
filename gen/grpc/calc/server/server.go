// Code generated by goa v3.5.2, DO NOT EDIT.
//
// calc gRPC server
//
// Command:
// $ goa gen goa-test/design

package server

import (
	"context"
	calc "goa-test/gen/calc"
	calcpb "goa-test/gen/grpc/calc/pb"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the calcpb.CalcServer interface.
type Server struct {
	MultiplyH   goagrpc.UnaryHandler
	AddH        goagrpc.UnaryHandler
	SubtractH   goagrpc.UnaryHandler
	DivideH     goagrpc.UnaryHandler
	GetNotesH   goagrpc.UnaryHandler
	GetNoteH    goagrpc.UnaryHandler
	CreateNoteH goagrpc.UnaryHandler
	DeleteNoteH goagrpc.UnaryHandler
	calcpb.UnimplementedCalcServer
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the calc service endpoints.
func New(e *calc.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		MultiplyH:   NewMultiplyHandler(e.Multiply, uh),
		AddH:        NewAddHandler(e.Add, uh),
		SubtractH:   NewSubtractHandler(e.Subtract, uh),
		DivideH:     NewDivideHandler(e.Divide, uh),
		GetNotesH:   NewGetNotesHandler(e.GetNotes, uh),
		GetNoteH:    NewGetNoteHandler(e.GetNote, uh),
		CreateNoteH: NewCreateNoteHandler(e.CreateNote, uh),
		DeleteNoteH: NewDeleteNoteHandler(e.DeleteNote, uh),
	}
}

// NewMultiplyHandler creates a gRPC handler which serves the "calc" service
// "multiply" endpoint.
func NewMultiplyHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeMultiplyRequest, EncodeMultiplyResponse)
	}
	return h
}

// Multiply implements the "Multiply" method in calcpb.CalcServer interface.
func (s *Server) Multiply(ctx context.Context, message *calcpb.MultiplyRequest) (*calcpb.MultiplyResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "multiply")
	ctx = context.WithValue(ctx, goa.ServiceKey, "calc")
	resp, err := s.MultiplyH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*calcpb.MultiplyResponse), nil
}

// NewAddHandler creates a gRPC handler which serves the "calc" service "add"
// endpoint.
func NewAddHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeAddRequest, EncodeAddResponse)
	}
	return h
}

// Add implements the "Add" method in calcpb.CalcServer interface.
func (s *Server) Add(ctx context.Context, message *calcpb.AddRequest) (*calcpb.AddResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "add")
	ctx = context.WithValue(ctx, goa.ServiceKey, "calc")
	resp, err := s.AddH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*calcpb.AddResponse), nil
}

// NewSubtractHandler creates a gRPC handler which serves the "calc" service
// "subtract" endpoint.
func NewSubtractHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeSubtractRequest, EncodeSubtractResponse)
	}
	return h
}

// Subtract implements the "Subtract" method in calcpb.CalcServer interface.
func (s *Server) Subtract(ctx context.Context, message *calcpb.SubtractRequest) (*calcpb.SubtractResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "subtract")
	ctx = context.WithValue(ctx, goa.ServiceKey, "calc")
	resp, err := s.SubtractH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*calcpb.SubtractResponse), nil
}

// NewDivideHandler creates a gRPC handler which serves the "calc" service
// "divide" endpoint.
func NewDivideHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeDivideRequest, EncodeDivideResponse)
	}
	return h
}

// Divide implements the "Divide" method in calcpb.CalcServer interface.
func (s *Server) Divide(ctx context.Context, message *calcpb.DivideRequest) (*calcpb.DivideResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "divide")
	ctx = context.WithValue(ctx, goa.ServiceKey, "calc")
	resp, err := s.DivideH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*calcpb.DivideResponse), nil
}

// NewGetNotesHandler creates a gRPC handler which serves the "calc" service
// "getNotes" endpoint.
func NewGetNotesHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeGetNotesRequest, EncodeGetNotesResponse)
	}
	return h
}

// GetNotes implements the "GetNotes" method in calcpb.CalcServer interface.
func (s *Server) GetNotes(ctx context.Context, message *calcpb.GetNotesRequest) (*calcpb.GetNotesResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "getNotes")
	ctx = context.WithValue(ctx, goa.ServiceKey, "calc")
	resp, err := s.GetNotesH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*calcpb.GetNotesResponse), nil
}

// NewGetNoteHandler creates a gRPC handler which serves the "calc" service
// "getNote" endpoint.
func NewGetNoteHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeGetNoteRequest, EncodeGetNoteResponse)
	}
	return h
}

// GetNote implements the "GetNote" method in calcpb.CalcServer interface.
func (s *Server) GetNote(ctx context.Context, message *calcpb.GetNoteRequest) (*calcpb.GetNoteResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "getNote")
	ctx = context.WithValue(ctx, goa.ServiceKey, "calc")
	resp, err := s.GetNoteH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*calcpb.GetNoteResponse), nil
}

// NewCreateNoteHandler creates a gRPC handler which serves the "calc" service
// "createNote" endpoint.
func NewCreateNoteHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeCreateNoteRequest, EncodeCreateNoteResponse)
	}
	return h
}

// CreateNote implements the "CreateNote" method in calcpb.CalcServer interface.
func (s *Server) CreateNote(ctx context.Context, message *calcpb.CreateNoteRequest) (*calcpb.CreateNoteResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "createNote")
	ctx = context.WithValue(ctx, goa.ServiceKey, "calc")
	resp, err := s.CreateNoteH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*calcpb.CreateNoteResponse), nil
}

// NewDeleteNoteHandler creates a gRPC handler which serves the "calc" service
// "deleteNote" endpoint.
func NewDeleteNoteHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeDeleteNoteRequest, EncodeDeleteNoteResponse)
	}
	return h
}

// DeleteNote implements the "DeleteNote" method in calcpb.CalcServer interface.
func (s *Server) DeleteNote(ctx context.Context, message *calcpb.DeleteNoteRequest) (*calcpb.DeleteNoteResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "deleteNote")
	ctx = context.WithValue(ctx, goa.ServiceKey, "calc")
	resp, err := s.DeleteNoteH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*calcpb.DeleteNoteResponse), nil
}
