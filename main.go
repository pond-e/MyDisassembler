package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		// Check if file path and entry point are provided
		if len(os.Args) < 3 {
			log.Fatal("Usage: go run main.go <file_path> <entry>")
		}

		// Read the file path from command line argument
		data, err := os.ReadFile(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		// Parse entry point
		var entry uint64
		_, err = fmt.Sscanf(os.Args[2], "%x", &entry)
		if err != nil {
			log.Fatal("Entry point must be a hexadecimal number")
		}

		if entry >= 0x400000 {
			entry -= 0x400000
		}

		// Validate entry point
		if entry >= uint64(len(data)) {
			log.Fatal("Entry point exceeds file size")
		}

		memory := NewMemory()
		memory.ReadDump(os.Args[1])

		// get all sections entry points

		// disassemble all sections
		end := entry + 8
		// get end address
		for entry <= end {
			state := NewState(memory.dump, entry)
			state.step(entry)
			// Format the output to match Intel syntax
			fmt.Printf("%s %s\n", state.disassembledInstruction[len(state.disassembledInstruction)-1],
				strings.Join(state.disassembledOperands, ", "))
			entry += state.disassembledInstructionSize
		}
	*/
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <pe-file>\n", os.Args[0])
		os.Exit(1)
	}

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	ReadPe(file)
}
