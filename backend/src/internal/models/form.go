package models

type Form struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	Theme           string          `json:"theme"`
	FloatingShapes  string          `json:"floatingShapes"`
	Questions       []Question      `json:"questions"`
	ThankYouMessage ThankYouMessage `json:"thankYouMessage"`
}

type ThankYouMessage struct {
	Title    string       `json:"title"`
	Subtitle string       `json:"subtitle"`
	Icon     string       `json:"icon"`
	Button   ButtonConfig `json:"button"`
}

type ButtonConfig struct {
	Text   string `json:"text"`
	URL    string `json:"url"`
	NewTab bool   `json:"newTab"`
}
