package llm

import (
	"context"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms/openai"
)

func CreateLLM() (*openai.LLM, context.Context) {

	llm, err := openai.New(
		openai.WithToken(os.Getenv("DEEPSEEK_API_KEY")),
		openai.WithModel("deepseek-v4-flash"),
		openai.WithBaseURL("https://api.deepseek.com/v1"),
	)

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	return llm, ctx
	// response, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(response)
}
