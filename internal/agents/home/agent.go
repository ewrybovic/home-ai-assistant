package home

import "github.com/ewrybovic/home-ai-assistant/internal/agents"

type HomeAgent struct {
	*agents.Agent
}

func New() (*HomeAgent, error) {
	agent, err := agents.New("You are a home AI assistant. You must follow the ASD-STE100 simplified techincal english standard")
	if err != nil {
		return nil, err
	}

	return &HomeAgent{
		Agent: agent,
	}, nil
}
