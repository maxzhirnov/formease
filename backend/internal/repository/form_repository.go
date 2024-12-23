package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/maxzhirnov/formease/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FormRepository struct {
	collection *mongo.Collection
}

func NewFormRepository(db *mongo.Database) *FormRepository {
	return &FormRepository{
		collection: db.Collection("forms"),
	}
}

func (r *FormRepository) CreateForm(form *models.Form) error {
	ctx := context.Background()
	result, err := r.collection.InsertOne(ctx, form)
	if err != nil {
		return err
	}
	form.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *FormRepository) GetForm(id string) (*models.Form, error) {
	ctx := context.Background()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var form models.Form
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&form)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("form not found")
		}
		return nil, err
	}
	return &form, nil
}

func (r *FormRepository) ListForms(userID string) ([]models.Form, error) {
	ctx := context.Background()

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID format: %v", err)
	}

	// Add filter for userID
	filter := bson.M{"user_id": objectID}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var forms []models.Form
	if err = cursor.All(ctx, &forms); err != nil {
		return nil, err
	}
	return forms, nil
}

func (r *FormRepository) UpdateForm(form *models.Form) error {
	ctx := context.Background()
	_, err := r.collection.ReplaceOne(
		ctx,
		bson.M{"_id": form.ID},
		form,
	)
	return err
}

func (r *FormRepository) DeleteForm(id string) error {
	ctx := context.Background()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("form not found")
	}
	return nil
}
