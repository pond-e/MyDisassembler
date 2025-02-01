package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
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
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	sectionNames, offsets, sizes, err := ReadElf(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sectionNames[1])
	entry = offsets[1]
	fmt.Printf("entry: %x\n", entry)
	fmt.Printf("sizes[1]: %x\n", sizes[1])
	// Find the section containing the entry point
	// var sectionSize uint64
	// for i, offset := range offsets {
	// 	if entry >= offset && entry < offset+sizes[i] {
	// 		sectionSize = sizes[i]
	// 		break
	// 	}
	// }

	// Set end based on the section size
	end := entry + sizes[1]

	// get end address
	for entry < end {
		state := NewState(memory.dump, entry)
		state.step(entry)
		// Format the output to match Intel syntax
		fmt.Printf("%s %s\n", state.disassembledInstruction[len(state.disassembledInstruction)-1],
			strings.Join(state.disassembledOperands, ", "))
		entry += state.disassembledInstructionSize
		fmt.Printf("state.disassembledInstructionSize: %x\n", state.disassembledInstructionSize)
	}

	/*
		// readpe, readelf test
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
	*/
}
