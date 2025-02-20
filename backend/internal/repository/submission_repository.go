package repository

import (
	"context"

	"github.com/maxzhirnov/formease/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SubmissionRepository struct {
	collection *mongo.Collection
}

func NewSubmissionRepository(db *mongo.Database) *SubmissionRepository {
	return &SubmissionRepository{
		collection: db.Collection("submissions"),
	}
}

func (r *SubmissionRepository) CreateSubmission(s *models.Submission) error {
	ctx := context.Background()
	result, err := r.collection.InsertOne(ctx, s)
	if err != nil {
		return err
	}

	s.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}
