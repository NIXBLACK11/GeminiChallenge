package api

import (
	"context"
	"errors"
	"os"
	"github.com/joho/godotenv"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GetGeminiResponse(Resume string) (string, error) {
	ctx := context.Background()
	// Access your API key as an environment variable
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("error loading .env file")
	}
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return "", err
	}
	defer client.Close()

	// The Gemini 1.5 models are versatile and work with both text-only and multimodal prompts
	model := client.GenerativeModel("gemini-1.5-flash")

	// Creating the input
	input := "This is my resume: " + Resume

	input += "\nbased on these data give me the tags for job search Example: the language, experience so that I can do the google search.Gove just tags and , seperated and give tags like intern or junior mid expert etc based on the resume"

	input += "This is the format I want ->Languages:JavaScript,TypeScript\nFrameworksLibraries:React,Next.js\nDatabases:MongoDB,MySQL\nToolsPlatforms:Git,GitHub\nExperienceLevel:Intern,Junior\nSkills:NLP,CI/CD"

	input += "The output should not have * in it and any other text other than the format"
	resp, err := model.GenerateContent(ctx, genai.Text(input))
	if err != nil {
		return "", err
	}

	if resp != nil {
		candidates := resp.Candidates
		if candidates != nil {
			for _, candidate := range candidates {
				content := candidate.Content
				if content != nil {
					part := content.Parts[0]
					// Type assertion to convert genai.Part to string
					if textPart, ok := part.(genai.Text); ok {
						return string(textPart), nil
					} else {
						return "", errors.New("part is not of type TextPart")
					}
				} else {
					return "", errors.New("content returned nil")
				}
			}
		} else {
			return "", errors.New("candidates is nil")
		}
	} else {
		return "", errors.New("error in response")
	}

	return "", errors.New("internal error")
}
