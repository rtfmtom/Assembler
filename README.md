## Related Repositories

[Digital](https://github.com/rtfmtom/Digital)  
[Life-Hack](https://github.com/rtfmtom/Life-Hack)  
[CPU](https://github.com/rtfmtom/CPU)  


# Assembler

Assembler for the Hack CPU.  

A summary reference of the Hack assembly language is provided below. A more in depth explanation of the Hack CPU architecture and instruction set can be found [here](https://github.com/rtfmtom/CPU/blob/main/README.md#cpu).

## Language Reference

Hack assembly has two instruction types: the A-instruction and C-instruction.

### A-instruction 

`@N` where `N` is any 15-bit integer (0 to 32,767 unsigned, -16,384 to 16,383 signed). For example, `@18` loads the 15-bit integer value 18 into the A-Register.

### C-instruction

Can be either `[dest]=[computation]` or `[computation];[jump]`

Valid destinations are:  
`A` — A-Register  
`D` — D-Register  
`M` — Memory location RAM[A] (addressed by the value in the A-Register)  
`null` — Value is not stored
```
D=A     // load the value presently stored in the A-Register to the D-Register
D=D+1   // increment the value stored in the D-Register
@100    // load 100 into the A-Register (selects RAM[100] as a side effect)
M=D     // load the value stored in the D-Register to RAM[100]
```

### Jumps

Jumps are preceded by an A-instruction that specifies which instruction address to jump to. The destination field must be omitted (null) for jump instructions.
```
@10
0;JMP   // unconditional jump to instruction at address 10 in ROM

@4
D+A;JGT // jump to instruction at address 4 in ROM if D+A is greater than 0
```

_Jump Reference Table_

| jump | Effect                      |
|------|-----------------------------|
| null | no jump                     |
| JGT  | if *comp* > 0 jump          |
| JEQ  | if *comp* = 0 jump          |
| JGE  | if *comp* ≥ 0 jump          |
| JLT  | if *comp* < 0 jump          |
| JNE  | if *comp* ≠ 0 jump          |
| JLE  | if *comp* ≤ 0 jump          |
| JMP  | unconditional jump          |

### Operators
```
-  subtraction
+  addition
!  bitwise NOT
&  bitwise AND
|  bitwise OR
```

### Symbols

Symbols are user-defined names that represent memory addresses. They make code more readable by using meaningful names instead of numeric addresses.
```
@counter    // references the memory address assigned to symbol 'counter'
M=0         // initialize counter to 0

@sum        // references the memory address assigned to symbol 'sum'
M=0         // initialize sum to 0
```

The assembler maintains a symbol table and automatically assigns memory addresses to symbols starting at RAM address 16, unless the symbol is predefined.

**Predefined symbols:**

| Symbol       | Address | Description                    |
|--------------|---------|--------------------------------|
| R0-R15       | 0-15    | Virtual registers              |
| SCREEN       | 16384   | Screen memory map base address |
| KBD          | 24576   | Keyboard memory map address    |
| SP           | 0       | Stack pointer                  |
| LCL          | 1       | Local segment base             |
| ARG          | 2       | Argument segment base          |
| THIS         | 3       | This segment base              |
| THAT         | 4       | That segment base              |

### Labels

Labels mark instruction addresses and are defined using parentheses: `(LABEL_NAME)`. They do not generate machine code themselves but serve as symbolic references to ROM addresses for jump instructions.
```
    @counter
    M=0         // initialize counter

(LOOP)          // label marking ROM address of this instruction
    @counter
    M=M+1       // increment counter
    
    @10
    D=A
    @counter
    D=M-D       // check if counter == 10
    
    @LOOP
    D;JLT       // jump back to LOOP if counter < 10
    
    @END
    0;JMP       // jump to END when done

(END)
    @END
    0;JMP       // infinite loop
```

Labels are resolved during the first pass of assembly, and their addresses are stored in the symbol table before the second pass translates instructions to machine code.


