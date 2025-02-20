// storage/file_storage.go
package storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/maxzhirnov/formease/pkg/logger"
	"go.uber.org/zap"
)

type LocalFileStorage struct {
	config FileStorageConfig
}

func NewLocalFileStorage(config FileStorageConfig) *LocalFileStorage {
	// Ensure upload directory exists
	if err := os.MkdirAll(config.UploadDir, 0755); err != nil {
		logger.Error("Failed to create upload directory",
			zap.String("dir", config.UploadDir),
			zap.Error(err))
	}

	return &LocalFileStorage{
		config: config,
	}
}

func (s *LocalFileStorage) Store(filename string, file multipart.File) (string, error) {

	// Validate filename
	if !isValidFilename(filename) {
		logger.Error("Invalid filename")
		return "", fmt.Errorf("invalid filename: %s", filename)
	}

	// Check if file exists
	if s.Exists(filename) {
		logger.Error("File already exists")
		return "", fmt.Errorf("file already exists: %s", filename)
	}

	// Create file
	fullPath := s.GetFullPath(filename)
	dst, err := os.Create(fullPath)
	if err != nil {
		logger.Error("Failed to create file", zap.Error(err))
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	// Copy file content
	if _, err := io.Copy(dst, file); err != nil {
		logger.Error("Failed to copy file content", zap.Error(err))
		os.Remove(fullPath)
		return "", fmt.Errorf("failed to copy file content: %w", err)
	}

	// Return public URL
	return s.GetPublicURL(filename), nil
}

func (s *LocalFileStorage) Delete(filename string) error {

	fullPath := s.GetFullPath(filename)

	// Check file existence
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		logger.Error("File does not exist")
		return fmt.Errorf("file does not exist: %s", filename)
	}

	// Remove file
	if err := os.Remove(fullPath); err != nil {
		logger.Error("Failed to delete file", zap.Error(err))
		return fmt.Errorf("failed to delete file %s: %w", filename, err)
	}

	logger.Info("File deleted successfully")
	return nil
}

func (s *LocalFileStorage) Exists(filename string) bool {
	_, err := os.Stat(s.GetFullPath(filename))
	return !os.IsNotExist(err)
}

func (s *LocalFileStorage) GetFullPath(filename string) string {
	return filepath.Join(s.config.UploadDir, filename)
}

func (s *LocalFileStorage) GetPublicURL(filename string) string {
	// Ensure base URL is clean
	baseURL := strings.TrimRight(s.config.BaseURL, "/")
	return fmt.Sprintf("%s/%s", baseURL, filename)
}

func isValidFilename(filename string) bool {
	// Enhanced filename validation
	if filename == "" {
		return false
	}

	// Check for suspicious characters
	if strings.ContainsAny(filename, "/\\?%*:|\"<>") {
		return false
	}

	// Optional: Check file extension
	ext := filepath.Ext(filename)
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
	}

	return allowedExtensions[strings.ToLower(ext)]
}
