package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ewrybovic/home-ai-assistant/internal/handler"
	"github.com/ewrybovic/home-ai-assistant/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	agentService, err := service.NewAgentService()
	if err != nil {
		log.Fatal(err)
	}
	agentHandler := handler.NewAgentHandler(agentService)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/healthz"))

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/agent", agentHandler.AgentRoutes())
	})

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", r)
}
