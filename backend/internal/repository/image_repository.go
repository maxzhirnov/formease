package repository

import (
	"context"
	"fmt"

	"github.com/maxzhirnov/formease/internal/models"
	"github.com/maxzhirnov/formease/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type ImageRepository interface {
	Create(image *models.Image) error
	FindByUserID(userID string, page, limit int) ([]*models.Image, error)
	FindByID(id string) (*models.Image, error)
	Delete(id string) error
	CountByUserID(userID string) (int64, error)
}

type MongoImageRepository struct {
	db *mongo.Database
}

func NewMongoImageRepository(db *mongo.Database) *MongoImageRepository {
	return &MongoImageRepository{
		db: db,
	}
}

func (r *MongoImageRepository) Create(image *models.Image) error {
	collection := r.db.Collection("images")

	result, err := collection.InsertOne(context.Background(), image)
	if err != nil {
		logger.Error("Failed to insert image", zap.Error(err))
		return err
	}

	image.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *MongoImageRepository) CountByUserID(userID string) (int64, error) {

	collection := r.db.Collection("images")

	count, err := collection.CountDocuments(
		context.Background(),
		bson.M{"userId": userID},
	)

	if err != nil {
		logger.Error("Failed to count user images", zap.Error(err))
		return 0, fmt.Errorf("failed to count user images: %w", err)
	}

	return count, nil
}

func (r *MongoImageRepository) FindByUserID(userID string, page, limit int) ([]*models.Image, error) {
	collection := r.db.Collection("images")

	// Вычисляем skip для пагинации
	skip := (page - 1) * limit

	// Создаем опции для поиска
	options := options.Find().
		SetSort(bson.M{"createdAt": -1}). // Сортировка по дате создания (новые первые)
		SetSkip(int64(skip)).
		SetLimit(int64(limit))

	// Выполняем поиск
	cursor, err := collection.Find(context.Background(),
		bson.M{"userId": userID},
		options,
	)
	if err != nil {
		logger.Error("Failed to find images", zap.Error(err))
		return nil, fmt.Errorf("failed to find images: %w", err)
	}
	defer cursor.Close(context.Background())

	// Декодируем результаты
	var images []*models.Image
	if err = cursor.All(context.Background(), &images); err != nil {
		logger.Error("Failed to decode images", zap.Error(err))
		return nil, fmt.Errorf("failed to decode images: %w", err)
	}

	logger.Info("Successfully found images", zap.Int("count", len(images)))
	return images, nil
}

func (r *MongoImageRepository) FindByID(id string) (*models.Image, error) {
	// Конвертируем id в ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Failed to convert id to ObjectID", zap.Error(err))
		return nil, fmt.Errorf("invalid image ID format: %w", err)
	}

	collection := r.db.Collection("images")

	// Ищем изображение
	var image models.Image
	err = collection.FindOne(context.Background(),
		bson.M{"_id": objectID},
	).Decode(&image)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Error("Image not found", zap.String("id", id))
			return nil, fmt.Errorf("image not found: %s", id)
		}
		logger.Error("Failed to find image", zap.Error(err))
		return nil, fmt.Errorf("failed to find image: %w", err)
	}

	return &image, nil
}

func (r *MongoImageRepository) Delete(id string) error {
	// Конвертируем id в ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Failed to convert id to ObjectID", zap.Error(err))
		return fmt.Errorf("invalid image ID format: %w", err)
	}

	collection := r.db.Collection("images")

	// Удаляем изображение
	result, err := collection.DeleteOne(context.Background(),
		bson.M{"_id": objectID},
	)

	if err != nil {
		logger.Error("Failed to delete image", zap.Error(err))
		return fmt.Errorf("failed to delete image: %w", err)
	}

	// Проверяем, было ли что-то удалено
	if result.DeletedCount == 0 {
		logger.Error("Image not found for deletion", zap.String("id", id))
		return fmt.Errorf("image not found: %s", id)
	}

	return nil
}
