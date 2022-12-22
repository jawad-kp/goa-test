// Code generated by goa v3.10.2, DO NOT EDIT.
//
// calc gRPC server encoders and decoders
//
// Command:
// $ goa gen goa-test/design

package server

import (
	"context"
	calc "goa-test/gen/calc"
	calcpb "goa-test/gen/grpc/calc/pb"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeMultiplyResponse encodes responses from the "calc" service "multiply"
// endpoint.
func EncodeMultiplyResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(int)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "multiply", "int", v)
	}
	resp := NewProtoMultiplyResponse(result)
	return resp, nil
}

// DecodeMultiplyRequest decodes requests sent to "calc" service "multiply"
// endpoint.
func DecodeMultiplyRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *calcpb.MultiplyRequest
		ok      bool
	)
	{
		if message, ok = v.(*calcpb.MultiplyRequest); !ok {
			return nil, goagrpc.ErrInvalidType("calc", "multiply", "*calcpb.MultiplyRequest", v)
		}
	}
	var payload *calc.MultiplyPayload
	{
		payload = NewMultiplyPayload(message)
	}
	return payload, nil
}

// EncodeAddResponse encodes responses from the "calc" service "add" endpoint.
func EncodeAddResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(int)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "add", "int", v)
	}
	resp := NewProtoAddResponse(result)
	return resp, nil
}

// DecodeAddRequest decodes requests sent to "calc" service "add" endpoint.
func DecodeAddRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *calcpb.AddRequest
		ok      bool
	)
	{
		if message, ok = v.(*calcpb.AddRequest); !ok {
			return nil, goagrpc.ErrInvalidType("calc", "add", "*calcpb.AddRequest", v)
		}
	}
	var payload *calc.AddPayload
	{
		payload = NewAddPayload(message)
	}
	return payload, nil
}

// EncodeSubtractResponse encodes responses from the "calc" service "subtract"
// endpoint.
func EncodeSubtractResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(int)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "subtract", "int", v)
	}
	resp := NewProtoSubtractResponse(result)
	return resp, nil
}

// DecodeSubtractRequest decodes requests sent to "calc" service "subtract"
// endpoint.
func DecodeSubtractRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *calcpb.SubtractRequest
		ok      bool
	)
	{
		if message, ok = v.(*calcpb.SubtractRequest); !ok {
			return nil, goagrpc.ErrInvalidType("calc", "subtract", "*calcpb.SubtractRequest", v)
		}
	}
	var payload *calc.SubtractPayload
	{
		payload = NewSubtractPayload(message)
	}
	return payload, nil
}

// EncodeDivideResponse encodes responses from the "calc" service "divide"
// endpoint.
func EncodeDivideResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(float64)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "divide", "float64", v)
	}
	resp := NewProtoDivideResponse(result)
	return resp, nil
}

// DecodeDivideRequest decodes requests sent to "calc" service "divide"
// endpoint.
func DecodeDivideRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *calcpb.DivideRequest
		ok      bool
	)
	{
		if message, ok = v.(*calcpb.DivideRequest); !ok {
			return nil, goagrpc.ErrInvalidType("calc", "divide", "*calcpb.DivideRequest", v)
		}
	}
	var payload *calc.DividePayload
	{
		payload = NewDividePayload(message)
	}
	return payload, nil
}

// EncodeGetNotesResponse encodes responses from the "calc" service "getNotes"
// endpoint.
func EncodeGetNotesResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*calc.GetNotesResult)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "getNotes", "*calc.GetNotesResult", v)
	}
	resp := NewProtoGetNotesResponse(result)
	return resp, nil
}

// DecodeGetNotesRequest decodes requests sent to "calc" service "getNotes"
// endpoint.
func DecodeGetNotesRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *calcpb.GetNotesRequest
		ok      bool
	)
	{
		if message, ok = v.(*calcpb.GetNotesRequest); !ok {
			return nil, goagrpc.ErrInvalidType("calc", "getNotes", "*calcpb.GetNotesRequest", v)
		}
	}
	var payload *calc.GetNotesPayload
	{
		payload = NewGetNotesPayload(message)
	}
	return payload, nil
}

// EncodeCreateNoteResponse encodes responses from the "calc" service
// "createNote" endpoint.
func EncodeCreateNoteResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*calc.Note)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "createNote", "*calc.Note", v)
	}
	resp := NewProtoCreateNoteResponse(result)
	return resp, nil
}

// DecodeCreateNoteRequest decodes requests sent to "calc" service "createNote"
// endpoint.
func DecodeCreateNoteRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *calcpb.CreateNoteRequest
		ok      bool
	)
	{
		if message, ok = v.(*calcpb.CreateNoteRequest); !ok {
			return nil, goagrpc.ErrInvalidType("calc", "createNote", "*calcpb.CreateNoteRequest", v)
		}
	}
	var payload *calc.CreateNotePayload
	{
		payload = NewCreateNotePayload(message)
	}
	return payload, nil
}
