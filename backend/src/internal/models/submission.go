package models

import "time"

type Submission struct {
	ID          int       `json:"id"`
	FormID      int       `json:"formId"`
	Answers     []Answer  `json:"answers"`
	SubmittedAt time.Time `json:"submittedAt"`
}

type Answer struct {
	QuestionID int    `json:"questionId"`
	Value      string `json:"value"`
}
