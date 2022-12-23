// Code generated by goa v3.5.2, DO NOT EDIT.
//
// calc gRPC client encoders and decoders
//
// Command:
// $ goa gen goa-test/design

package client

import (
	"context"
	calc "goa-test/gen/calc"
	calcpb "goa-test/gen/grpc/calc/pb"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildMultiplyFunc builds the remote method to invoke for "calc" service
// "multiply" endpoint.
func BuildMultiplyFunc(grpccli calcpb.CalcClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Multiply(ctx, reqpb.(*calcpb.MultiplyRequest), opts...)
		}
		return grpccli.Multiply(ctx, &calcpb.MultiplyRequest{}, opts...)
	}
}

// EncodeMultiplyRequest encodes requests sent to calc multiply endpoint.
func EncodeMultiplyRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*calc.MultiplyPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "multiply", "*calc.MultiplyPayload", v)
	}
	return NewMultiplyRequest(payload), nil
}

// DecodeMultiplyResponse decodes responses from the calc multiply endpoint.
func DecodeMultiplyResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*calcpb.MultiplyResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "multiply", "*calcpb.MultiplyResponse", v)
	}
	res := NewMultiplyResult(message)
	return res, nil
}

// BuildAddFunc builds the remote method to invoke for "calc" service "add"
// endpoint.
func BuildAddFunc(grpccli calcpb.CalcClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Add(ctx, reqpb.(*calcpb.AddRequest), opts...)
		}
		return grpccli.Add(ctx, &calcpb.AddRequest{}, opts...)
	}
}

// EncodeAddRequest encodes requests sent to calc add endpoint.
func EncodeAddRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*calc.AddPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "add", "*calc.AddPayload", v)
	}
	return NewAddRequest(payload), nil
}

// DecodeAddResponse decodes responses from the calc add endpoint.
func DecodeAddResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*calcpb.AddResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "add", "*calcpb.AddResponse", v)
	}
	res := NewAddResult(message)
	return res, nil
}

// BuildSubtractFunc builds the remote method to invoke for "calc" service
// "subtract" endpoint.
func BuildSubtractFunc(grpccli calcpb.CalcClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Subtract(ctx, reqpb.(*calcpb.SubtractRequest), opts...)
		}
		return grpccli.Subtract(ctx, &calcpb.SubtractRequest{}, opts...)
	}
}

// EncodeSubtractRequest encodes requests sent to calc subtract endpoint.
func EncodeSubtractRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*calc.SubtractPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "subtract", "*calc.SubtractPayload", v)
	}
	return NewSubtractRequest(payload), nil
}

// DecodeSubtractResponse decodes responses from the calc subtract endpoint.
func DecodeSubtractResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*calcpb.SubtractResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "subtract", "*calcpb.SubtractResponse", v)
	}
	res := NewSubtractResult(message)
	return res, nil
}

// BuildDivideFunc builds the remote method to invoke for "calc" service
// "divide" endpoint.
func BuildDivideFunc(grpccli calcpb.CalcClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Divide(ctx, reqpb.(*calcpb.DivideRequest), opts...)
		}
		return grpccli.Divide(ctx, &calcpb.DivideRequest{}, opts...)
	}
}

// EncodeDivideRequest encodes requests sent to calc divide endpoint.
func EncodeDivideRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*calc.DividePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "divide", "*calc.DividePayload", v)
	}
	return NewDivideRequest(payload), nil
}

// DecodeDivideResponse decodes responses from the calc divide endpoint.
func DecodeDivideResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*calcpb.DivideResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "divide", "*calcpb.DivideResponse", v)
	}
	res := NewDivideResult(message)
	return res, nil
}

// BuildGetNotesFunc builds the remote method to invoke for "calc" service
// "getNotes" endpoint.
func BuildGetNotesFunc(grpccli calcpb.CalcClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.GetNotes(ctx, reqpb.(*calcpb.GetNotesRequest), opts...)
		}
		return grpccli.GetNotes(ctx, &calcpb.GetNotesRequest{}, opts...)
	}
}

// EncodeGetNotesRequest encodes requests sent to calc getNotes endpoint.
func EncodeGetNotesRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*calc.GetNotesPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "getNotes", "*calc.GetNotesPayload", v)
	}
	return NewGetNotesRequest(payload), nil
}

// DecodeGetNotesResponse decodes responses from the calc getNotes endpoint.
func DecodeGetNotesResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*calcpb.GetNotesResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "getNotes", "*calcpb.GetNotesResponse", v)
	}
	res := NewGetNotesResult(message)
	return res, nil
}

// BuildGetNoteFunc builds the remote method to invoke for "calc" service
// "getNote" endpoint.
func BuildGetNoteFunc(grpccli calcpb.CalcClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.GetNote(ctx, reqpb.(*calcpb.GetNoteRequest), opts...)
		}
		return grpccli.GetNote(ctx, &calcpb.GetNoteRequest{}, opts...)
	}
}

// EncodeGetNoteRequest encodes requests sent to calc getNote endpoint.
func EncodeGetNoteRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*calc.GetNotePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "getNote", "*calc.GetNotePayload", v)
	}
	return NewGetNoteRequest(payload), nil
}

// DecodeGetNoteResponse decodes responses from the calc getNote endpoint.
func DecodeGetNoteResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*calcpb.GetNoteResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "getNote", "*calcpb.GetNoteResponse", v)
	}
	res := NewGetNoteResult(message)
	return res, nil
}

// BuildCreateNoteFunc builds the remote method to invoke for "calc" service
// "createNote" endpoint.
func BuildCreateNoteFunc(grpccli calcpb.CalcClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.CreateNote(ctx, reqpb.(*calcpb.CreateNoteRequest), opts...)
		}
		return grpccli.CreateNote(ctx, &calcpb.CreateNoteRequest{}, opts...)
	}
}

// EncodeCreateNoteRequest encodes requests sent to calc createNote endpoint.
func EncodeCreateNoteRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*calc.CreateNotePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "createNote", "*calc.CreateNotePayload", v)
	}
	return NewCreateNoteRequest(payload), nil
}

// DecodeCreateNoteResponse decodes responses from the calc createNote endpoint.
func DecodeCreateNoteResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*calcpb.CreateNoteResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "createNote", "*calcpb.CreateNoteResponse", v)
	}
	if err := ValidateCreateNoteResponse(message); err != nil {
		return nil, err
	}
	res := NewCreateNoteResult(message)
	return res, nil
}
