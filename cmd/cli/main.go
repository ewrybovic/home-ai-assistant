package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/ewrybovic/home-ai-assistant/internal/agents"
	"github.com/ewrybovic/home-ai-assistant/internal/agents/home"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langgraphgo/graph"
)

func main() {
	runAgents()
	fmt.Println("Graph:")
	runGraph()
}

func runGraph() {
	agent, err := agents.New("You are a home assitance AI")
	if err != nil {
		log.Fatal(err)
	}

	g := graph.NewMessageGraph()
	g.AddNode("lawn", func(ctx context.Context, state []llms.MessageContent) ([]llms.MessageContent, error) {
		r, err := agent.Agent.GenerateContent(ctx, state, llms.WithTemperature(0.0))
		if err != nil {
			return nil, err
		}

		return append(
			state,
			llms.TextParts(llms.ChatMessageTypeAI, r.Choices[0].Content),
		), nil
	})

	g.AddEdge("lawn", graph.END)
	g.SetEntryPoint("lawn")

	runnable, err := g.Compile()
	if err != nil {
		log.Fatal(err)
	}

	res, err := runnable.Invoke(context.Background(), []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, "When should I mow my lawn"),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range res {
		if r.Role == llms.ChatMessageTypeHuman {
			continue
		}
		for _, part := range r.Parts {
			fmt.Println(part)
		}
	}
}

func runAgents() {
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
