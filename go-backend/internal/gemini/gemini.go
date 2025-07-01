package gemini

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

// CallGemini uses the official Gemini Go SDK to send the prompt and return the result.
// The API key is read from the GOOGLE_API_KEY environment variable.
// For backward compatibility, we keep the apiKey parameter but it's not used.
func CallGemini(prompt, apiKey string) (string, error) {
	// Note: apiKey parameter is ignored as the SDK reads from GOOGLE_API_KEY env var

	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create Gemini client: %w", err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash", // Using the standard pro model available in free tier
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("gemini API error: %w", err)
	}

	// Extract the text from the response
	return result.Text(), nil
}
