// Code generated by goa v3.5.2, DO NOT EDIT.
//
// calc HTTP client encoders and decoders
//
// Command:
// $ goa gen goa-test/design

package client

import (
	"bytes"
	"context"
	calc "goa-test/gen/calc"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildMultiplyRequest instantiates a HTTP request object with method and path
// set to call the "calc" service "multiply" endpoint
func (c *Client) BuildMultiplyRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		a int
		b int
	)
	{
		p, ok := v.(*calc.MultiplyPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("calc", "multiply", "*calc.MultiplyPayload", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: MultiplyCalcPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("calc", "multiply", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeMultiplyResponse returns a decoder for responses returned by the calc
// multiply endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeMultiplyResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body int
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "multiply", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("calc", "multiply", resp.StatusCode, string(body))
		}
	}
}

// BuildAddRequest instantiates a HTTP request object with method and path set
// to call the "calc" service "add" endpoint
func (c *Client) BuildAddRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		a int
		b int
	)
	{
		p, ok := v.(*calc.AddPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("calc", "add", "*calc.AddPayload", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddCalcPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("calc", "add", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeAddResponse returns a decoder for responses returned by the calc add
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeAddResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body int
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "add", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("calc", "add", resp.StatusCode, string(body))
		}
	}
}

// BuildSubtractRequest instantiates a HTTP request object with method and path
// set to call the "calc" service "subtract" endpoint
func (c *Client) BuildSubtractRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		a int
		b int
	)
	{
		p, ok := v.(*calc.SubtractPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("calc", "subtract", "*calc.SubtractPayload", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SubtractCalcPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("calc", "subtract", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeSubtractResponse returns a decoder for responses returned by the calc
// subtract endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeSubtractResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body int
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "subtract", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("calc", "subtract", resp.StatusCode, string(body))
		}
	}
}

// BuildDivideRequest instantiates a HTTP request object with method and path
// set to call the "calc" service "divide" endpoint
func (c *Client) BuildDivideRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		a int
		b int
	)
	{
		p, ok := v.(*calc.DividePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("calc", "divide", "*calc.DividePayload", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DivideCalcPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("calc", "divide", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDivideResponse returns a decoder for responses returned by the calc
// divide endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeDivideResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body float64
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "divide", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("calc", "divide", resp.StatusCode, string(body))
		}
	}
}

// BuildGetNotesRequest instantiates a HTTP request object with method and path
// set to call the "calc" service "getNotes" endpoint
func (c *Client) BuildGetNotesRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		userID string
	)
	{
		p, ok := v.(*calc.GetNotesPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("calc", "getNotes", "*calc.GetNotesPayload", v)
		}
		userID = p.UserID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetNotesCalcPath(userID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("calc", "getNotes", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetNotesResponse returns a decoder for responses returned by the calc
// getNotes endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetNotesResponse may return the following errors:
//   - "NoteMissing" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeGetNotesResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetNotesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "getNotes", err)
			}
			res := NewGetNotesResultOK(&body)
			return res, nil
		case http.StatusNotFound:
			var (
				body GetNotesNoteMissingResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "getNotes", err)
			}
			err = ValidateGetNotesNoteMissingResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("calc", "getNotes", err)
			}
			return nil, NewGetNotesNoteMissing(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("calc", "getNotes", resp.StatusCode, string(body))
		}
	}
}

// BuildGetNoteRequest instantiates a HTTP request object with method and path
// set to call the "calc" service "getNote" endpoint
func (c *Client) BuildGetNoteRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		uuid string
	)
	{
		p, ok := v.(*calc.GetNotePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("calc", "getNote", "*calc.GetNotePayload", v)
		}
		uuid = p.UUID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetNoteCalcPath(uuid)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("calc", "getNote", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetNoteResponse returns a decoder for responses returned by the calc
// getNote endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetNoteResponse may return the following errors:
//   - "NoteMissing" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeGetNoteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetNoteResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "getNote", err)
			}
			res := NewGetNoteResultOK(&body)
			return res, nil
		case http.StatusNotFound:
			var (
				body GetNoteNoteMissingResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "getNote", err)
			}
			err = ValidateGetNoteNoteMissingResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("calc", "getNote", err)
			}
			return nil, NewGetNoteNoteMissing(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("calc", "getNote", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateNoteRequest instantiates a HTTP request object with method and
// path set to call the "calc" service "createNote" endpoint
func (c *Client) BuildCreateNoteRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		userID string
	)
	{
		p, ok := v.(*calc.CreateNotePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("calc", "createNote", "*calc.CreateNotePayload", v)
		}
		userID = p.UserID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateNoteCalcPath(userID)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("calc", "createNote", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateNoteRequest returns an encoder for requests sent to the calc
// createNote server.
func EncodeCreateNoteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*calc.CreateNotePayload)
		if !ok {
			return goahttp.ErrInvalidType("calc", "createNote", "*calc.CreateNotePayload", v)
		}
		body := NewCreateNoteRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("calc", "createNote", err)
		}
		return nil
	}
}

// DecodeCreateNoteResponse returns a decoder for responses returned by the
// calc createNote endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeCreateNoteResponse may return the following errors:
//   - "BadRequest" (type *goa.ServiceError): http.StatusBadRequest
//   - error: internal error
func DecodeCreateNoteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateNoteResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "createNote", err)
			}
			err = ValidateCreateNoteResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("calc", "createNote", err)
			}
			res := NewCreateNoteResultCreated(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body CreateNoteBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "createNote", err)
			}
			err = ValidateCreateNoteBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("calc", "createNote", err)
			}
			return nil, NewCreateNoteBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("calc", "createNote", resp.StatusCode, string(body))
		}
	}
}

// unmarshalNoteResponseBodyToCalcNote builds a value of type *calc.Note from a
// value of type *NoteResponseBody.
func unmarshalNoteResponseBodyToCalcNote(v *NoteResponseBody) *calc.Note {
	if v == nil {
		return nil
	}
	res := &calc.Note{
		Title: v.Title,
		Body:  v.Body,
		UUID:  v.UUID,
	}

	return res
}

// unmarshalNoteResponseResponseBodyToCalcNoteResponse builds a value of type
// *calc.NoteResponse from a value of type *NoteResponseResponseBody.
func unmarshalNoteResponseResponseBodyToCalcNoteResponse(v *NoteResponseResponseBody) *calc.NoteResponse {
	res := &calc.NoteResponse{
		UUID: v.UUID,
	}

	return res
}
