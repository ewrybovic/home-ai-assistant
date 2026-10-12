package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ewrybovic/home-ai-assistant/internal/service"
	"github.com/go-chi/chi/v5"
)

type AgentHandler struct {
	agentService *service.AgentService
}

type AskAgentRequest struct {
	Question string `json:"question"`
}

func NewAgentHandler(agentService *service.AgentService) *AgentHandler {
	return &AgentHandler{
		agentService: agentService,
	}
}

func (a *AgentHandler) AgentRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", a.askAgentHandler)

	return r
}

func (a *AgentHandler) askAgentHandler(w http.ResponseWriter, r *http.Request) {
	var req AskAgentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	answer, err := a.agentService.AskAgent(r.Context(), req.Question)
	if err != nil {
		http.Error(w, "Error generating answer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(answer))
}
