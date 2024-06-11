package api

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GetGemini() {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// The Gemini 1.5 models are versatile and work with both text-only and multimodal prompts
	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text("Write a story about a magic backpack."))
	if err != nil {
		log.Fatal(err)
	}

	if resp!=nil {
		candidates := resp.Candidates
		if candidates!=nil {
			for _, candidate := range candidates {
				content := candidate.Content
				if content!=nil {
					text := content.Parts[0]
					fmt.Println("Gemini output: ", text)
				} else {
					fmt.Println("content is nil")
				}
			}
		} else {
			fmt.Println("candidates is nil")
		}
	} else {
		fmt.Println("Error in response")
	}
}