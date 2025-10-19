package main

type ComputationCode int
type DestinationCode int
type JumpCode int

const (
	Zero ComputationCode = iota
	One
	NegativeOne
	D
	A
	M
	NotD
	NotA
	NotM
	NegativeD
	NegativeA
	NegativeM
	DPlusOne
	APlusOne
	MPlusOne
	DMinusOne
	AMinusOne
	MMinusOne
	DPlusA
	DPlusM
	DMinusA
	DMinusM
	AMinusD
	MMinusD
	DAndA
	DAndM
	DOrA
	DOrM
)

const (
	NullDest DestinationCode = iota
	MDest
	DDest
	DMDest
	ADest
	AMDest
	ADDest
	ADMDest
)

const (
	NullJump JumpCode = iota
	JGT
	JEQ
	JGE
	JLT
	JNE
	JLE
	JMP
)

var computationCodes = map[string]ComputationCode{
	"0":   Zero,
	"1":   One,
	"-1":  NegativeOne,
	"D":   D,
	"A":   A,
	"M":   M,
	"!D":  NotD,
	"!A":  NotA,
	"!M":  NotM,
	"-D":  NegativeD,
	"-A":  NegativeA,
	"-M":  NegativeM,
	"D+1": DPlusOne,
	"1+D": DPlusOne,
	"A+1": APlusOne,
	"1+A": APlusOne,
	"M+1": MPlusOne,
	"1+M": MPlusOne,
	"D-1": DMinusOne,
	"A-1": AMinusOne,
	"M-1": MMinusOne,
	"D+A": DPlusA,
	"A+D": DPlusA,
	"D+M": DPlusM,
	"M+D": DPlusM,
	"D-A": DMinusA,
	"D-M": DMinusM,
	"A-D": AMinusD,
	"M-D": MMinusD,
	"D&A": DAndA,
	"A&D": DAndA,
	"D&M": DAndM,
	"M&D": DAndM,
	"D|A": DOrA,
	"A|D": DOrA,
	"D|M": DOrM,
	"M|D": DOrM,
}

var destinationCodes = map[string]DestinationCode{
	"ADest": ADest,
	"DDest": DDest,
	"MDest": MDest,
	"DM":    DMDest,
	"MD":    DMDest,
	"AM":    AMDest,
	"MA":    AMDest,
	"AD":    ADDest,
	"DA":    ADDest,
	"ADM":   ADMDest,
	"AMD":   ADMDest,
	"DAM":   ADMDest,
	"DMA":   ADMDest,
	"MAD":   ADMDest,
	"MDA":   ADMDest,
}

var jumpCodes = map[string]JumpCode{
	"JGT": JGT,
	"JEQ": JEQ,
	"JGE": JGE,
	"JLT": JLT,
	"JNE": JNE,
	"JLE": JLE,
	"JMP": JMP,
}

var compBinary = map[ComputationCode]string{
	Zero:        "0101010",
	One:         "0111111",
	NegativeOne: "0111010",
	D:           "0001100",
	A:           "0110000",
	M:           "1110000",
	NotD:        "0001101",
	NotA:        "0110001",
	NotM:        "1110001",
	NegativeD:   "0001111",
	NegativeA:   "0110011",
	NegativeM:   "1110011",
	DPlusOne:    "0011111",
	APlusOne:    "0110111",
	MPlusOne:    "1110111",
	DMinusOne:   "0001110",
	AMinusOne:   "0110010",
	MMinusOne:   "1110010",
	DPlusA:      "0000010",
	DPlusM:      "1000010",
	DMinusA:     "0010011",
	DMinusM:     "1010011",
	AMinusD:     "0000111",
	MMinusD:     "1000111",
	DAndA:       "0000000",
	DAndM:       "1000000",
	DOrA:        "0010101",
	DOrM:        "1010101",
}

var destBinary = map[DestinationCode]string{
	NullDest: "000",
	MDest:    "001",
	DDest:    "010",
	DMDest:   "011",
	ADest:    "100",
	AMDest:   "101",
	ADDest:   "110",
	ADMDest:  "111",
}

var jumpBinary = map[JumpCode]string{
	NullJump: "000",
	JGT:      "001",
	JEQ:      "010",
	JGE:      "011",
	JLT:      "100",
	JNE:      "101",
	JLE:      "110",
	JMP:      "111",
}
