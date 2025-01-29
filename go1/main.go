package main

import (
	"fmt"
)

func main() {
	fmt.Println("x64 Disassembler")
	disassemble([]byte{0x48, 0x89, 0xe5}) // Example: mov rbp, rsp
}

func disassemble(code []byte) {
	// Simple example of disassembly logic
	if len(code) == 3 && code[0] == 0x48 && code[1] == 0x89 && code[2] == 0xe5 {
		fmt.Println("mov rbp, rsp")
	} else {
		fmt.Println("Unknown instruction")
	}
}
