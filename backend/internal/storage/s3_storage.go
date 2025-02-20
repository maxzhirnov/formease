package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/maxzhirnov/formease/pkg/logger"
	"go.uber.org/zap"
)

type YandexS3Storage struct {
	client     *s3.Client
	bucketName string
	baseURL    string
}

func NewYandexS3Storage(config FileStorageConfig) (*YandexS3Storage, error) {
	logger.Info("Creating Yandex S3 storage", zap.Any("config", config))
	// Создаем кастомный resolver для Яндекс.Облака
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               "https://storage.yandexcloud.net", // Именно так
			SigningRegion:     config.Region,
			HostnameImmutable: true,
		}, nil
	})

	// Создаем провайдер креденшалов
	credProvider := credentials.NewStaticCredentialsProvider(
		config.AccessKeyID,
		config.SecretAccessKey,
		"",
	)

	// Загружаем AWS конфиг
	awsConfig, err := awsConfig.LoadDefaultConfig(
		context.TODO(),
		awsConfig.WithRegion(config.Region),
		awsConfig.WithCredentialsProvider(credProvider),
		awsConfig.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	// Создаем S3 клиент
	client := s3.NewFromConfig(awsConfig)

	return &YandexS3Storage{
		client:     client,
		bucketName: config.BucketName,
		baseURL:    fmt.Sprintf("%s/%s", config.Endpoint, config.BucketName),
	}, nil
}

func (s *YandexS3Storage) Store(filename string, file multipart.File) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Загрузка файла
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		logger.Error("Failed to upload to Yandex S3",
			zap.String("filename", filename),
			zap.Error(err))
		return "", fmt.Errorf("upload failed: %w", err)
	}

	return s.GetPublicURL(filename), nil
}

func (s *YandexS3Storage) Delete(filename string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(filename),
	})

	if err != nil {
		logger.Error("Failed to delete from Yandex S3",
			zap.String("filename", filename),
			zap.Error(err))
		return fmt.Errorf("deletion failed: %w", err)
	}

	return nil
}

func (s *YandexS3Storage) Exists(filename string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(filename),
	})

	return err == nil
}

func (s *YandexS3Storage) GetFullPath(filename string) string {
	return filename
}

func (s *YandexS3Storage) GetPublicURL(filename string) string {
	return fmt.Sprintf("%s/%s", s.baseURL, filename)
}
