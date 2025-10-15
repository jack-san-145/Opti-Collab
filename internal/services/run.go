package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"opti-collab/models"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func AnalyzeCode(code string, language string) (*models.CodeAnalysisResponse, error) {
	// Load .env if not already loaded
	_ = godotenv.Load()

	prompt := fmt.Sprintf(`You are a code optimization and analysis AI.

	Analyze the following user code step-by-step. Your tasks are:
	1. **Identify redundant, repeated, or unnecessary code blocks, variables, or functions** directly from the user’s original code — do not infer them from the optimized code.
	2. **List all unused variables** that are declared but never used in the code.
	3. **List all unused functions** that are defined but never called.
	4. **Explain redundancy** by showing the exact redundant lines or blocks from the original code.
	5. **Generate a fully optimized version** of the code with all redundant logic removed while preserving correct functionality and indentation.
	6. **Calculate performance metrics** such as CPU time and memory usage for the user's original code.
	7. **Estimate the code optimization level (0–100)** based on the efficiency and cleanliness of the user's original code.

	Input (JSON):
	{
	"language": "%s",
	"code": %q
	}

	Output must be **only** this JSON structure:

	{
	"code_optimization_level": <integer 0–100>,
	"cpu_performance": "<number>ms",
	"memory_usage": "<number>kb",
	"error": <string|null>,
	"output": "<string>",
	"redundant_block": "<string|null>",
	"unused_variables": "<string|null>",
	"unused_functions": "<string|null>",
	"suggested_optimized_code": "<string>"
	}

	Rules:
	- Always return **valid JSON** and nothing else.
	- Maintain **exact indentation and formatting** in redundant_block, unused_variables, unused_functions, and suggested_optimized_code.
	- If no redundant code exists, redundant_block = null.
	- If no unused variables exist, unused_variables = null.
	- If no unused functions exist, unused_functions = null.
	- If code runs successfully, error = null.
	- Never infer redundant_block, unused_variables, or unused_functions from optimized code; detect them only from the **original user code**.
	- Respond ONLY with JSON and nothing outside it.
	`, language, code)

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("missing GEMINI_API_KEY")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	model := client.GenerativeModel("models/gemini-2.5-flash")

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("no response from Gemini")
	}

	// Convert genai.Part to string
	text := fmt.Sprint(resp.Candidates[0].Content.Parts[0])

	// Strip ```json or ``` wrapping if present
	text = strings.TrimSpace(text)
	if strings.HasPrefix(text, "```") {
		lines := strings.SplitN(text, "\n", 2)
		if len(lines) == 2 {
			text = lines[1]
		} else {
			text = ""
		}
	}
	text = strings.TrimSuffix(text, "```")
	text = strings.TrimSpace(text)

	// Parse the JSON from Gemini
	var analysis models.CodeAnalysisResponse
	if err := json.Unmarshal([]byte(text), &analysis); err != nil {
		return nil, fmt.Errorf("failed to parse response as JSON: %v\nRaw response: %s", err, text)
	}

	fmt.Println("analysis - ", analysis)
	return &analysis, nil
}
