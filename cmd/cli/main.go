package main

import (
	"context"
	"log"

	"github.com/ewrybovic/home-ai-assistant/internal/agents/home"
)

func main() {
	homeAgent, err := home.New()
	if err != nil {
		log.Fatal(err)
	}
	homeAgent.AddHumanMessage("How often should I mow my lawn")

	err = homeAgent.Generate(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
