package main

import (
	"errors"
	"fmt"
)

type ErrIntegerOverflow struct {
	Register string
	BitWidth int
	Signed   bool
}

func (e *ErrIntegerOverflow) Error() string {
	var minBound, maxBound int64
	if e.Signed {
		maxBound = (1 << (e.BitWidth - 1)) - 1
		minBound = -(1 << (e.BitWidth - 1))
		return fmt.Sprintf("integer overflow - Register %s (%d-bit signed) valid range: %d to %d",
			e.Register, e.BitWidth, minBound, maxBound)
	}
	maxBound = (1 << e.BitWidth) - 1
	return fmt.Sprintf("integer overflow - Register %s (%d-bit unsigned) valid range: 0 to %d",
		e.Register, e.BitWidth, maxBound)
}

var (
	ErrComputationNil       = errors.New("error: computation should never be nil")
	ErrMalformedDestination = errors.New("error: malformed C-Instruction OpCode; expected Destination")
	ErrMalformedComputation = errors.New("error: malformed C-Instruction OpCode; expected Computation")
	ErrMalformedJump        = errors.New("error: malformed C-Instruction OpCode; expected Jump")
	ErrInvalidSymbol        = errors.New("error: invalid symbol name")
	ErrTooManySemicolons    = errors.New("error: parse: C-instruction contains multiple semicolons")
	ErrTooManyEquals        = errors.New("error: parse: C-instruction contains multiple equals signs")
	ErrEmptyInstruction     = errors.New("error: parse: C-instruction is empty after parsing")
)
