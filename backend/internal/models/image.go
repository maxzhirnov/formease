package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Image struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID      string             `bson:"userId" json:"userId"`
	URL         string             `bson:"url" json:"url"`
	Filename    string             `bson:"filename" json:"filename"`
	Size        int64              `bson:"size" json:"size"`
	Type        string             `bson:"type" json:"type"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	Title       string             `bson:"title,omitempty" json:"title,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
}

type ImageResponse struct {
	Images []*Image `json:"images"`
	Total  int64    `json:"total"`
	Page   int      `json:"page"`
	Limit  int      `json:"limit"`
}
