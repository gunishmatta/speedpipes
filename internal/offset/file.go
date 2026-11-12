package offset

import (
	"context"
	"errors"
	"os"
	"strconv"
	"strings"
)

type FileStore struct {
	filePath string
}

func NewFileStore(filePath string) *FileStore {
	return &FileStore{
		filePath: filePath,
	}
}

func (fs *FileStore) Load(ctx context.Context) (int64, error) {
	data, err := os.ReadFile(fs.filePath)
	if errors.Is(err, os.ErrNotExist) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(strings.TrimSpace(string(data)), 10, 64)
}

func (fs *FileStore) Save(ctx context.Context, offset int64) error {
	data := []byte(strconv.FormatInt(offset, 10))
	return os.WriteFile(fs.filePath, data, 0644)
}
