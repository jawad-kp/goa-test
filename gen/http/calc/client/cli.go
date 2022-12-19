// Code generated by goa v3.10.2, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen goa-test/design

package client

import (
	"fmt"
	calc "goa-test/gen/calc"
	"strconv"
)

// BuildMultiplyPayload builds the payload for the calc multiply endpoint from
// CLI flags.
func BuildMultiplyPayload(calcMultiplyA string, calcMultiplyB string) (*calc.MultiplyPayload, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(calcMultiplyA, 10, strconv.IntSize)
		a = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for a, must be INT")
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(calcMultiplyB, 10, strconv.IntSize)
		b = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for b, must be INT")
		}
	}
	v := &calc.MultiplyPayload{}
	v.A = a
	v.B = b

	return v, nil
}

// BuildAddPayload builds the payload for the calc add endpoint from CLI flags.
func BuildAddPayload(calcAddA string, calcAddB string) (*calc.AddPayload, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(calcAddA, 10, strconv.IntSize)
		a = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for a, must be INT")
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(calcAddB, 10, strconv.IntSize)
		b = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for b, must be INT")
		}
	}
	v := &calc.AddPayload{}
	v.A = a
	v.B = b

	return v, nil
}

// BuildSubtractPayload builds the payload for the calc subtract endpoint from
// CLI flags.
func BuildSubtractPayload(calcSubtractA string, calcSubtractB string) (*calc.SubtractPayload, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(calcSubtractA, 10, strconv.IntSize)
		a = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for a, must be INT")
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(calcSubtractB, 10, strconv.IntSize)
		b = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for b, must be INT")
		}
	}
	v := &calc.SubtractPayload{}
	v.A = a
	v.B = b

	return v, nil
}

// BuildDividePayload builds the payload for the calc divide endpoint from CLI
// flags.
func BuildDividePayload(calcDivideA string, calcDivideB string) (*calc.DividePayload, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(calcDivideA, 10, strconv.IntSize)
		a = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for a, must be INT")
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(calcDivideB, 10, strconv.IntSize)
		b = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for b, must be INT")
		}
	}
	v := &calc.DividePayload{}
	v.A = a
	v.B = b

	return v, nil
}
