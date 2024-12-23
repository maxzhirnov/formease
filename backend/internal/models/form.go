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
	ID            int          `bson:"id" json:"id"`
	Type          string       `bson:"type" json:"type"`
	Question      string       `bson:"question" json:"question"`
	Subtext       string       `bson:"subtext" json:"subtext"`
	Image         string       `bson:"image" json:"image"`
	InputType     string       `bson:"inputType,omitempty" json:"inputType,omitempty"`
	Placeholder   string       `bson:"placeholder,omitempty" json:"placeholder,omitempty"`
	Validation    string       `bson:"validation,omitempty" json:"validation,omitempty"`
	MaxSelections int          `bson:"maxSelections,omitempty" json:"maxSelections,omitempty"`
	Options       []Option     `bson:"options,omitempty" json:"options,omitempty"`
	NextQuestion  NextQuestion `bson:"nextQuestion" json:"nextQuestion"`
}

type Option struct {
	Text  string `bson:"text" json:"text"`
	Icon  string `bson:"icon" json:"icon"`
	Color string `bson:"color" json:"color"`
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
	NewTab bool   `bson:"new_tab" json:"new_tab"`
}
