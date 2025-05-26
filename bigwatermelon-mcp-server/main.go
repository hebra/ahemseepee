package main

import (
	"encoding/json"
	"github.com/hebra/ahemseepee/bigwatermelon-mcp-server/internal"
	"log/slog"
	"net/http"
	"os"
)

var log = slog.New(slog.NewTextHandler(os.Stderr, nil))

type MCPServer struct {
	handlers map[string]func(params json.RawMessage) (interface{}, error)
}

func NewMCPServer() *MCPServer {
	return &MCPServer{
		handlers: make(map[string]func(params json.RawMessage) (interface{}, error)),
	}
}

func (s *MCPServer) RegisterHandler(action string, handler func(params json.RawMessage) (interface{}, error)) {
	s.handlers[action] = handler
}

func (s *MCPServer) HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req internal.MCPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	handler, exists := s.handlers[req.Action]
	if !exists {
		sendResponse(w, internal.MCPResponse{
			Status:    "error",
			Error:     "Unknown action",
			RequestID: req.RequestID,
		})
		return
	}

	result, err := handler(req.Parameters)
	if err != nil {
		sendResponse(w, internal.MCPResponse{
			Status:    "error",
			Error:     err.Error(),
			RequestID: req.RequestID,
		})
		return
	}

	sendResponse(w, internal.MCPResponse{
		Status:    "success",
		Data:      result,
		RequestID: req.RequestID,
	})
}

func sendResponse(w http.ResponseWriter, resp internal.MCPResponse) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Error("Error sending response: ", "Error", err)
		return
	}
}

func main() {
	log.Info("Starting offers extractor...")

	server := NewMCPServer()

	server.RegisterHandler("get-big-watermelon-deals", func(params json.RawMessage) (interface{}, error) {
		return map[string]internal.ResponseData{"message": internal.FetchOffers()}, nil
	})

	http.HandleFunc("/mcp", server.HandleRequest)
	log.Info("Starting MCP server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Error("Error starting MCP server: ", "Error", err)
	}
}
