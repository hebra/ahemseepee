package main

import (
	"context"
	"encoding/json"
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/ThinkInAIXYZ/go-mcp/transport"
	"github.com/gin-gonic/gin"
	"github.com/hebra/ahemseepee/bigwatermelon-mcp-server/internal"
	"log/slog"
	"os"
)

var log = slog.New(slog.NewTextHandler(os.Stderr, nil))

type DealsRequest struct {
}

func main() {
	log.Info("Starting offers extractor...")

	messageEndpointURL := "/message"

	sseTransport, mcpHandler, err := transport.NewSSEServerTransportAndHandler(messageEndpointURL)
	if err != nil {
		log.Error("Error creating SEE transport and handler.", "Error", err)
		os.Exit(1)
	}

	mcpServer, _ := server.NewServer(sseTransport)

	tool, err := protocol.NewTool("get-big-watermelon-deals",
		"Get today's deals from Big Watermelon",
		DealsRequest{})

	if err != nil {
		log.Error("Failed to create tool.", "Error", err)
		os.Exit(1)
	}
	mcpServer.RegisterTool(tool, getDailyDealsHandler)

	go func() {
		err := mcpServer.Run()
		if err != nil {
			log.Error("Failed to start MCP server.", "Error", err)
			os.Exit(1)
		}
	}()

	defer func(mcpServer *server.Server, userCtx context.Context) {
		err := mcpServer.Shutdown(userCtx)
		if err != nil {
			log.Error("Failed to shutdown MCP server.", "Error", err)
			os.Exit(1)
		}
	}(mcpServer, context.Background())

	r := gin.Default()
	r.GET("/sse", func(ctx *gin.Context) {
		mcpHandler.HandleSSE().ServeHTTP(ctx.Writer, ctx.Request)
	})
	r.POST(messageEndpointURL, func(ctx *gin.Context) {
		mcpHandler.HandleMessage().ServeHTTP(ctx.Writer, ctx.Request)
	})

	if err = r.Run(":8080"); err != nil {
		log.Error("Failed to start HTTP server.", "Error", err)
		os.Exit(1)
	}
}

func getDailyDealsHandler(ctx context.Context, req *protocol.CallToolRequest) (*protocol.CallToolResult, error) {

	bytes, err := json.Marshal(internal.FetchOffers())
	if err != nil {
		log.Error("Error marshalling JSON.", "Error", err)
		return nil, err
	}

	return &protocol.CallToolResult{
		Content: []protocol.Content{
			&protocol.TextContent{
				Type: "text",
				Text: string(bytes),
			},
		},
	}, nil

}
