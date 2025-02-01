package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Elf64Ehdr struct {
	Ident     [16]byte
	Type      uint16
	Machine   uint16
	Version   uint32
	Entry     uint64
	Phoff     uint64
	Shoff     uint64
	Flags     uint32
	Ehsize    uint16
	Phentsize uint16
	Phnum     uint16
	Shentsize uint16
	Shnum     uint16
	Shstrndx  uint16
}

type Elf64Phdr struct {
	Type   uint32
	Flags  uint32
	Offset uint64
	Vaddr  uint64
	Paddr  uint64
	Filesz uint64
	Memsz  uint64
	Align  uint64
}

type Elf64Shdr struct {
	Name      uint32
	Type      uint32
	Flags     uint64
	Addr      uint64
	Offset    uint64
	Size      uint64
	Link      uint32
	Info      uint32
	Addralign uint64
	Entsize   uint64
}

type Elf64ShdrWithName struct {
	Name      string
	Type      uint32
	Flags     uint64
	Addr      uint64
	Offset    uint64
	Size      uint64
	Link      uint32
	Info      uint32
	Addralign uint64
	Entsize   uint64
}

func ReadProgramHeaders(file *os.File, ehdr *Elf64Ehdr) []Elf64Phdr {
	file.Seek(int64(ehdr.Phoff), 0)
	phdrs := make([]Elf64Phdr, ehdr.Phnum)
	for i := 0; i < int(ehdr.Phnum); i++ {
		binary.Read(file, binary.LittleEndian, &phdrs[i])
	}
	return phdrs
}

func dumpStringTable(file *os.File, offset, size uint64) []byte {
	file.Seek(int64(offset), 0)
	strData := make([]byte, size)
	binary.Read(file, binary.LittleEndian, &strData)
	return strData
}

func getString(data []byte, index uint32) string {
	end := index
	for end < uint32(len(data)) && data[end] != 0 {
		end++
	}
	return string(data[index:end])
}

func ReadSectionHeaders(file *os.File, ehdr *Elf64Ehdr) []Elf64ShdrWithName {
	return MakeSectionHeaderWithName(file, ehdr)
}

func MakeSectionHeaderWithName(file *os.File, ehdr *Elf64Ehdr) []Elf64ShdrWithName {
	file.Seek(int64(ehdr.Shoff), 0)

	// Load section headers into a slice
	shdrs := make([]Elf64Shdr, ehdr.Shnum)
	shdrwns := make([]Elf64ShdrWithName, ehdr.Shnum)
	for i := 0; i < int(ehdr.Shnum); i++ {
		binary.Read(file, binary.LittleEndian, &shdrs[i])
	}

	// Load the section header string table
	stringTable := dumpStringTable(file, shdrs[ehdr.Shstrndx].Offset, shdrs[ehdr.Shstrndx].Size)

	for i := 0; i < int(ehdr.Shnum); i++ {
		sectionName := getString(stringTable, shdrs[i].Name)
		shdrwns[i].Name = sectionName
		shdrwns[i].Type = shdrs[i].Type
		shdrwns[i].Flags = shdrs[i].Flags
		shdrwns[i].Addr = shdrs[i].Addr
		shdrwns[i].Offset = shdrs[i].Offset
		shdrwns[i].Size = shdrs[i].Size
		shdrwns[i].Link = shdrs[i].Link
		shdrwns[i].Info = shdrs[i].Info
		shdrwns[i].Addralign = shdrs[i].Addralign
		shdrwns[i].Entsize = shdrs[i].Entsize
	}

	return shdrwns
}

func ReadELFHeader(file *os.File) (*Elf64Ehdr, error) {
	ehdr := new(Elf64Ehdr)
	err := binary.Read(file, binary.LittleEndian, ehdr)
	if err != nil {
		return nil, err
	}
	return ehdr, nil
}

func ReadElf(file *os.File) ([]string, []uint64, []uint64, error) {
	ehdr, err := ReadELFHeader(file)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error reading ELF header: %v", err)
	}

	shdrs := ReadSectionHeaders(file, ehdr)

	sectionNames := make([]string, len(shdrs))
	offsets := make([]uint64, len(shdrs))
	sizes := make([]uint64, len(shdrs))

	for i, shdr := range shdrs {
		sectionNames[i] = shdr.Name
		offsets[i] = shdr.Offset
		sizes[i] = shdr.Size
	}

	return sectionNames, offsets, sizes, nil
}
