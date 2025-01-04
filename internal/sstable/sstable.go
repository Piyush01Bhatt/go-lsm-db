package sstable

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/Piyush01Bhatt/go-lsm-db/internal/skiplist"
)

type SSTable struct {
	dataFile  *os.File
	indexFile *os.File
	memtable  *skiplist.Skiplist
}

func NewSSTable(memtable *skiplist.Skiplist, dataFilePath, indexFilePath string) (*SSTable, error) {
	dataFile, err := os.Create(dataFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create data file: %v", err)
	}

	indexFile, err := os.Create(indexFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create index file: %v", err)
	}

	return &SSTable{
		dataFile:  dataFile,
		indexFile: indexFile,
		memtable:  memtable,
	}, nil
}

func (sstab *SSTable) Write() error {
	// Initialize variables for data and index writing
	offset := int64(0)

	for it := sstab.memtable.Iterator(); it.HasNext(); it.Next() {
		key, value := it.KeyValue()

		// Write the key length and key (binary encoding)
		keyLen := uint32(len(key))
		if err := binary.Write(sstab.dataFile, binary.LittleEndian, keyLen); err != nil {
			return fmt.Errorf("failed to write keylength: %v", err)
		}
		if _, err := sstab.dataFile.Write([]byte(key)); err != nil {
			return fmt.Errorf("failed to write key bytes: %v", err)
		}

		// Write the value length and value (binary encoding)
		valueLen := uint32(len(value))
		if err := binary.Write(sstab.dataFile, binary.LittleEndian, valueLen); err != nil {
			return fmt.Errorf("failed to write valuelength: %v", err)
		}
		if _, err := sstab.dataFile.Write([]byte(value)); err != nil {
			return fmt.Errorf("failed to write value bytes: %v", err)
		}

		// Write the key's index (key offset) to the index file
		indexEntry := fmt.Sprintf("%s:%d", key, offset)
		if _, err := sstab.indexFile.Write([]byte(indexEntry + "\n")); err != nil {
			return fmt.Errorf("failed to write index entry: %v", err)
		}

		// Update the offset
		// size of key, value length, and value
		offset += int64(4 + len(key) + 4 + len(value))
	}
	return nil
}

// Close closes the SSTable files
func (sstab *SSTable) Close() error {
	if err := sstab.dataFile.Close(); err != nil {
		return err
	}
	if err := sstab.indexFile.Close(); err != nil {
		return err
	}
	return nil
}
