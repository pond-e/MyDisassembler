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

type IMAGE_SECTION_HEADER struct {
	Name                 [8]BYTE
	VirtualSize          DWORD
	VirtualAddress       DWORD
	SizeOfRawData        DWORD
	PointerToRawData     DWORD
	PointerToRelocations DWORD
	PointerToLinenumbers DWORD
	NumberOfRelocations  WORD
	NumberOfLinenumbers  WORD
	Characteristics      DWORD
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

func ReadSectionHeader(file *os.File, ntHeader *IMAGE_NT_HEADERS64, dosHeader *IMAGE_DOS_HEADER) ([]IMAGE_SECTION_HEADER, error) {
	// Calculate the offset to the first section header
	sectionHeaderOffset := int64(ntHeader.FileHeader.SizeOfOptionalHeader) + int64(dosHeader.E_lfanew) + int64(binary.Size(IMAGE_FILE_HEADER{})) + int64(binary.Size(DWORD(0)))
	// fmt.Printf("sectionHeaderOffset: 0x%x\n", sectionHeaderOffset)

	// Move the file pointer to the section header offset
	_, err := file.Seek(sectionHeaderOffset, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("error seeking to section header: %v", err)
	}

	// Read the section headers
	sectionHeaders := make([]IMAGE_SECTION_HEADER, ntHeader.FileHeader.NumberOfSections)
	for i := 0; i < int(ntHeader.FileHeader.NumberOfSections); i++ {
		err = binary.Read(file, binary.LittleEndian, &sectionHeaders[i])
		if err != nil {
			return nil, fmt.Errorf("error reading section header: %v", err)
		}
	}

	return sectionHeaders, nil
}

func ReadPe(file *os.File) (DWORD, DWORD) {
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

	// Read section headers
	// sectionHeaders, err := ReadSectionHeader(file, ntHeader)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error reading section headers: %v\n", err)
	// 	os.Exit(1)
	// }

	sectionHeaders, err := ReadSectionHeader(file, ntHeader, dosHeader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading section headers: %v\n", err)
		os.Exit(1)
	}

	// Display the SizeOfRawData and PointerToRawData for each section header
	// for i, sectionHeader := range sectionHeaders {
	// 	fmt.Printf("Section %d: VirtualSize: 0x%x, PointerToRawData: 0x%x\n", i+1, sectionHeader.VirtualSize, sectionHeader.PointerToRawData)
	// }

	return sectionHeaders[0].PointerToRawData, sectionHeaders[0].VirtualSize
}
