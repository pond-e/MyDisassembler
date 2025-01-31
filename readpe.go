package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

type WORD uint16
type LONG int32

type IMAGE_DOS_HEADER struct {
	E_magic    WORD
	E_cblp     WORD
	E_cp       WORD
	E_crlc     WORD
	E_cparhdr  WORD
	E_minalloc WORD
	E_maxalloc WORD
	E_ss       WORD
	E_sp       WORD
	E_csum     WORD
	E_ip       WORD
	E_cs       WORD
	E_lfarlc   WORD
	E_ovno     WORD
	E_res      [4]WORD
	E_oemid    WORD
	E_oeminfo  WORD
	E_res2     [10]WORD
	E_lfanew   LONG
}

func ReadImageDosHeader(file *os.File) (*IMAGE_DOS_HEADER, error) {
	dosHeader := new(IMAGE_DOS_HEADER)
	err := binary.Read(file, binary.LittleEndian, dosHeader)
	if err != nil {
		return nil, err
	}
	return dosHeader, nil
}

func ReadPe(file *os.File) {
	dosHeader, err := ReadImageDosHeader(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading Dos header: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("DOS Header:\n")
	fmt.Printf("  Magic number:                 0x%X\n", dosHeader.E_magic)
	fmt.Printf("  Bytes on last page of file:   %d\n", dosHeader.E_cblp)
	fmt.Printf("  Pages in file:                %d\n", dosHeader.E_cp)
	fmt.Printf("  Relocations:                  %d\n", dosHeader.E_crlc)
	fmt.Printf("  Size of header in paragraphs: %d\n", dosHeader.E_cparhdr)
	fmt.Printf("  Minimum extra paragraphs:     %d\n", dosHeader.E_minalloc)
	fmt.Printf("  Maximum extra paragraphs:     %d\n", dosHeader.E_maxalloc)
	fmt.Printf("  Initial (relative) SS value:  0x%X\n", dosHeader.E_ss)
	fmt.Printf("  Initial SP value:             0x%X\n", dosHeader.E_sp)
	fmt.Printf("  Checksum:                     0x%X\n", dosHeader.E_csum)
	fmt.Printf("  Initial IP value:             0x%X\n", dosHeader.E_ip)
	fmt.Printf("  Initial (relative) CS value:  0x%X\n", dosHeader.E_cs)
	fmt.Printf("  File address of relocation table: 0x%X\n", dosHeader.E_lfarlc)
	fmt.Printf("  Overlay number:               %d\n", dosHeader.E_ovno)
	fmt.Printf("  OEM identifier:               %d\n", dosHeader.E_oemid)
	fmt.Printf("  OEM information:              %d\n", dosHeader.E_oeminfo)
	fmt.Printf("  File address of new exe header: 0x%X\n", dosHeader.E_lfanew)
}
