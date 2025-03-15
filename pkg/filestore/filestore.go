package filestore

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type FileStore struct {
	basePath string
}

func NewFileStore(basePath string) *FileStore {
	return &FileStore{
		basePath: basePath,
	}
}

func (fs *FileStore) SaveFile(file *multipart.FileHeader, movieID uint) (string, string, error) {
	// Create directory if not exists
	uploadDir := filepath.Join(fs.basePath, fmt.Sprintf("movie_%d", movieID))
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", "", err
	}

	// Open uploaded file
	src, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer src.Close()

	// Calculate SHA-256 hash
	hash := sha256.New()
	if _, err := io.Copy(hash, src); err != nil {
		return "", "", err
	}
	fileHash := hex.EncodeToString(hash.Sum(nil))

	// Reset file pointer
	src.Seek(0, 0)

	// Create safe filename
	ext := strings.ToLower(filepath.Ext(file.Filename))
	safeName := fmt.Sprintf("%s%s", fileHash[:12], ext)
	filePath := filepath.Join(uploadDir, safeName)

	// Create destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", "", err
	}
	defer dst.Close()

	// Copy file
	src.Seek(0, 0)
	if _, err = io.Copy(dst, src); err != nil {
		return "", "", err
	}

	return filePath, fileHash, nil
}

func (fs *FileStore) DeleteFile(filePath string) error {
	return os.Remove(filePath)
}



// links when i read 

// https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html


// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type

// https://go.dev/doc/database/sql-injection

