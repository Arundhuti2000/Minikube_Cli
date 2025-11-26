package server

import (
	"context"
	"encoding/json"

	"github.com/Arundhuti2000/Minikube_Cli/internal/tools"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type MinikubeServer struct {
	mcpServer *server.MCPServer
}

func NewMinikubeServer(name, version string) *MinikubeServer {
	// Initialize the MCP server with basic capabilities
	s := server.NewMCPServer(
		name,
		version,
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	return &MinikubeServer{
		mcpServer: s,
	}
}

// RegisterTools registers the minikube tools with the MCP server
func (s *MinikubeServer) RegisterTools() {
	// Register minikube_start tool
	s.mcpServer.AddTool(mcp.NewTool("minikube_start",
		mcp.WithDescription("Starts the Minikube cluster"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		output, err := tools.StartCluster()
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(output), nil
	})

	// Register minikube_stop tool
	s.mcpServer.AddTool(mcp.NewTool("minikube_stop",
		mcp.WithDescription("Stops the Minikube cluster"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		output, err := tools.StopCluster()
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return mcp.NewToolResultText(output), nil
	})

	// Register minikube_status tool
	s.mcpServer.AddTool(mcp.NewTool("minikube_status",
		mcp.WithDescription("Gets the status of the Minikube cluster"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		status, err := tools.GetStatus()
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		// Convert status to JSON string for output
		statusBytes, _ := json.MarshalIndent(status, "", "  ")
		return mcp.NewToolResultText(string(statusBytes)), nil
	})
}

// Start runs the server using Stdio transport
func (s *MinikubeServer) Start() error {
	return server.ServeStdio(s.mcpServer)
}
