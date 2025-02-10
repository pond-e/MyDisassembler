package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Check if file path and entry point are provided
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <file_path>")
	}

	memory := NewMemory()
	memory.ReadDump(os.Args[1])

	// disassemble all sections
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}

	var entry uint64
	var end uint64

	var sectionNames []string
	var offsets, sizes []uint64

	if memory.dump[0] == 0x4d {
		peOffset, size := ReadPe(file)
		entry = uint64(peOffset)
		end = entry + uint64(size)
	} else if memory.dump[0] == 0x7f {
		sectionNames, offsets, sizes, err = ReadElf(file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(sectionNames[1])
		entry = offsets[1]
		end = entry + sizes[1]
	} else {
		fmt.Fprintf(os.Stderr, "Unknown file format\n")
		os.Exit(1)
	}

	// fmt.Printf("entry: %x\n", entry)
	for entry < end {
		state := NewState(memory.dump, entry)
		state.step(entry)
		// Format the output to match Intel syntax
		fmt.Printf("%s %s\n", state.disassembledInstruction[len(state.disassembledInstruction)-1],
			strings.Join(state.disassembledOperands, ", "))
		entry += state.disassembledInstructionSize
	}
}
