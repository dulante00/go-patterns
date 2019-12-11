package data

import "io"

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

type StorageType int

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}


func NewStore(t StorageType) Store {
	switch t {
	case MemoryStorage:
		return newMemoryStorage()
	case DiskStorage:
		return newDiskStorage()
	case TempStorage:
		return newTempStorage()
	default:
		return newTempStorage()
	}
}
