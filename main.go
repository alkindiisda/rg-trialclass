package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	goopenai "github.com/CasualCodersProjects/gopenai"
	"github.com/CasualCodersProjects/gopenai/types"
)

func main() {
	godotenv.Load(".env")
	//Whatsapp()
	//Discord()

}

func AIResponse(question string) (response string, err error) {
	godotenv.Load(".env")
	openAI := goopenai.NewOpenAI(&goopenai.OpenAIOpts{
		APIKey: os.Getenv("AI_KEY"),
	})

	request := types.NewDefaultCompletionRequest("The following is a conversation with an AI assistant. The assistant is helpful, creative, clever, and very friendly.\n\nHuman: Hello, who are you?\nAI: I am an AI created by OpenAI. How can I help you today?\nHuman: " + question + "\nAI:")
	request.Model = "text-davinci-003"
	request.Temperature = 0.9
	request.MaxTokens = 150
	request.TopP = 1
	request.FrequencyPenalty = 0
	request.PresencePenalty = 0.6
	request.Stop = []string{" Human:", " AI:"}

	resp, err := openAI.CreateCompletion(request)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("Response not Found!")
	}

	return resp.Choices[0].Text, nil
}

// Reference Discord Go: https://github.com/bwmarrin/discordgo
