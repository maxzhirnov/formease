package service

import (
	"fmt"

	"github.com/maxzhirnov/formease/pkg/logger"
	"github.com/maxzhirnov/formease/pkg/yandexgpt"
	"go.uber.org/zap"
)

type YandexGPTService struct {
	client *yandexgpt.Client
}

func NewYandexGPTService(apiKey, folderID string) *YandexGPTService {
	return &YandexGPTService{
		client: yandexgpt.NewClient(apiKey, folderID),
	}
}

func (s *YandexGPTService) GetCompletion(prompt string, systemPrompt string) (string, error) {
	response, err := s.client.Complete(yandexgpt.CompletionRequest{
		UserPrompt:   prompt,
		SystemPrompt: systemPrompt,
		Temperature:  0.7, // You might want to adjust this
	})
	if err != nil {
		logger.Error("Error getting completion from YandexGPT",
			zap.Error(err),
			zap.String("prompt", prompt),
			zap.String("systemPrompt", systemPrompt))
		return "", fmt.Errorf("failed to get completion: %w", err)
	}

	// Log the full response for debugging
	logger.Info("YandexGPT raw response", zap.Any("response", response))

	if len(response.Result.Alternatives) == 0 {
		return "", fmt.Errorf("no completion alternatives received")
	}

	completion := response.Result.Alternatives[0].Message.Text
	logger.Info("Completion generated successfully",
		zap.String("completion", completion[:min(len(completion), 100)]+"..."), // Log first 100 chars
		zap.String("role", response.Result.Alternatives[0].Message.Role),
		zap.Any("usage", response.Result.Usage))

	return completion, nil
}

// Helper function for string truncation
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
