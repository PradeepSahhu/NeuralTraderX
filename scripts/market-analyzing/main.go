package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {

	llm, err := openai.New(
		openai.WithToken(os.Getenv("DEEPSEEK_API_KEY")),
		openai.WithModel("deepseek-v4-flash"),
		openai.WithBaseURL("https://api.deepseek.com/v1"),
	)

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	response, err := llms.GenerateFromSinglePrompt(ctx, llm, "How are you today? and are you deepseek or openai chatgpt?")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}
