package main

func assemble(lines [][]byte) ([]string, error) {

	symbolTable, err := buildSymbolTable(lines)
	if err != nil {
		return []string{}, err
	}

	var instructions []string

	for _, line := range lines {
		line = cleanLine(line)
		if len(line) == 0 || isLabel(line) {
			continue
		}

		if line[0] == '@' {
			var a AInstruction
			result, err := a.Encode(instructions, line, symbolTable)
			if err != nil {
				return []string{}, err
			}
			instructions = result
		} else {
			var c CInstruction
			result, err := c.Encode(instructions, line)
			if err != nil {
				return []string{}, err
			}
			instructions = result
		}
	}

	return instructions, nil
}
