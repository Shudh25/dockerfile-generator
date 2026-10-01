package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Message Message `json:"message"`
}

const prompt = `ONLY Generate an ideal Dockerfile for {language} with best practices. Do not provide any description.
Include:
- Base image
- Installing dependencies
- Setting working directory
- Adding source code
- Running the application
- Multistage builds if applicable`

func generateDockerfile(language string) (string, error) {
	// Ollama's default is http://localhost:11434
	// remote host can be set with OLLAMA_HOST environment variable ex. http://192.168.1.68:11434
	host := os.Getenv("OLLAMA_HOST")
	if host == "" {
		host = "http://localhost:11434"
	}

	// Allow OLLAMA_HOST to be set as "localhost:11434"
	if !strings.HasPrefix(host, "http://") &&
		!strings.HasPrefix(host, "https://") {
		host = "http://" + host
	}

	model := os.Getenv("OLLAMA_MODEL")
	if model == "" {
		model = "llama3.1:8b"
	}

	reqBody := ChatRequest{
		Model: model,
		Messages: []Message{
			{
				Role:    "user",
				Content: strings.Replace(prompt, "{language}", language, 1),
			},
		},
		Stream: false,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(
		host+"/api/chat",
		"application/json",
		bytes.NewReader(body),
	)
	if err != nil {
		return "", fmt.Errorf("failed to connect to Ollama: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Ollama returned HTTP %d", resp.StatusCode)
	}

	var result ChatResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode Ollama response: %w", err)
	}

	return result.Message.Content, nil
}

func main() {
	var language string

	fmt.Print("Enter the programming language: ")
	fmt.Scanln(&language)

	if strings.TrimSpace(language) == "" {
		fmt.Println("Language cannot be empty.")
		return
	}

	dockerfile, err := generateDockerfile(language)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nGenerated Dockerfile:\n")
	fmt.Println(dockerfile)
}
