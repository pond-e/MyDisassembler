package main

import (
	"os"
)

type Memory struct {
	dump []byte
}

func NewMemory() *Memory {
	return &Memory{
		dump: nil,
	}
}

func (m *Memory) ReadDump(filePath string) (*Memory, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	m.dump = data
	return m, nil
}

func (m *Memory) Read(entry uint64, size int) uint64 {
	value := uint64(0)
	for i := 0; i < size; i++ {
		// Ensure we don't read past the end of dump
		if int(entry)+i >= len(m.dump) {
			break
		}
		// Shift and combine bytes
		value |= uint64(m.dump[entry+uint64(i)]) << (i * size * 2)
	}
	return value
}
