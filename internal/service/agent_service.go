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

func (s *AgentService) AskAgent(ctx context.Context, question string) (<-chan []byte, <-chan error) {
	chunkChan := make(chan []byte)
	errChan := make(chan error)

	s.homeAgent.AddHumanMessage(question)

	go func() {
		err := s.homeAgent.GenerateChan(ctx, chunkChan)
		if err != nil {
			errChan <- err
		}
		close(errChan)
	}()

	return chunkChan, errChan
}
