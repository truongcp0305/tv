package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"tuvi/models"
)

const defaultLMStudioEndpoint = "http://127.0.0.1:1234/v1/chat/completions"
const defaultLMStudioModel = "local-model"
const defaultLMStudioTemperature = 0.6
const defaultLMStudioTimeoutSec = 600

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatCompletionRequest struct {
	Model       string        `json:"model"`
	Messages    []chatMessage `json:"messages"`
	Temperature float64       `json:"temperature"`
}

type chatCompletionResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type LMStudioClient struct {
	Endpoint    string
	Model       string
	Temperature float64
	TimeoutSec  int
}

func NewLMStudioClient() LMStudioClient {
	return LMStudioClient{
		Endpoint:    getEnvOrDefault("LMSTUDIO_ENDPOINT", defaultLMStudioEndpoint),
		Model:       getEnvOrDefault("LMSTUDIO_MODEL", defaultLMStudioModel),
		Temperature: defaultLMStudioTemperature,
		TimeoutSec:  getEnvIntOrDefault("LMSTUDIO_TIMEOUT_SEC", defaultLMStudioTimeoutSec),
	}
}

func (c LMStudioClient) chat(messages []chatMessage) (string, error) {
	endpoint := strings.TrimSpace(c.Endpoint)
	if endpoint == "" {
		endpoint = defaultLMStudioEndpoint
	}

	model := strings.TrimSpace(c.Model)
	if model == "" {
		model = defaultLMStudioModel
	}

	temperature := c.Temperature
	if temperature == 0 {
		temperature = defaultLMStudioTemperature
	}

	timeoutSec := c.TimeoutSec
	if timeoutSec <= 0 {
		timeoutSec = defaultLMStudioTimeoutSec
	}

	payload := chatCompletionRequest{
		Model:       model,
		Messages:    messages,
		Temperature: temperature,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{Timeout: time.Duration(timeoutSec) * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("goi LM Studio that bai: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("LM Studio tra ve status khong hop le: %d", resp.StatusCode)
	}

	var result chatCompletionResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("parse LM Studio response: %w", err)
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("LM Studio khong tra ve choices hop le")
	}

	content := strings.TrimSpace(result.Choices[0].Message.Content)
	if content == "" {
		return "", fmt.Errorf("noi dung tra ve rong")
	}

	return content, nil
}

func (c LMStudioClient) AnalyzeDetailPlace(mainId int, gender int, page models.HoroscopePage) (string, error) {
	userPrompt := BuildMessageDetailPlace(mainId, gender, page)
	if strings.TrimSpace(userPrompt) == "" {
		return "", fmt.Errorf("khong build duoc user prompt")
	}

	return c.chat([]chatMessage{
		{Role: "system", Content: SYSTEM_PORMT},
		{Role: "user", Content: userPrompt},
	})
}

func (c LMStudioClient) AnalyzeAppearance(gender int, page models.HoroscopePage) (string, error) {
	userPrompt := BuildMessageAppearance(gender, page)
	if strings.TrimSpace(userPrompt) == "" {
		return "", fmt.Errorf("khong build duoc user prompt")
	}

	return c.chat([]chatMessage{
		{Role: "system", Content: SYSTEM_PORMT},
		{Role: "user", Content: userPrompt},
	})
}

func getEnvOrDefault(key string, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func getEnvIntOrDefault(key string, fallback int) int {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(raw)
	if err != nil {
		return fallback
	}
	return parsed
}
