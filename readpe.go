package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type WORD uint16
type DWORD uint32
type ULONGLONG uint64
type BYTE byte
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

type IMAGE_FILE_HEADER struct {
	Machine              WORD
	NumberOfSections     WORD
	TimeDateStamp        DWORD
	PointerToSymbolTable DWORD
	NumberOfSymbols      DWORD
	SizeOfOptionalHeader WORD
	Characteristics      WORD
}

type IMAGE_OPTIONAL_HEADER64 struct {
	Magic                       WORD
	MajorLinkerVersion          BYTE
	MinorLinkerVersion          BYTE
	SizeOfCode                  DWORD
	SizeOfInitializedData       DWORD
	SizeOfUninitializedData     DWORD
	AddressOfEntryPoint         DWORD
	BaseOfCode                  DWORD
	ImageBase                   ULONGLONG
	SectionAlignment            DWORD
	FileAlignment               DWORD
	MajorOperatingSystemVersion WORD
	MinorOperatingSystemVersion WORD
	MajorImageVersion           WORD
	MinorImageVersion           WORD
	MajorSubsystemVersion       WORD
	MinorSubsystemVersion       WORD
	Win32VersionValue           DWORD
	SizeOfImage                 DWORD
	SizeOfHeaders               DWORD
	CheckSum                    DWORD
	Subsystem                   WORD
	DllCharacteristics          WORD
	SizeOfStackReserve          ULONGLONG
	SizeOfStackCommit           ULONGLONG
	SizeOfHeapReserve           ULONGLONG
	SizeOfHeapCommit            ULONGLONG
	LoaderFlags                 DWORD
	NumberOfRvaAndSizes         DWORD
}

type IMAGE_NT_HEADERS64 struct {
	Signature      DWORD
	FileHeader     IMAGE_FILE_HEADER
	OptionalHeader IMAGE_OPTIONAL_HEADER64
}

func ReadImageDosHeader(file *os.File) (*IMAGE_DOS_HEADER, error) {
	dosHeader := new(IMAGE_DOS_HEADER)
	err := binary.Read(file, binary.LittleEndian, dosHeader)
	if err != nil {
		return nil, err
	}
	return dosHeader, nil
}

func ReadNTHeader(file *os.File, offset int64) (*IMAGE_NT_HEADERS64, error) {
	// Move the file pointer to the NT header offset
	_, err := file.Seek(offset, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("error seeking to NT header: %v", err)
	}

	ntHeader := new(IMAGE_NT_HEADERS64)
	// Read the NT header
	err = binary.Read(file, binary.LittleEndian, ntHeader)
	if err != nil {
		return nil, fmt.Errorf("error reading NT header: %v", err)
	}

	return ntHeader, nil
}

func ReadPe(file *os.File) {
	dosHeader, err := ReadImageDosHeader(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading DOS header: %v\n", err)
		os.Exit(1)
	}
	ntHeader, err := ReadNTHeader(file, int64(dosHeader.E_lfanew))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading NT header: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("NT Header:\n")
	fmt.Printf("  Signature:                    0x%X\n", ntHeader.Signature)
	fmt.Printf("  File Header:\n")
	fmt.Printf("    Machine:                   0x%X\n", ntHeader.FileHeader.Machine)
	fmt.Printf("    Number of Sections:        %d\n", ntHeader.FileHeader.NumberOfSections)
	fmt.Printf("    Time Date Stamp:           0x%X\n", ntHeader.FileHeader.TimeDateStamp)
	fmt.Printf("    Pointer to Symbol Table:   0x%X\n", ntHeader.FileHeader.PointerToSymbolTable)
	fmt.Printf("    Number of Symbols:         %d\n", ntHeader.FileHeader.NumberOfSymbols)
	fmt.Printf("    Size of Optional Header:   %d\n", ntHeader.FileHeader.SizeOfOptionalHeader)
	fmt.Printf("    Characteristics:           0x%X\n", ntHeader.FileHeader.Characteristics)
	fmt.Printf("  Optional Header:\n")
	fmt.Printf("    Magic:                     0x%X\n", ntHeader.OptionalHeader.Magic)
	fmt.Printf("    Major Linker Version:      %d\n", ntHeader.OptionalHeader.MajorLinkerVersion)
	fmt.Printf("    Minor Linker Version:      %d\n", ntHeader.OptionalHeader.MinorLinkerVersion)
	fmt.Printf("    Size of Code:              %d\n", ntHeader.OptionalHeader.SizeOfCode)
	fmt.Printf("    Size of Initialized Data:  %d\n", ntHeader.OptionalHeader.SizeOfInitializedData)
	fmt.Printf("    Size of Uninitialized Data:%d\n", ntHeader.OptionalHeader.SizeOfUninitializedData)
	fmt.Printf("    Address of Entry Point:    0x%X\n", ntHeader.OptionalHeader.AddressOfEntryPoint)
	fmt.Printf("    Base of Code:              0x%X\n", ntHeader.OptionalHeader.BaseOfCode)
	fmt.Printf("    Image Base:                0x%X\n", ntHeader.OptionalHeader.ImageBase)
	fmt.Printf("    Section Alignment:         %d\n", ntHeader.OptionalHeader.SectionAlignment)
	fmt.Printf("    File Alignment:            %d\n", ntHeader.OptionalHeader.FileAlignment)
	fmt.Printf("    Major Operating System Version: %d\n", ntHeader.OptionalHeader.MajorOperatingSystemVersion)
	fmt.Printf("    Minor Operating System Version: %d\n", ntHeader.OptionalHeader.MinorOperatingSystemVersion)
	fmt.Printf("    Major Image Version:       %d\n", ntHeader.OptionalHeader.MajorImageVersion)
	fmt.Printf("    Minor Image Version:       %d\n", ntHeader.OptionalHeader.MinorImageVersion)
	fmt.Printf("    Major Subsystem Version:   %d\n", ntHeader.OptionalHeader.MajorSubsystemVersion)
	fmt.Printf("    Minor Subsystem Version:   %d\n", ntHeader.OptionalHeader.MinorSubsystemVersion)
	fmt.Printf("    Win32 Version Value:       0x%X\n", ntHeader.OptionalHeader.Win32VersionValue)
	fmt.Printf("    Size of Image:             %d\n", ntHeader.OptionalHeader.SizeOfImage)
	fmt.Printf("    Size of Headers:           %d\n", ntHeader.OptionalHeader.SizeOfHeaders)
	fmt.Printf("    CheckSum:                  0x%X\n", ntHeader.OptionalHeader.CheckSum)
	fmt.Printf("    Subsystem:                 0x%X\n", ntHeader.OptionalHeader.Subsystem)
	fmt.Printf("    DLL Characteristics:       0x%X\n", ntHeader.OptionalHeader.DllCharacteristics)
	fmt.Printf("    Size of Stack Reserve:     0x%X\n", ntHeader.OptionalHeader.SizeOfStackReserve)
	fmt.Printf("    Size of Stack Commit:      0x%X\n", ntHeader.OptionalHeader.SizeOfStackCommit)
	fmt.Printf("    Size of Heap Reserve:      0x%X\n", ntHeader.OptionalHeader.SizeOfHeapReserve)
	fmt.Printf("    Size of Heap Commit:       0x%X\n", ntHeader.OptionalHeader.SizeOfHeapCommit)
	fmt.Printf("    Loader Flags:              0x%X\n", ntHeader.OptionalHeader.LoaderFlags)
	fmt.Printf("    Number of RVA and Sizes:   %d\n", ntHeader.OptionalHeader.NumberOfRvaAndSizes)
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
