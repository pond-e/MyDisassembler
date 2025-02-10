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

	// INSERT_YOUR_CODE
	// Determine if the file is a PE or ELF file
	// fileHeader := make([]byte, 4)
	// _, err = file.Read(fileHeader)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error reading file header: %v\n", err)
	// 	os.Exit(1)
	// }

	var entry uint64
	var end uint64

	var sectionNames []string
	var offsets, sizes []uint64

	if memory.dump[0] == 0x4d {
		peOffset, size := ReadPe(file)
		entry = uint64(peOffset)
		end = entry + uint64(size)
		// You can add additional logic here if needed for PE files
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
	// Find the section containing the entry point
	// var sectionSize uint64
	// for i, offset := range offsets {
	// 	if entry >= offset && entry < offset+sizes[i] {
	// 		sectionSize = sizes[i]
	// 		break
	// 	}
	// }
	// fmt.Printf("entry: %x\n", entry)
	for entry < end {
		state := NewState(memory.dump, entry)
		state.step(entry)
		// Format the output to match Intel syntax
		fmt.Printf("%s %s\n", state.disassembledInstruction[len(state.disassembledInstruction)-1],
			strings.Join(state.disassembledOperands, ", "))
		entry += state.disassembledInstructionSize
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
