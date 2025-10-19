package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
)

var outputFile string
var outputFormat string

func isValidOutputFormats(outputFormat string, validOutputFormats []string) bool {
	for _, validFormat := range validOutputFormats {
		if outputFormat == validFormat {
			return true
		}
	}
	return false
}

func main() {
	flag.StringVar(&outputFile, "output", "out.txt", "output file (shorthand: -o)")
	flag.StringVar(&outputFile, "o", "out.txt", "output file")
	flag.StringVar(&outputFormat, "mode", "hex", "output mode: hex, bin (shorthand: -m)")
	flag.StringVar(&outputFormat, "m", "hex", "output mode: hex, bin")
	flag.Parse()

	validOutputFormats := []string{"hex", "bin", "x", "b"}
	if !isValidOutputFormats(outputFormat, validOutputFormats) {
		fmt.Fprintf(os.Stderr, "Error: Invalid output format '%s'. Valid output formats: %v\n", outputFormat, validOutputFormats)
		flag.Usage()
		os.Exit(1)
	}

	if len(flag.Args()) < 1 {
		log.Fatal("Input file name required")
	}

	data, err := os.ReadFile(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	lines := bytes.Split(data, []byte("\n"))
	instructions, err := assemble(lines)
	if err != nil {
		log.Fatal(err)
	}

	for _, in := range instructions {
		fmt.Println(in)
	}
}
