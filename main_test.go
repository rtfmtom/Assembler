package main

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

func TestAddProgram(t *testing.T) {
	inputFile := "testdata/Add.asm"
	outputFormat = "bin"

	data, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	lines := bytes.Split(data, []byte("\n"))
	instructions, err := assemble(lines)

	expected := []string{
		"0000000000000010",
		"1110110000010000",
		"0000000000000011",
		"1110000010010000",
		"0000000000000000",
		"1110001100001000",
	}

	if len(instructions) != len(expected) {
		t.Fatalf("Expected %d instructions, got %d", len(expected), len(instructions))
	}

	for i, instr := range instructions {
		if instr != expected[i] {
			t.Errorf("Instruction %d: expected %s, got %s", i, expected[i], instr)
		}
	}
}

func TestMaxLProgram(t *testing.T) {

	inputFile := "testdata/MaxL.asm"
	outputFormat = "bin"

	data, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	lines := bytes.Split(data, []byte("\n"))
	instructions, err := assemble(lines)

	expected := []string{
		"0000000000000000",
		"1111110000010000",
		"0000000000000001",
		"1111010011010000",
		"0000000000001010",
		"1110001100000001",
		"0000000000000001",
		"1111110000010000",
		"0000000000001100",
		"1110101010000111",
		"0000000000000000",
		"1111110000010000",
		"0000000000000010",
		"1110001100001000",
		"0000000000001110",
		"1110101010000111",
	}

	if len(instructions) != len(expected) {
		t.Fatalf("Expected %d instructions, got %d", len(expected), len(instructions))
	}

	for i, instr := range instructions {
		if instr != expected[i] {
			t.Errorf("Instruction %d: expected %s, got %s", i, expected[i], instr)
		}
	}
}

func TestAddProgramHex(t *testing.T) {
	inputFile := "testdata/Add.asm"
	outputFormat = "hex"

	data, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	lines := bytes.Split(data, []byte("\n"))
	instructions, err := assemble(lines)

	expected := []string{
		"0002",
		"ec10",
		"0003",
		"e090",
		"0000",
		"e308",
	}

	if len(instructions) != len(expected) {
		t.Fatalf("Expected %d instructions, got %d", len(expected), len(instructions))
	}

	for i, instr := range instructions {
		if instr != expected[i] {
			t.Errorf("Instruction %d: expected %s, got %s", i, expected[i], instr)
		}
	}
}

func TestMaxLProgramHex(t *testing.T) {
	inputFile := "testdata/MaxL.asm"
	outputFormat = "hex"

	data, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	lines := bytes.Split(data, []byte("\n"))
	instructions, err := assemble(lines)

	expected := []string{
		"0000",
		"fc10",
		"0001",
		"f4d0",
		"000a",
		"e301",
		"0001",
		"fc10",
		"000c",
		"ea87",
		"0000",
		"fc10",
		"0002",
		"e308",
		"000e",
		"ea87",
	}

	if len(instructions) != len(expected) {
		t.Fatalf("Expected %d instructions, got %d", len(expected), len(instructions))
	}

	for i, instr := range instructions {
		if instr != expected[i] {
			t.Errorf("Instruction %d: expected %s, got %s", i, expected[i], instr)
		}
	}
}

func TestMaxProgramHex(t *testing.T) {
	inputFile := "testdata/Max.asm"
	outputFormat = "hex"

	data, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	lines := bytes.Split(data, []byte("\n"))
	instructions, err := assemble(lines)

	expected := []string{
		"0000",
		"fc10",
		"0001",
		"f4d0",
		"000a",
		"e301",
		"0001",
		"fc10",
		"000c",
		"ea87",
		"0000",
		"fc10",
		"0002",
		"e308",
		"000e",
		"ea87",
	}

	if len(instructions) != len(expected) {
		t.Fatalf("Expected %d instructions, got %d", len(expected), len(instructions))
	}

	for i, instr := range instructions {
		if instr != expected[i] {
			t.Errorf("Instruction %d: expected %s, got %s", i, expected[i], instr)
		}
	}
}

func TestAInstructionIntegerOverflow(t *testing.T) {
	outputFormat = "bin"

	data := []byte("@32769")
	lines := bytes.Split(data, []byte("\n"))

	_, err := assemble(lines)

	if err == nil {
		t.Fatal("Expected integer overflow error, got nil")
	}

	var overflowErr *ErrIntegerOverflow
	if !errors.As(err, &overflowErr) {
		t.Fatalf("Expected *ErrIntegerOverflow, got %T: %v", err, err)
	}

	if overflowErr.Register != "A" {
		t.Errorf("Expected Register 'A', got '%s'", overflowErr.Register)
	}
	if overflowErr.BitWidth != 15 {
		t.Errorf("Expected BitWidth 15, got %d", overflowErr.BitWidth)
	}
	if overflowErr.Signed != false {
		t.Errorf("Expected Signed false, got %t", overflowErr.Signed)
	}

	expectedMsg := "integer overflow - Register A (15-bit unsigned) valid range: 0 to 32767"
	if overflowErr.Error() != expectedMsg {
		t.Errorf("Expected error message %q, got %q", expectedMsg, overflowErr.Error())
	}
}
