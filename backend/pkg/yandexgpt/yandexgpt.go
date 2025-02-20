// Package yandexgpt provides functionality to interact with Yandex GPT API
package yandexgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// DefaultBaseURL is the base URL for Yandex GPT API
	DefaultBaseURL = "https://llm.api.cloud.yandex.net/foundationModels/v1/completion"

	// DefaultModel is the default model URI
	DefaultModel = "gpt://b1gakioh5lutqcssd8ph/yandexgpt"

	// DefaultSystemPrompt is the default system prompt
	DefaultSystemPrompt = "Provide user with what he requested"
)

// Client represents a Yandex GPT API client
type Client struct {
	baseURL    string
	apiKey     string
	folderID   string
	httpClient *http.Client
}

// Message represents a chat message
type Message struct {
	Role string `json:"role"`
	Text string `json:"text"`
}

// CompletionOptions represents completion parameters
type CompletionOptions struct {
	Stream      bool    `json:"stream"`
	Temperature float64 `json:"temperature"`
	MaxTokens   int     `json:"maxTokens,omitempty"`
}

// RequestData represents the request payload
type RequestData struct {
	ModelURI          string            `json:"modelUri"`
	CompletionOptions CompletionOptions `json:"completionOptions"`
	Messages          []Message         `json:"messages"`
}

// Response represents the API response
type Response struct {
	Result struct {
		Alternatives []struct {
			Message struct {
				Role string `json:"role"`
				Text string `json:"text"`
			} `json:"message"`
			Status string `json:"status"`
		} `json:"alternatives"`
		Usage struct {
			InputTextTokens  json.Number `json:"inputTextTokens"`
			OutputTextTokens json.Number `json:"outputTextTokens"`
			TotalTokens      json.Number `json:"totalTokens"`
		} `json:"usage"`
	} `json:"result"`
}

// NewClient creates a new Yandex GPT client
func NewClient(apiKey, folderID string) *Client {
	return &Client{
		baseURL:  DefaultBaseURL,
		apiKey:   apiKey,
		folderID: folderID,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// CompletionRequest represents the parameters for a completion request
type CompletionRequest struct {
	SystemPrompt string
	UserPrompt   string
	Model        string
	Temperature  float64
	MaxTokens    int
}

// Complete sends a completion request to the API
func (c *Client) Complete(req CompletionRequest) (*Response, error) {
	// Set default values if not provided
	if req.Model == "" {
		req.Model = DefaultModel
	}
	if req.SystemPrompt == "" {
		req.SystemPrompt = DefaultSystemPrompt
	}
	if req.Temperature == 0 {
		req.Temperature = 0.8
	}

	data := RequestData{
		ModelURI: req.Model,
		CompletionOptions: CompletionOptions{
			Stream:      false,
			Temperature: req.Temperature,
			MaxTokens:   req.MaxTokens,
		},
		Messages: []Message{
			{
				Role: "system",
				Text: req.SystemPrompt,
			},
			{
				Role: "user",
				Text: req.UserPrompt,
			},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	request, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Api-Key "+c.apiKey)
	request.Header.Set("x-folder-id", c.folderID)

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return &response, nil
}
