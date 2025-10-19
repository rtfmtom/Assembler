package main

import "fmt"

type CInstruction struct {
	Jump    []byte
	Comp    []byte
	Dest    []byte
	Encoded []string
}

type AInstruction struct {
	Prefix string
	Value  string
}

type SymbolTable struct {
	table       map[string]int
	nextAddress int
}

func NewSymbolTable() *SymbolTable {
	// nextAddress is initialized to 16 to reserve registers R0-R15
	st := &SymbolTable{
		table:       make(map[string]int),
		nextAddress: 16,
	}

	for i := 0; i < 16; i++ {
		st.table[fmt.Sprintf("R%d", i)] = i
	}

	st.table["SP"] = 0
	st.table["LCL"] = 1
	st.table["ARG"] = 2
	st.table["THIS"] = 3
	st.table["THAT"] = 4
	st.table["SCREEN"] = 16384
	st.table["KBD"] = 24576

	return st
}

func (st *SymbolTable) AddEntry(symbol string, address int) {
	st.table[symbol] = address
}

func (st *SymbolTable) Contains(symbol string) bool {
	_, ok := st.table[symbol]
	return ok
}

func (st *SymbolTable) Address(symbol string) (int, bool) {
	addr, ok := st.table[symbol]
	return addr, ok
}

func (st *SymbolTable) AddSymbol(symbol string) int {
	if !st.Contains(symbol) {
		st.table[symbol] = st.nextAddress
		st.nextAddress++
	}
	return st.table[symbol]
}
