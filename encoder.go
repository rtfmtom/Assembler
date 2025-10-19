package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func (c *CInstruction) Encode(instructions []string, token []byte) ([]string, error) {
	token = bytes.Trim(token, "\r\n")

	if err := c.Parse(token); err != nil {
		return []string{}, err
	}
	c.Encoded = append(c.Encoded, "111")
	if c.Comp == nil {
		return []string{}, ErrComputationNil
	}

	b, err := c.LookupComp()
	if err != nil {
		return []string{}, err
	}
	c.Encoded = append(c.Encoded, b)

	if c.Dest != nil {
		b, err = c.LookupDest()
		if err != nil {
			return []string{}, err
		}
		c.Encoded = append(c.Encoded, b)
	} else {
		c.Encoded = append(c.Encoded, destBinary[NullDest])
	}

	if c.Jump != nil {
		b, err = c.LookupJump()
		if err != nil {
			return []string{}, err
		}
		c.Encoded = append(c.Encoded, b)
	} else {
		c.Encoded = append(c.Encoded, jumpBinary[NullJump])
	}

	v, err := strconv.ParseUint(strings.Join(c.Encoded, ""), 2, 16)
	if err != nil {
		return []string{}, fmt.Errorf("error: %w", err)
	}

	var encoded string
	switch outputFormat {
	case "hex", "x":
		encoded = fmt.Sprintf("%04x", v)
	case "bin", "b":
		encoded = fmt.Sprintf("%016b", v)
	default:
		encoded = fmt.Sprintf("%04x", v)
	}

	return append(instructions, encoded), nil
}

func (a *AInstruction) Encode(instructions []string, token []byte, symbolTable *SymbolTable) ([]string, error) {
	token = bytes.Trim(token, "\r\n")

	symbolOrValue := string(token[1:])

	var v int

	if num, err := strconv.Atoi(symbolOrValue); err == nil {
		v = num
	} else {
		if !isValidSymbol(symbolOrValue) {
			return nil, fmt.Errorf("%w: %s", ErrInvalidSymbol, symbolOrValue)
		}

		if addr, ok := symbolTable.Address(symbolOrValue); ok {
			v = addr
		} else {
			v = symbolTable.AddSymbol(symbolOrValue)
		}
	}

	if v < 0 || v >= (1<<15) {
		return []string{}, &ErrIntegerOverflow{
			Register: "A",
			BitWidth: 15,
			Signed:   false,
		}
	}

	var encoded string
	switch outputFormat {
	case "hex", "x":
		encoded = fmt.Sprintf("%04x", v)
	case "bin", "b":
		encoded = fmt.Sprintf("0%015b", v)
	default:
		encoded = fmt.Sprintf("%04x", v)
	}

	return append(instructions, encoded), nil
}

func (c *CInstruction) LookupDest() (string, error) {
	destStr := string(c.Dest)
	switch destStr {
	case "A":
		c.Dest = []byte("ADest")
	case "D":
		c.Dest = []byte("DDest")
	case "M":
		c.Dest = []byte("MDest")
	}

	if pattern, ok := destinationCodes[string(c.Dest)]; ok {
		return destBinary[pattern], nil
	}
	return "", fmt.Errorf("%w, got: %q", ErrMalformedDestination, string(c.Dest))
}

func (c *CInstruction) LookupComp() (string, error) {
	if pattern, ok := computationCodes[string(c.Comp)]; ok {
		return compBinary[pattern], nil
	}
	return "", fmt.Errorf("%w, got: %q", ErrMalformedComputation, string(c.Comp))
}

func (c *CInstruction) LookupJump() (string, error) {
	if pattern, ok := jumpCodes[string(c.Jump)]; ok {
		return jumpBinary[pattern], nil
	}
	return "", fmt.Errorf("%w, got: %q", ErrMalformedJump, string(c.Jump))
}
