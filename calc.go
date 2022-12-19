package calcapi

import (
	"context"
	"goa-test/gen/calc"
	"log"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Multiply implements multiply.
func (s *calcsrvc) Multiply(_ context.Context, p *calc.MultiplyPayload) (res int, err error) {
	res = p.A * p.A
	return
}

// Add implements add.
func (s *calcsrvc) Add(_ context.Context, p *calc.AddPayload) (res int, err error) {
	res = p.A + p.B
	return
}

// Subtract implements subtract.
func (s *calcsrvc) Subtract(_ context.Context, p *calc.SubtractPayload) (res int, err error) {
	res = p.A - p.B
	return
}

// Divide implements divide.
func (s *calcsrvc) Divide(_ context.Context, p *calc.DividePayload) (res float64, err error) {
	res = float64(p.A) / float64(p.B)
	return
}
