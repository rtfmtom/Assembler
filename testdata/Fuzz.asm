// Comprehensive Hack CPU Instruction Test
// Tests all computation, destination, and jump instruction types

// Initialize test values
@100
D=A
@200
M=D
M=M+1

// Test all computations with D destination
@0
D=0
D=1
D=-1
D=D
D=A
@200
D=M
D=!D
D=!A
@200
D=!M
D=-D
D=-A
@200
D=-M
D=D+1
D=A+1
@200
D=M+1
D=D-1
D=A-1
@200
D=M-1
D=D+A
@200
D=D+M
D=D-A
@200
D=D-M
D=A-D
@200
D=M-D
D=D&A
@200
D=D&M
D=D|A
@200
D=D|M

// Test all destination types
@300
M=1
D=1
MD=1
A=1
AM=1
AD=1
AMD=1

// Test jump conditions
@10
D=A
@JGT_TEST
D;JGT
(JGT_TEST)

@0
D=A
@JEQ_TEST
D;JEQ
(JEQ_TEST)

@5
D=A
@JGE_TEST
D;JGE
(JGE_TEST)

@5
D=A
D=-D
@JLT_TEST
D;JLT
(JLT_TEST)

@5
D=A
@JNE_TEST
D;JNE
(JNE_TEST)

D=-1
@JLE_TEST
D;JLE
(JLE_TEST)

@JMP_TEST
0;JMP
(JMP_TEST)

// End program infinite loop
(END)
@END
0;JMP