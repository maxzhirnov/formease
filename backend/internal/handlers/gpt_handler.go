package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/maxzhirnov/formease/internal/models"
	"github.com/maxzhirnov/formease/internal/service"
	"github.com/maxzhirnov/formease/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type GPTHandler struct {
	formService      *service.FormService
	yandexGPTService *service.YandexGPTService
}

type GenerateFormRequest struct {
	Topic        string   `json:"topic" binding:"required"`
	FormType     string   `json:"formType" binding:"required"`
	NumQuestions int      `json:"numQuestions" binding:"required,min=1,max=10"`
	Preferences  []string `json:"preferences,omitempty"`
}

func NewGPTHandler(formService *service.FormService, yandexGPTService *service.YandexGPTService) *GPTHandler {
	return &GPTHandler{
		formService:      formService,
		yandexGPTService: yandexGPTService,
	}
}

func (h *GPTHandler) GenerateForm(c *gin.Context) {
	logger.Info("Starting form generation")

	// Parse request
	var req GenerateFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request data", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Get user ID from context
	userID, exists := c.Get("userID")
	if !exists {
		logger.Error("UserID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Create prompt for GPT
	prompt := h.createPrompt(req)
	systemPrompt := `You are a form generation assistant. Generate a form in JSON format following these rules:
    1. Follow the exact JSON structure as provided in the example
    2. Make sure all IDs are sequential and valid
    3. Include appropriate validation rules for input fields
    4. Create logical next question conditions
    5. Return only valid JSON without any additional text
	question types can be: 'input', 'single-choice', 'multiple-choice', 'rating'
	theme can be: 'tech', 'flat', 'dark', 'light'
	floatingShapesTheme can be: 'spring', 'summer', 'autumn', 'winter'
	validation for input fields can be: 'email', 'phone', 'text'
	Here's example JSON structure:
	{
    "name": "Sample Form",
    "theme": "tech",
    "floatingShapesTheme": "spring",
    "questions": [
      {
        "id": 1,
        "type": "single-choice",
        "question": "What brings you here today?",
        "subtext": "Help us personalize your experience",
        "image": "/img/demo.jpg",
        "options": [
          { "text": "Building a Website", "icon": "🎨" },
          { "text": "Mobile App", "icon": "📱" },
          { "text": "Just Exploring", "icon": "🔍" }
        ],
        "nextQuestion": {
          "conditions": [
            { "answer": "Building a Website", "nextId": 2 },
            { "answer": "Mobile App", "nextId": 3 },
            { "answer": "Just Exploring", "nextId": 4 }
          ]
        }
      },
      {
          "id": 2,
          "type": "multiple-choice",
          "question": "What features do you need?",
          "subtext": "Select up to 4 key features",
          "image": "/img/demo.jpg",
          "maxSelections": 4,
          "options": [
            { "text": "E-commerce", "icon": "🛍️" },
            { "text": "Blog", "icon": "✍️" },
            { "text": "Authentication", "icon": "🔒" },
            { "text": "Analytics", "icon": "📊" },
            { "text": "SEO Tools", "icon": "🎯" }
          ],
          "nextQuestion": {
            "conditions": [],
            "default": 5
          }
        },
        {
          "id": 3,
          "type": "multiple-choice",
          "question": "Which platforms are you targeting?",
          "subtext": "Select all that apply",
          "image": "/img/demo.jpg",
          "maxSelections": 3,
          "options": [
            { "text": "iOS", "icon": "🍎" },
            { "text": "Android", "icon": "🤖" },
            { "text": "Both", "icon": "📱" }
          ],
          "nextQuestion": {
            "conditions": [],
            "default": 5
          }
        },
        {
          "id": 4,
          "type": "input",
          "question": "What's your main interest?",
          "subtext": "Tell us what you'd like to learn more about",
          "image": "/img/demo.jpg",
          "inputType": "text",
          "placeholder": "Enter your interest",
          "validation": "text",
          "nextQuestion": {
            "conditions": [],
            "default": 5
          }
        },
        {
          "id": 5,
          "type": "input",
          "question": "What's your email?",
          "subtext": "We'll send you personalized recommendations",
          "image": "/img/demo.jpg",
          "inputType": "email",
          "placeholder": "your@email.com",
          "validation": "email",
          "nextQuestion": {
            "conditions": [],
            "default": 6
          }
        },
        {
			"id": 6,
			"type": "rating",
			"question": "Оцените нашу работу?",
			"subtext": "",
			"image": "",
			"minValue": 1,
			"maxValue": 5,
			"step": 1,
			"showLabels": true,
			"minLabel": "Плохо",
			"maxLabel": "Отлично",
			"icon": "⭐️",
			"nextQuestion": {
				"conditions": []
			}
		}
    ],
    "thankYouMessage": {
      "title": "Спасибо вам огромное!",
      "subtitle": "За то что прошли нашу форму!.",
      "icon": "✨",
      "button": {
        "text": "Перейти",
        "url": "https://hemaks.org",
        "newTab": true
      }
    }
  }
	`

	// Call YandexGPT
	response, err := h.yandexGPTService.GetCompletion(prompt, systemPrompt)
	if err != nil {
		logger.Error("Failed to generate form with GPT", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate form"})
		return
	}

	// Clean up the response - remove markdown code blocks
	response = cleanJSONResponse(response)

	logger.Info("Cleaned JSON response", zap.String("json", response))

	// Parse the generated form
	var generatedForm models.Form
	if err := json.Unmarshal([]byte(response), &generatedForm); err != nil {
		logger.Error("Failed to parse generated form", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid generated form format"})
		return
	}

	// Validate the generated form
	if err := h.validateAndFixForm(&generatedForm); err != nil {
		logger.Error("Generated form validation failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set user ID and draft status
	objectID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		logger.Error("Invalid user ID format", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}
	generatedForm.UserID = objectID
	generatedForm.IsDraft = true

	// Save the form
	if err := h.formService.CreateForm(&generatedForm); err != nil {
		logger.Error("Failed to save generated form", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save form"})
		return
	}

	logger.Info("Form generated and saved successfully", zap.String("formId", generatedForm.ID.Hex()))
	c.JSON(http.StatusCreated, generatedForm)
}

func (h *GPTHandler) createPrompt(req GenerateFormRequest) string {
	return fmt.Sprintf(`Generate a form with the following specifications:
    - Topic: %s
    - Form Type: %s
    - Number of Questions: %d
    - Preferences: %v
    
    The form should include:
    1. A meaningful name and theme
    2. Logical question flow with proper navigation
    3. Appropriate input validation
    4. A thank you message
    5. Proper icons and visual elements
    
    Generate the complete form JSON structure following the provided format.`,
		req.Topic, req.FormType, req.NumQuestions, req.Preferences)
}
func (h *GPTHandler) validateAndFixForm(form *models.Form) error {
	// Critical validations that can't have defaults
	if form.Name == "" {
		return fmt.Errorf("form name is required")
	}

	if len(form.Questions) == 0 {
		return fmt.Errorf("form must contain at least one question")
	}

	// Set default theme if empty
	if form.Theme == "" {
		form.Theme = "default"
	}
	if form.FloatingShapes == "" {
		form.FloatingShapes = "default"
	}

	// Fix and validate questions
	questionIDs := make(map[int]bool)
	for i := range form.Questions {
		// Fix question ID if duplicate or invalid
		if questionIDs[form.Questions[i].ID] || form.Questions[i].ID <= 0 {
			form.Questions[i].ID = len(questionIDs) + 1
		}
		questionIDs[form.Questions[i].ID] = true

		// Validate and fix question type
		if err := h.validateAndFixQuestionType(&form.Questions[i]); err != nil {
			return err
		}

		// Fix next question logic
		h.validateAndFixNextQuestion(&form.Questions[i], len(form.Questions))
	}

	// Fix thank you message
	h.validateAndFixThankYouMessage(&form.ThankYouMessage)

	return nil
}

func (h *GPTHandler) validateAndFixQuestionType(q *models.Question) error {
	// Set default image if not provided or invalid
	if q.Image == "" || q.Image == "/img/demo.jpg" {
		q.Image = fmt.Sprintf("/img/default/%s.jpg", q.Type)
	}

	switch q.Type {
	case "single-choice", "multiple-choice":
		if len(q.Options) == 0 {
			// Add default options if none provided
			q.Options = []models.Option{
				{Text: "Option 1", Icon: "✨"},
				{Text: "Option 2", Icon: "🌟"},
			}
		}
		// Set default maxSelections for multiple-choice
		if q.Type == "multiple-choice" && q.MaxSelections <= 0 {
			q.MaxSelections = len(q.Options)
		}

	case "input":
		if q.InputType == "" {
			q.InputType = "text"
		}
		if q.Validation == "" {
			q.Validation = "/.+/"
		}
		if q.Placeholder == "" {
			q.Placeholder = "Enter your answer here"
		}

	default:
		// Set default type if invalid
		q.Type = "single-choice"
		q.Options = []models.Option{
			{Text: "Option 1", Icon: "✨"},
			{Text: "Option 2", Icon: "🌟"},
		}
	}

	// Set default subtext if empty
	if q.Subtext == "" {
		q.Subtext = "Please provide your answer"
	}

	return nil
}

func (h *GPTHandler) validateAndFixNextQuestion(q *models.Question, totalQuestions int) {
	// Fix invalid next question IDs in conditions
	validConditions := make([]models.Condition, 0)
	for _, condition := range q.NextQuestion.Conditions {
		if condition.NextID > 0 && condition.NextID <= totalQuestions {
			validConditions = append(validConditions, condition)
		}
	}
	q.NextQuestion.Conditions = validConditions

	// Fix default next question
	if q.NextQuestion.Default <= 0 || q.NextQuestion.Default > totalQuestions {
		if q.ID < totalQuestions {
			q.NextQuestion.Default = q.ID + 1
		} else {
			q.NextQuestion.Default = 0 // Last question
		}
	}
}

func (h *GPTHandler) validateAndFixThankYouMessage(msg *models.ThankYouMessage) {
	if msg.Title == "" {
		msg.Title = "Thank You!"
	}
	if msg.Subtitle == "" {
		msg.Subtitle = "We appreciate your feedback"
	}
	if msg.Icon == "" {
		msg.Icon = "✨"
	}
	if msg.Button.Text == "" {
		msg.Button.Text = "Continue"
	}
	if msg.Button.URL == "" {
		msg.Button.URL = "/"
	}
}

func cleanJSONResponse(jsonStr string) string {
	// First remove markdown code blocks if present
	jsonStr = strings.TrimSpace(jsonStr)
	jsonStr = strings.TrimPrefix(jsonStr, "```")
	jsonStr = strings.TrimSuffix(jsonStr, "```")
	jsonStr = strings.TrimPrefix(jsonStr, "json")
	jsonStr = strings.TrimSpace(jsonStr)

	// Fix escaped regex patterns
	// Replace \d with \\d
	jsonStr = strings.ReplaceAll(jsonStr, `\d`, `\\d`)
	// Replace other common regex escapes
	jsonStr = strings.ReplaceAll(jsonStr, `\.`, `\\.`)
	jsonStr = strings.ReplaceAll(jsonStr, `\+`, `\\+`)
	jsonStr = strings.ReplaceAll(jsonStr, `\*`, `\\*`)
	jsonStr = strings.ReplaceAll(jsonStr, `\?`, `\\?`)
	jsonStr = strings.ReplaceAll(jsonStr, `\[`, `\\[`)
	jsonStr = strings.ReplaceAll(jsonStr, `\]`, `\\]`)
	jsonStr = strings.ReplaceAll(jsonStr, `\(`, `\\(`)
	jsonStr = strings.ReplaceAll(jsonStr, `\)`, `\\)`)

	return jsonStr
}
