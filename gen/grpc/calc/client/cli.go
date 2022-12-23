// Code generated by goa v3.5.2, DO NOT EDIT.
//
// calc gRPC client CLI support package
//
// Command:
// $ goa gen goa-test/design

package client

import (
	"encoding/json"
	"fmt"
	calc "goa-test/gen/calc"
	calcpb "goa-test/gen/grpc/calc/pb"
)

// BuildMultiplyPayload builds the payload for the calc multiply endpoint from
// CLI flags.
func BuildMultiplyPayload(calcMultiplyMessage string) (*calc.MultiplyPayload, error) {
	var err error
	var message calcpb.MultiplyRequest
	{
		if calcMultiplyMessage != "" {
			err = json.Unmarshal([]byte(calcMultiplyMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"a\": 5752212304782361938,\n      \"b\": 5269901698896315977\n   }'")
			}
		}
	}
	v := &calc.MultiplyPayload{
		A: int(message.A),
		B: int(message.B),
	}

	return v, nil
}

// BuildAddPayload builds the payload for the calc add endpoint from CLI flags.
func BuildAddPayload(calcAddMessage string) (*calc.AddPayload, error) {
	var err error
	var message calcpb.AddRequest
	{
		if calcAddMessage != "" {
			err = json.Unmarshal([]byte(calcAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"a\": 7047189229146369891,\n      \"b\": 3896950918919820613\n   }'")
			}
		}
	}
	v := &calc.AddPayload{
		A: int(message.A),
		B: int(message.B),
	}

	return v, nil
}

// BuildSubtractPayload builds the payload for the calc subtract endpoint from
// CLI flags.
func BuildSubtractPayload(calcSubtractMessage string) (*calc.SubtractPayload, error) {
	var err error
	var message calcpb.SubtractRequest
	{
		if calcSubtractMessage != "" {
			err = json.Unmarshal([]byte(calcSubtractMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"a\": 460832770702694527,\n      \"b\": 3464015629245697118\n   }'")
			}
		}
	}
	v := &calc.SubtractPayload{
		A: int(message.A),
		B: int(message.B),
	}

	return v, nil
}

// BuildDividePayload builds the payload for the calc divide endpoint from CLI
// flags.
func BuildDividePayload(calcDivideMessage string) (*calc.DividePayload, error) {
	var err error
	var message calcpb.DivideRequest
	{
		if calcDivideMessage != "" {
			err = json.Unmarshal([]byte(calcDivideMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"a\": 2105883920431331421,\n      \"b\": 9040457143694300258\n   }'")
			}
		}
	}
	v := &calc.DividePayload{
		A: int(message.A),
		B: int(message.B),
	}

	return v, nil
}

// BuildGetNotesPayload builds the payload for the calc getNotes endpoint from
// CLI flags.
func BuildGetNotesPayload(calcGetNotesMessage string) (*calc.GetNotesPayload, error) {
	var err error
	var message calcpb.GetNotesRequest
	{
		if calcGetNotesMessage != "" {
			err = json.Unmarshal([]byte(calcGetNotesMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"userID\": \"Nihil odit iusto vitae.\"\n   }'")
			}
		}
	}
	v := &calc.GetNotesPayload{
		UserID: message.UserId,
	}

	return v, nil
}

// BuildGetNotePayload builds the payload for the calc getNote endpoint from
// CLI flags.
func BuildGetNotePayload(calcGetNoteMessage string) (*calc.GetNotePayload, error) {
	var err error
	var message calcpb.GetNoteRequest
	{
		if calcGetNoteMessage != "" {
			err = json.Unmarshal([]byte(calcGetNoteMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"uuid\": \"Veritatis voluptas atque labore at fuga.\"\n   }'")
			}
		}
	}
	v := &calc.GetNotePayload{
		UUID: message.Uuid,
	}

	return v, nil
}

// BuildCreateNotePayload builds the payload for the calc createNote endpoint
// from CLI flags.
func BuildCreateNotePayload(calcCreateNoteMessage string) (*calc.CreateNotePayload, error) {
	var err error
	var message calcpb.CreateNoteRequest
	{
		if calcCreateNoteMessage != "" {
			err = json.Unmarshal([]byte(calcCreateNoteMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"NoteDetails\": {\n         \"Body\": \"Ut itaque sit corrupti velit.\",\n         \"Title\": \"Non et.\"\n      },\n      \"userID\": \"Aliquam vel facilis.\"\n   }'")
			}
		}
	}
	v := &calc.CreateNotePayload{
		UserID: message.UserId,
	}
	if message.NoteDetails != nil {
		v.NoteDetails = protobufCalcpbNoteDetailsToCalcNoteDetails(message.NoteDetails)
	}

	return v, nil
}
