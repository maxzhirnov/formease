package models

type QuestionType string

const (
	SingleChoice   QuestionType = "single-choice"
	MultipleChoice QuestionType = "multiple-choice"
	Input          QuestionType = "input"
)

type Question struct {
	ID            int          `json:"id"`
	FormID        int          `json:"form_id"`
	Type          QuestionType `json:"type"`
	Question      string       `json:"question"`
	Subtext       string       `json:"subtext"`
	Image         string       `json:"image"`
	Gradient      string       `json:"gradient"`
	BgColor       string       `json:"bgColor"`
	InputType     string       `json:"inputType,omitempty"`
	Placeholder   string       `json:"placeholder,omitempty"`
	Validation    string       `json:"validation,omitempty"`
	MaxSelections int          `json:"maxSelections,omitempty"`
	Options       []Option     `json:"options,omitempty"`
	NextQuestion  NextQuestion `json:"nextQuestion"`
}

type Option struct {
	Text  string `json:"text"`
	Icon  string `json:"icon"`
	Color string `json:"color"`
}

type NextQuestion struct {
	Conditions []Condition `json:"conditions"`
	Default    int         `json:"default,omitempty"`
}

type Condition struct {
	Answer string `json:"answer"`
	NextID int    `json:"nextId"`
}
