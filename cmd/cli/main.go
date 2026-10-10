package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/ewrybovic/home-ai-assistant/internal/agents/home"
)

func main() {
	var agents []*home.HomeAgent
	homeAgent1, err := home.New()
	if err != nil {
		log.Fatal(err)
	}
	homeAgent1.AddHumanMessage("How often should I mow my lawn?")
	agents = append(agents, homeAgent1)

	homeAgent2, err := home.New()
	if err != nil {
		log.Fatal(err)
	}
	homeAgent2.AddHumanMessage("When do I need to change the air filter in my HVAC?")
	agents = append(agents, homeAgent2)

	var wg sync.WaitGroup
	for _, agent := range agents {
		wg.Go(func() {
			output, err := agent.GenerateString(context.Background())
			if err != nil {
				// Not correct to do this, testing it out
				log.Fatal(err)
			}

			fmt.Printf("%s\n--- Output ---\n%s\n", agent.SystemPrompt, output)
		})
	}
	wg.Wait()
}
