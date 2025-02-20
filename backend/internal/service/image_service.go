package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/maxzhirnov/formease/internal/models"
	"github.com/maxzhirnov/formease/internal/repository"
	"github.com/maxzhirnov/formease/internal/storage"
	"github.com/maxzhirnov/formease/pkg/logger"
	"go.uber.org/zap"
)

const (
	maxFileSize = 10 << 20 // 10 MB
	minFileSize = 1 << 10  // 1 KB
)

var allowedMimeTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/gif":  true,
	"image/webp": true,
}

type ImageService struct {
	repo      repository.ImageRepository
	fileStore storage.FileStorage
}

func NewImageService(repo repository.ImageRepository, fileStore storage.FileStorage) *ImageService {
	return &ImageService{
		repo:      repo,
		fileStore: fileStore,
	}
}

func (s *ImageService) UploadImage(image *models.Image, file multipart.File) (*models.Image, error) {

	// Проверка квоты пользователя
	if err := s.checkUserQuota(image.UserID); err != nil {
		return nil, err
	}

	// Валидация файла
	if err := s.validateImage(file, image); err != nil {
		return nil, err
	}

	// Генерация уникального имени файла
	filename := s.generateUniqueFilename(image.Filename)
	image.Filename = filename

	// Загрузка файла в хранилище
	fileURL, err := s.fileStore.Store(filename, file)
	if err != nil {
		logger.Error("Failed to store file", zap.Error(err))
		return nil, err
	}
	image.URL = fileURL

	// Сохранение информации в базе данных
	if err := s.repo.Create(image); err != nil {
		// В случае ошибки удаляем загруженный файл
		if deleteErr := s.fileStore.Delete(filename); deleteErr != nil {
			logger.Error("Failed to delete file after db error", zap.Error(deleteErr))
		}
		return nil, err
	}

	return image, nil
}

// validateImage проверяет размер и тип файла
func (s *ImageService) validateImage(file multipart.File, image *models.Image) error {

	// Проверка размера файла
	if image.Size > maxFileSize {
		logger.Error("File too large",
			zap.Int64("size", image.Size),
			zap.Int64("maxSize", maxFileSize))
		return errors.New("file size exceeds maximum limit of 10MB")
	}

	if image.Size < minFileSize {
		logger.Error("File too small",
			zap.Int64("size", image.Size),
			zap.Int64("minSize", minFileSize))
		return errors.New("file size below minimum limit of 1KB")
	}

	// Проверка типа файла
	buff := make([]byte, 512)
	_, err := file.Read(buff)
	if err != nil {
		logger.Error("Failed to read file header", zap.Error(err))
		return errors.New("failed to read file")
	}

	// Важно вернуть указатель чтения в начало файла
	if _, err := file.Seek(0, 0); err != nil {
		logger.Error("Failed to reset file pointer", zap.Error(err))
		return errors.New("failed to process file")
	}

	// Определение MIME типа
	filetype := http.DetectContentType(buff)
	if !allowedMimeTypes[filetype] {
		logger.Error("Invalid file type",
			zap.String("type", filetype),
			zap.Any("allowedTypes", allowedMimeTypes))
		return errors.New("invalid file type. Only JPEG, PNG, GIF and WebP are allowed")
	}

	// Проверка расширения файла
	ext := strings.ToLower(filepath.Ext(image.Filename))
	if !isValidExtension(ext) {
		logger.Error("Invalid file extension", zap.String("extension", ext))
		return errors.New("invalid file extension")
	}

	return nil
}

// generateUniqueFilename создает уникальное имя файла
func (s *ImageService) generateUniqueFilename(originalFilename string) string {
	// Получаем расширение файла
	ext := filepath.Ext(originalFilename)

	// Создаем имя файла из timestamp и UUID
	timestamp := time.Now().Format("20060102-150405")
	uniqueID := uuid.New().String()[:8]

	// Формируем имя файла: timestamp_uniqueID.extension
	return fmt.Sprintf("%s_%s%s", timestamp, uniqueID, ext)
}

// Вспомогательная функция для проверки расширения
func isValidExtension(ext string) bool {
	validExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
	}
	return validExtensions[ext]
}

// Добавим также метод для проверки квоты пользователя (опционально)
func (s *ImageService) checkUserQuota(userID string) error {

	// Получаем количество изображений пользователя
	count, err := s.repo.CountByUserID(userID)
	if err != nil {
		logger.Error("Failed to get user images count", zap.Error(err))
		return err
	}

	// Проверяем квоту (например, максимум 100 изображений)
	if count >= 100 {
		logger.Error("User quota exceeded",
			zap.String("userId", userID),
			zap.Int64("imageCount", count))
		return errors.New("image quota exceeded")
	}

	return nil
}

func (s *ImageService) FindByUserID(userID string, page, limit int) ([]*models.Image, error) {

	images, err := s.repo.FindByUserID(userID, page, limit)
	if err != nil {
		logger.Error("Failed to find images", zap.Error(err))
		return nil, fmt.Errorf("failed to find images: %w", err)
	}

	return images, nil
}

func (s *ImageService) CountByUserID(userID string) (int64, error) {
	return s.repo.CountByUserID(userID)
}

func (s *ImageService) FindByID(imageID string) (*models.Image, error) {
	// Retrieve the image from the repository
	image, err := s.repo.FindByID(imageID)
	if err != nil {
		logger.Error("Failed to find image by ID",
			zap.String("imageID", imageID),
			zap.Error(err))
		return nil, fmt.Errorf("failed to find image: %w", err)
	}

	return image, nil
}

func (s *ImageService) Delete(imageID string) error {
	// First, find the image to get its filename
	image, err := s.FindByID(imageID)
	if err != nil {
		return err
	}

	// Delete the file from storage
	if err := s.fileStore.Delete(image.Filename); err != nil {
		logger.Error("Failed to delete file from storage",
			zap.String("filename", image.Filename),
			zap.Error(err))
		// Continue with database deletion even if file deletion fails
	}

	// Delete the image record from the database
	if err := s.repo.Delete(imageID); err != nil {
		logger.Error("Failed to delete image from database",
			zap.String("imageID", imageID),
			zap.Error(err))
		return fmt.Errorf("failed to delete image: %w", err)
	}

	logger.Info("Image deleted successfully",
		zap.String("imageID", imageID),
		zap.String("filename", image.Filename))

	return nil
}
