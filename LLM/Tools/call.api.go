package tools

import (
	"context"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/tools"
)

type CustomTool struct{}

func (c CustomTool) Name() string {
	return "custom_tool"
}

func (c CustomTool) Description() string {
	return "Performs custom operations"
}

// the agent will call this method for the tool execution
func (c CustomTool) Call(ctx context.Context, input string) (string, error) {
	// Your custom logic here
	return "Tool result", nil
}

// Link : https://github.com/tmc/langchaingo/blob/main/docs/docs/modules/agents/index.mdx

func NewsCall(llm *openai.LLM, ctx context.Context) {
	api_key := "REMOVED_API_KEY"

	// Create tools
	tools := []tools.Tool{
		tools.Calculator{},
		tools.WebSearch{},
	}

	// Create agent
	agent := agents.NewOneShotAgent(llm, tools)

	// Create executor
	executor := agents.NewExecutor(agent)

	result, err := executor.Call(ctx, map[string]any{
		"call the api to get the news ",
	})
}
