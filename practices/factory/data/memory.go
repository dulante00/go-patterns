package data

import (
	"fmt"
	"io"
)

type memory struct {
	name string
	rawMap map[string]*memRawData
}

type memRawData struct {
	rawData []byte
}

func newMemoryStorage() Store {
	return &memory{}
}

func (m *memory) Read(p []byte) (n int, err error) {
	n = len(m.rawMap[m.name].rawData)
	fmt.Printf("rawData: %v\n", n)
	return n, nil
}


func (m *memory)  Write(p []byte) (n int, err error) {
	rd := &memRawData{
		rawData: p,
	}
	m.rawMap[m.name] = rd
	n = len(p)
	fmt.Printf("<memRawData>: %v", string(m.rawMap[m.name].rawData))
	return n , nil
}

func (m *memory) Close() error {
	return nil
}

func (m *memory)  Open(name string) (io.ReadWriteCloser, error) {
	_m := map[string]*memRawData{
		name: {[]byte{}},
	}
	return &memory{name:name, rawMap: _m}, nil
}