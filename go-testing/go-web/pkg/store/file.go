package store

import (
	"encoding/json"
	"os"
)

type Storage interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type StorageType string

const (
	FileType  StorageType = "file"
	MongoType StorageType = "mongo"
)

type fileStore struct {
	FilePath string
}

func NewStore(storeType StorageType, path string) Storage {
	switch storeType {
	case FileType:
		return &fileStore{path}
	}
	return nil
}

func (fs *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

func (fs *fileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}
