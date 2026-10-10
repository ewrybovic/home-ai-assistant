package agents

import (
	"context"
	"fmt"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type Agent struct {
	SystemPrompt string
	Agent        *ollama.LLM
	Messages     []llms.MessageContent
}

func New(systemPrompt string) (*Agent, error) {
	gemma4, err := ollama.New(ollama.WithModel("gemma4:e4b"))
	if err != nil {
		return nil, err
	}

	agent := &Agent{
		SystemPrompt: systemPrompt,
		Agent:        gemma4,
	}
	agent.AddSystemMessage(agent.SystemPrompt)

	return agent, nil
}

func (a *Agent) AddHumanMessage(message string) {
	a.Messages = append(a.Messages, llms.TextParts(llms.ChatMessageTypeHuman, message))
}

func (a *Agent) AddSystemMessage(message string) {
	a.Messages = append(a.Messages, llms.TextParts(llms.ChatMessageTypeSystem, message))
}

func (a *Agent) Generate(ctx context.Context) error {
	_, err := a.Agent.GenerateContent(ctx, a.Messages, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))
		return nil
	}))
	fmt.Println()
	return err
}

func (a *Agent) GenerateString(ctx context.Context) (string, error) {
	var sb strings.Builder
	_, err := a.Agent.GenerateContent(ctx, a.Messages, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		sb.WriteString(string(chunk))
		return nil
	}))
	return sb.String(), err
}
