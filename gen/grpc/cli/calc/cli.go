// Code generated by goa v3.10.2, DO NOT EDIT.
//
// calc gRPC client CLI support package
//
// Command:
// $ goa gen goa-test/design

package cli

import (
	"flag"
	"fmt"
	calcc "goa-test/gen/grpc/calc/client"
	"os"

	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `calc (multiply|add|subtract|divide|get-notes|create-note)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` calc multiply --message '{
      "a": 3546414149089559360,
      "b": 1618669811688599017
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		calcFlags = flag.NewFlagSet("calc", flag.ContinueOnError)

		calcMultiplyFlags       = flag.NewFlagSet("multiply", flag.ExitOnError)
		calcMultiplyMessageFlag = calcMultiplyFlags.String("message", "", "")

		calcAddFlags       = flag.NewFlagSet("add", flag.ExitOnError)
		calcAddMessageFlag = calcAddFlags.String("message", "", "")

		calcSubtractFlags       = flag.NewFlagSet("subtract", flag.ExitOnError)
		calcSubtractMessageFlag = calcSubtractFlags.String("message", "", "")

		calcDivideFlags       = flag.NewFlagSet("divide", flag.ExitOnError)
		calcDivideMessageFlag = calcDivideFlags.String("message", "", "")

		calcGetNotesFlags       = flag.NewFlagSet("get-notes", flag.ExitOnError)
		calcGetNotesMessageFlag = calcGetNotesFlags.String("message", "", "")

		calcCreateNoteFlags       = flag.NewFlagSet("create-note", flag.ExitOnError)
		calcCreateNoteMessageFlag = calcCreateNoteFlags.String("message", "", "")
	)
	calcFlags.Usage = calcUsage
	calcMultiplyFlags.Usage = calcMultiplyUsage
	calcAddFlags.Usage = calcAddUsage
	calcSubtractFlags.Usage = calcSubtractUsage
	calcDivideFlags.Usage = calcDivideUsage
	calcGetNotesFlags.Usage = calcGetNotesUsage
	calcCreateNoteFlags.Usage = calcCreateNoteUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "calc":
			svcf = calcFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "calc":
			switch epn {
			case "multiply":
				epf = calcMultiplyFlags

			case "add":
				epf = calcAddFlags

			case "subtract":
				epf = calcSubtractFlags

			case "divide":
				epf = calcDivideFlags

			case "get-notes":
				epf = calcGetNotesFlags

			case "create-note":
				epf = calcCreateNoteFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "calc":
			c := calcc.NewClient(cc, opts...)
			switch epn {
			case "multiply":
				endpoint = c.Multiply()
				data, err = calcc.BuildMultiplyPayload(*calcMultiplyMessageFlag)
			case "add":
				endpoint = c.Add()
				data, err = calcc.BuildAddPayload(*calcAddMessageFlag)
			case "subtract":
				endpoint = c.Subtract()
				data, err = calcc.BuildSubtractPayload(*calcSubtractMessageFlag)
			case "divide":
				endpoint = c.Divide()
				data, err = calcc.BuildDividePayload(*calcDivideMessageFlag)
			case "get-notes":
				endpoint = c.GetNotes()
				data, err = calcc.BuildGetNotesPayload(*calcGetNotesMessageFlag)
			case "create-note":
				endpoint = c.CreateNote()
				data, err = calcc.BuildCreateNotePayload(*calcCreateNoteMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// calcUsage displays the usage of the calc command and its subcommands.
func calcUsage() {
	fmt.Fprintf(os.Stderr, `The calc service performs operations on numbers.
Usage:
    %[1]s [globalflags] calc COMMAND [flags]

COMMAND:
    multiply: Multiply implements multiply.
    add: Add implements add.
    subtract: Subtract implements subtract.
    divide: Divide implements divide.
    get-notes: GetNotes implements getNotes.
    create-note: CreateNote implements createNote.

Additional help:
    %[1]s calc COMMAND --help
`, os.Args[0])
}
func calcMultiplyUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calc multiply -message JSON

Multiply implements multiply.
    -message JSON: 

Example:
    %[1]s calc multiply --message '{
      "a": 3546414149089559360,
      "b": 1618669811688599017
   }'
`, os.Args[0])
}

func calcAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calc add -message JSON

Add implements add.
    -message JSON: 

Example:
    %[1]s calc add --message '{
      "a": 4669454048700716669,
      "b": 1294861662385682529
   }'
`, os.Args[0])
}

func calcSubtractUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calc subtract -message JSON

Subtract implements subtract.
    -message JSON: 

Example:
    %[1]s calc subtract --message '{
      "a": 2574101197078288820,
      "b": 9032300625042805317
   }'
`, os.Args[0])
}

func calcDivideUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calc divide -message JSON

Divide implements divide.
    -message JSON: 

Example:
    %[1]s calc divide --message '{
      "a": 6480680201785989957,
      "b": 577828123511009763
   }'
`, os.Args[0])
}

func calcGetNotesUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calc get-notes -message JSON

GetNotes implements getNotes.
    -message JSON: 

Example:
    %[1]s calc get-notes --message '{
      "userID": "Voluptas sed neque consequatur maiores aut sed."
   }'
`, os.Args[0])
}

func calcCreateNoteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calc create-note -message JSON

CreateNote implements createNote.
    -message JSON: 

Example:
    %[1]s calc create-note --message '{
      "userID": "Quis dolorum corporis officiis rerum."
   }'
`, os.Args[0])
}
