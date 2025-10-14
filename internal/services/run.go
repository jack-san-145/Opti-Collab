package services

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func AnalyzeCode(code string, language string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("missing GEMINI_API_KEY")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("models/gemini-2.5-flash")

	prompt := fmt.Sprintf("You are a coding assistant. Analyze, optimize, and explain the following %s code:\n\n%s", language, code)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	// Extract the text response
	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		text := resp.Candidates[0].Content.Parts[0]
		fmt.Println("gemini text - ", text)
		return fmt.Sprintf("%v", text), nil
	}

	return "No response", nil
}
