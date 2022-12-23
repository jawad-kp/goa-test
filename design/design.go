package design

import (
	. "goa.design/goa/v3/dsl"
)

var (
	Note = Type("Note", func() {
		Field(1, "Title", String, "The title of the Note")
		Field(2, "Body", String, "The Body of the Note")

	})
	NoteResponse = Type("NoteResponse", func() {
		Field(1, "UUID", String, "The UUID of the Created Note")
	})
)

var _ = API("calc", func() {
	Title("Calculator Service")
	Description("Service for multiplying numbers, a Goa teaser")
	Server("calc", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers.")

	Method("multiply", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(Int)

		HTTP(func() {
			GET("/multiply/{a}/{b}")
		})

		GRPC(func() {
		})
	})
	Method("add", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(Int)

		HTTP(func() {
			GET("/add/{a}/{b}")
		})

		GRPC(func() {
		})
	})
	Method("subtract", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(Int)

		HTTP(func() {
			GET("/subtract/{a}/{b}")
		})

		GRPC(func() {
		})
	})
	Method("divide", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})

		Result(Float64)

		HTTP(func() {
			GET("/divide/{a}/{b}")
		})

		GRPC(func() {
		})
	})

	Method("getNotes", func() {
		Payload(func() {
			Field(1, "userID", String, "The email of the user")
			Required("userID")
		})

		Result(func() {
			Field(1, "notes", ArrayOf(Note), "list of notes")
		})

		HTTP(func() {
			GET("/notes/{userID}")
		})

		GRPC(func() {
		})
	})
	Method("createNote", func() {
		Payload(func() {
			Field(1, "userID", String, "The UserID for the note")
			Field(2, "Note", Note, "The Note data to be saved")
			Required("userID")
		})
		Result(func() {
			Field(1, "NoteResponse", NoteResponse, "The Created Note")
			Required("NoteResponse")
		})
		Error("BadRequest")
		HTTP(func() {
			POST("/notes/{userID}")
			Body("Note")
			Response(StatusCreated)
			Response("BadRequest", StatusBadRequest)

		})
		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
