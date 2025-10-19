package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func cleanLine(line []byte) []byte {
	line = bytes.TrimSpace(line)
	if index := bytes.Index(line, []byte("//")); index != -1 {
		line = line[:index]
		line = bytes.TrimSpace(line)
	}
	return line
}

func isLabel(line []byte) bool {
	return len(line) > 0 && line[0] == '(' && line[len(line)-1] == ')'
}

func extractLabel(line []byte) string {
	return string(line[1 : len(line)-1])
}

func isValidSymbol(symbol string) bool {
	if len(symbol) == 0 {
		return false
	}

	first := rune(symbol[0])
	if !unicode.IsLetter(first) && first != '_' && first != '.' && first != '$' && first != ':' {
		return false
	}

	for _, ch := range symbol[1:] {
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) && ch != '_' && ch != '.' && ch != '$' && ch != ':' && ch != '-' {
			return false
		}
	}

	return true
}

func buildSymbolTable(lines [][]byte) (*SymbolTable, error) {
	symbolTable := NewSymbolTable()

	romAddress := 0
	for _, line := range lines {
		line = cleanLine(line)
		if len(line) == 0 {
			continue
		}

		if isLabel(line) {
			label := extractLabel(line)
			if !isValidSymbol(label) {
				return nil, fmt.Errorf("%w: %s", ErrInvalidSymbol, label)
			}
			symbolTable.AddEntry(label, romAddress)
		} else {
			romAddress++
		}
	}
	return symbolTable, nil
}

func (c *CInstruction) Parse(token []byte) error {
	parts := bytes.Split(token, []byte(";"))
	if len(parts) > 2 {
		return fmt.Errorf("%w: %q (expected format: [Dest=]Comp[;Jump])",
			ErrTooManySemicolons, string(token))
	}

	if len(parts) == 2 {
		c.Jump = parts[1]
	}

	condition := bytes.Split(parts[0], []byte("="))

	if len(condition) > 2 {
		return fmt.Errorf("%w: %q (expected format: [Dest=]Comp[;Jump])",
			ErrTooManyEquals, string(token))
	}

	if len(condition) < 1 {
		return fmt.Errorf("%w: %q", ErrEmptyInstruction, string(token))
	}

	if len(condition) == 1 {
		c.Comp = condition[0]
	}

	if len(condition) == 2 {
		c.Dest = condition[0]
		c.Comp = condition[1]
	}

	return nil
}
