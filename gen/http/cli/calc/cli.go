// Code generated by goa v3.5.2, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen goa-test/design

package cli

import (
	"flag"
	"fmt"
	calcc "goa-test/gen/http/calc/client"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
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
	return os.Args[0] + ` calc multiply --a 8614350958614154271 --b 4285228716133201605` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		calcFlags = flag.NewFlagSet("calc", flag.ContinueOnError)

		calcMultiplyFlags = flag.NewFlagSet("multiply", flag.ExitOnError)
		calcMultiplyAFlag = calcMultiplyFlags.String("a", "REQUIRED", "Left operand")
		calcMultiplyBFlag = calcMultiplyFlags.String("b", "REQUIRED", "Right operand")

		calcAddFlags = flag.NewFlagSet("add", flag.ExitOnError)
		calcAddAFlag = calcAddFlags.String("a", "REQUIRED", "Left operand")
		calcAddBFlag = calcAddFlags.String("b", "REQUIRED", "Right operand")

		calcSubtractFlags = flag.NewFlagSet("subtract", flag.ExitOnError)
		calcSubtractAFlag = calcSubtractFlags.String("a", "REQUIRED", "Left operand")
		calcSubtractBFlag = calcSubtractFlags.String("b", "REQUIRED", "Right operand")

		calcDivideFlags = flag.NewFlagSet("divide", flag.ExitOnError)
		calcDivideAFlag = calcDivideFlags.String("a", "REQUIRED", "Left operand")
		calcDivideBFlag = calcDivideFlags.String("b", "REQUIRED", "Right operand")

		calcGetNotesFlags      = flag.NewFlagSet("get-notes", flag.ExitOnError)
		calcGetNotesUserIDFlag = calcGetNotesFlags.String("user-id", "REQUIRED", "The email of the user")

		calcCreateNoteFlags      = flag.NewFlagSet("create-note", flag.ExitOnError)
		calcCreateNoteBodyFlag   = calcCreateNoteFlags.String("body", "REQUIRED", "")
		calcCreateNoteUserIDFlag = calcCreateNoteFlags.String("user-id", "REQUIRED", "The UserID for the note")
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
			c := calcc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "multiply":
				endpoint = c.Multiply()
				data, err = calcc.BuildMultiplyPayload(*calcMultiplyAFlag, *calcMultiplyBFlag)
			case "add":
				endpoint = c.Add()
				data, err = calcc.BuildAddPayload(*calcAddAFlag, *calcAddBFlag)
			case "subtract":
				endpoint = c.Subtract()
				data, err = calcc.BuildSubtractPayload(*calcSubtractAFlag, *calcSubtractBFlag)
			case "divide":
				endpoint = c.Divide()
				data, err = calcc.BuildDividePayload(*calcDivideAFlag, *calcDivideBFlag)
			case "get-notes":
				endpoint = c.GetNotes()
				data, err = calcc.BuildGetNotesPayload(*calcGetNotesUserIDFlag)
			case "create-note":
				endpoint = c.CreateNote()
				data, err = calcc.BuildCreateNotePayload(*calcCreateNoteBodyFlag, *calcCreateNoteUserIDFlag)
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
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calc multiply -a INT -b INT

Multiply implements multiply.
    -a INT: Left operand
    -b INT: Right operand

Example:
    %[1]s calc multiply --a 8614350958614154271 --b 4285228716133201605
`, os.Args[0])
}

func calcAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calc add -a INT -b INT

Add implements add.
    -a INT: Left operand
    -b INT: Right operand

Example:
    %[1]s calc add --a 6595082841550375708 --b 3523480046783923250
`, os.Args[0])
}

func calcSubtractUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calc subtract -a INT -b INT

Subtract implements subtract.
    -a INT: Left operand
    -b INT: Right operand

Example:
    %[1]s calc subtract --a 4433681783854432708 --b 6518345420858817596
`, os.Args[0])
}

func calcDivideUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calc divide -a INT -b INT

Divide implements divide.
    -a INT: Left operand
    -b INT: Right operand

Example:
    %[1]s calc divide --a 6513733567661029654 --b 862594352938458553
`, os.Args[0])
}

func calcGetNotesUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calc get-notes -user-id STRING

GetNotes implements getNotes.
    -user-id STRING: The email of the user

Example:
    %[1]s calc get-notes --user-id "Laboriosam et consequatur tempore ex quae."
`, os.Args[0])
}

func calcCreateNoteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calc create-note -body JSON -user-id STRING

CreateNote implements createNote.
    -body JSON: 
    -user-id STRING: The UserID for the note

Example:
    %[1]s calc create-note --body '{
      "Body": "Accusantium culpa odit eaque.",
      "Title": "Veritatis id iure."
   }' --user-id "Non quod et facere."
`, os.Args[0])
}
