package service

import (
	"context"

	"github.com/ewrybovic/home-ai-assistant/internal/agents/home"
)

type AgentService struct {
	homeAgent *home.HomeAgent
}

func NewAgentService() (*AgentService, error) {
	var service AgentService

	homeAgent, err := home.New()
	if err != nil {
		return nil, err
	}

	service.homeAgent = homeAgent
	return &service, nil
}

func (a *AgentService) AskAgent(ctx context.Context, question string) (string, error) {
	a.homeAgent.AddHumanMessage(question)
	return a.homeAgent.GenerateString(ctx)
}
