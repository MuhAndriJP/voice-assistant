package openai

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

func OpenAI(userInput string) (res string, err error) {
	apiKey := os.Getenv("OPENAI_KEY")

	response, err := sendChatRequest(apiKey, userInput)
	if err != nil {
		return
	}

	res = response.Choices[0].Message.Content
	return
}

func sendChatRequest(apiKey, userInput string) (*OpenAIResponse, error) {
	url := os.Getenv("OPENAI_URL")

	requestBody := OpenAIRequest{
		Model: ModelAI,
		Messages: []Message{
			{
				Role:    User,
				Content: userInput,
			},
		},
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(requestJSON)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response OpenAIResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
