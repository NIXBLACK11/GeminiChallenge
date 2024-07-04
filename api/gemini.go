package api

import (
	"context"
	"errors"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GetGemini(input string) (string, error){
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("Error loading .env file")
	}
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return "", err
	}
	defer client.Close()

	// The Gemini 1.5 models are versatile and work with both text-only and multimodal prompts
	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(input))
	if err != nil {
		return "", err
	}

	if resp!=nil {
		candidates := resp.Candidates
		if candidates!=nil {
			for _, candidate := range candidates {
				content := candidate.Content
				if content!=nil {
					text := content.Parts[0]
					fmt.Println(text)
					// return string(text), nil
				} else {
					return "", errors.New("content returned nil")
				}
			}
		} else {
			return "", errors.New("candidates is nil")
		}
	} else {
		return "", errors.New("Error in response")
	}

	return "", errors.New("Internal error")
}