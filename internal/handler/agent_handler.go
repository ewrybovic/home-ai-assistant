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

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.WriteHeader(http.StatusOK)

	chunkChan, errChan := a.agentService.AskAgent(r.Context(), req.Question)
	for {
		select {
		case <-r.Context().Done():
			return
		case err := <-errChan:
			if err != nil {
				w.Write([]byte("Error generating response: " + err.Error()))
			}
			return
		case chunk, open := <-chunkChan:
			if !open {
				return
			}

			w.Write(chunk)
			flusher.Flush()
		}
	}
}
