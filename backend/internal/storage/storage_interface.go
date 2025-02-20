package storage

import (
	"mime/multipart"
)

type FileStorageConfig struct {
	UploadDir string
	BaseURL   string

	// Добавленные поля для S3
	Endpoint        string
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
}

type FileStorage interface {
	Store(filename string, file multipart.File) (fileURL string, err error)
	Delete(filename string) error
	Exists(filename string) bool
	GetFullPath(filename string) string
	GetPublicURL(filename string) string
}
