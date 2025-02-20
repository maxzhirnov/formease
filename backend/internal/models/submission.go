package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Submission struct {
	ID        primitive.ObjectID `json:"id"`
	Answers   string             `json:"answers"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
