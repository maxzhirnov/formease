package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Form struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID          primitive.ObjectID `bson:"userId" json:"userId"`
	IsDraft         bool               `bson:"isDraft" json:"isDraft"`
	Name            string             `bson:"name" json:"name"`
	Theme           string             `bson:"theme" json:"theme"`
	FloatingShapes  string             `bson:"floatingShapesTheme" json:"floatingShapesTheme"`
	Questions       []Question         `bson:"questions" json:"questions"`
	ThankYouMessage ThankYouMessage    `bson:"thankYouMessage" json:"thankYouMessage"`
}

type Question struct {
	ID       int    `bson:"id" json:"id"`
	Type     string `bson:"type" json:"type"`
	Question string `bson:"question" json:"question"`
	Subtext  string `bson:"subtext" json:"subtext"`
	Image    string `bson:"image" json:"image"`
	// Input question fields
	InputType   string `bson:"inputType,omitempty" json:"inputType,omitempty"`
	Placeholder string `bson:"placeholder,omitempty" json:"placeholder,omitempty"`
	Validation  string `bson:"validation,omitempty" json:"validation,omitempty"`
	// Choice question fields
	MaxSelections int      `bson:"maxSelections,omitempty" json:"maxSelections,omitempty"`
	Options       []Option `bson:"options,omitempty" json:"options,omitempty"`
	// Rating question fields
	MinValue   int     `bson:"minValue,omitempty" json:"minValue,omitempty"`
	MaxValue   int     `bson:"maxValue,omitempty" json:"maxValue,omitempty"`
	Step       float64 `bson:"step,omitempty" json:"step,omitempty"`
	ShowLabels bool    `bson:"showLabels,omitempty" json:"showLabels,omitempty"`
	MinLabel   string  `bson:"minLabel,omitempty" json:"minLabel,omitempty"`
	MaxLabel   string  `bson:"maxLabel,omitempty" json:"maxLabel,omitempty"`
	Icon       string  `bson:"icon,omitempty" json:"icon,omitempty"`
	// Common fields
	NextQuestion NextQuestion `bson:"nextQuestion" json:"nextQuestion"`
}

type Option struct {
	ID    int    `bson:"id" json:"id"`
	Text  string `bson:"text" json:"text"`
	Icon  string `bson:"icon" json:"icon"`
	Image string `bson:"image" json:"image"`
}

type NextQuestion struct {
	Conditions []Condition `bson:"conditions" json:"conditions"`
	Default    int         `bson:"default,omitempty" json:"default,omitempty"`
}

type Condition struct {
	Answer string `bson:"answer" json:"answer"`
	NextID int    `bson:"nextId" json:"nextId"`
}

type ThankYouMessage struct {
	Title    string       `bson:"title" json:"title"`
	Subtitle string       `bson:"subtitle" json:"subtitle"`
	Icon     string       `bson:"icon" json:"icon"`
	Button   ButtonConfig `bson:"button" json:"button"`
}

type ButtonConfig struct {
	Text   string `bson:"text" json:"text"`
	URL    string `bson:"url" json:"url"`
	NewTab bool   `bson:"newTab" json:"newTab"`
}
