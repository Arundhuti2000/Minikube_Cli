# Step-by-Step Build Guide 🛠️

This guide will walk you through building the Minikube MCP server from scratch. Follow each step in order.

---

## 🎯 Stage 1: Project Setup

### Step 1: Verify Prerequisites

Before starting, make sure everything is installed:

```bash
# Check Go installation
go version
# Should output: go version go1.21.x or higher

# Check Minikube
minikube version
# Should output: minikube version: vX.X.X

# Check kubectl
kubectl version --client
# Should output kubectl client version
```

If anything is missing, install it using the links in the README.

### Step 2: Create Project Structure

```bash
# Create main project directory (you're already here!)
mkdir -p cmd/minikube-mcp
mkdir -p internal/cmd
mkdir -p internal/server
mkdir -p internal/tools
mkdir -p internal/executor
mkdir -p pkg/minikube
```

### Step 3: Initialize Go Module

If you haven't already:

```bash
go mod init github.com/yourusername/minikube-mcp
```

### Step 4: Install Dependencies

```bash
# Install MCP SDK
go get github.com/mark3labs/mcp-go@latest

# Install Cobra for CLI
go get github.com/spf13/cobra@latest

# Update go.mod and go.sum
go mod tidy
```

**✅ Checkpoint:** Run `ls` and verify you see the directory structure. Run `cat go.mod` and verify dependencies are listed.

---

## 🎯 Stage 2: Hello World CLI

### Step 5: Create the Main Entry Point

Create `cmd/minikube-mcp/main.go`:

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello from Minikube MCP!")
	os.Exit(0)
}
```

### Step 6: Test Your First Build

```bash
# Build it
go build -o bin/minikube-mcp ./cmd/minikube-mcp

# Run it
./bin/minikube-mcp
```

**Expected output:** `Hello from Minikube MCP!`

**✅ Checkpoint:** If you see the message, great! You've built your first Go binary.

### Step 7: Set Up Cobra CLI

Now let's add proper CLI structure. Create `internal/cmd/root.go`:

```go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	transport  string
	logLevel   string
	accessLevel string
)

var rootCmd = &cobra.Command{
	Use:   "minikube-mcp",
	Short: "Minikube MCP Server",
	Long:  `A Model Context Protocol server for Minikube Kubernetes clusters.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Minikube MCP Server starting...")
		fmt.Printf("Transport: %s\n", transport)
		fmt.Printf("Log Level: %s\n", logLevel)
		fmt.Printf("Access Level: %s\n", accessLevel)
	},
}

func init() {
	rootCmd.Flags().StringVar(&transport, "transport", "stdio", "Transport mechanism (stdio, sse, http)")
	rootCmd.Flags().StringVar(&logLevel, "log-level", "info", "Log level (debug, info, warn, error)")
	rootCmd.Flags().StringVar(&accessLevel, "access-level", "readonly", "Access level (readonly, readwrite)")
}

func Execute() error {
	return rootCmd.Execute()
}
```

### Step 8: Update Main to Use Cobra

Update `cmd/minikube-mcp/main.go`:

```go
package main

import (
	"fmt"
	"os"

	"github.com/yourusername/minikube-mcp/internal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
```

### Step 9: Build and Test CLI Flags

```bash
# Rebuild
make build
# or
go build -o bin/minikube-mcp ./cmd/minikube-mcp

# Test with flags
./bin/minikube-mcp --help
./bin/minikube-mcp --transport stdio --log-level debug
```

**✅ Checkpoint:** You should see help text and be able to pass flags!

---

## 🎯 Stage 3: Understanding MCP Protocol

Before writing MCP code, let's understand what we're building.

### MCP Communication Flow

```
1. AI Assistant sends JSON-RPC request:
{
  "jsonrpc": "2.0",
  "method": "tools/call",
  "params": {
    "name": "minikube_cluster_status",
    "arguments": {}
  },
  "id": 1
}

2. Your MCP Server processes it:
   - Receives JSON via stdin
   - Parses the tool name
   - Executes: `minikube status`
   - Formats response

3. Your MCP Server responds:
{
  "jsonrpc": "2.0",
  "result": {
    "content": [
      {
        "type": "text",
        "text": "Cluster is running..."
      }
    ]
  },
  "id": 1
}
```

### Key MCP Concepts

1. **Tools:** Functions the AI can call (like `minikube_cluster_status`)
2. **Transport:** How data is sent (stdio = stdin/stdout)
3. **JSON-RPC:** The protocol format
4. **Tool Schema:** Defines what parameters a tool accepts

---

## 🎯 Stage 4: Your First MCP Tool

### Step 10: Create the MCP Server

Create `internal/server/server.go`:

```go
package server

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/server"
)

type MinikubeMCPServer struct {
	mcpServer *server.MCPServer
}

func NewMinikubeMCPServer() *MinikubeMCPServer {
	s := &MinikubeMCPServer{}
	
	// Create MCP server
	s.mcpServer = server.NewMCPServer(
		"minikube-mcp",
		"1.0.0",
		server.WithToolCapabilities(true),
	)
	
	// Register tools
	s.registerTools()
	
	return s
}

func (s *MinikubeMCPServer) registerTools() {
	// We'll add tools here
	fmt.Println("Registering tools...")
}

func (s *MinikubeMCPServer) Start(transport string) error {
	fmt.Printf("Starting MCP server with transport: %s\n", transport)
	
	ctx := context.Background()
	
	// Start stdio transport
	if transport == "stdio" {
		return s.mcpServer.Serve(ctx)
	}
	
	return fmt.Errorf("unsupported transport: %s", transport)
}
```

### Step 11: Create Your First Tool

Create `internal/tools/cluster.go`:

```go
package tools

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
)

// ClusterStatusTool returns the status of the Minikube cluster
func ClusterStatusTool() mcp.Tool {
	return mcp.Tool{
		Name:        "minikube_cluster_status",
		Description: "Get the current status of the Minikube cluster",
		InputSchema: mcp.ToolInputSchema{
			Type:       "object",
			Properties: map[string]interface{}{},
			Required:   []string{},
		},
		Handler: handleClusterStatus,
	}
}

func handleClusterStatus(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Execute minikube status command
	cmd := exec.CommandContext(ctx, "minikube", "status")
	output, err := cmd.CombinedOutput()
	
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get status: %v", err)), nil
	}
	
	// Return the result
	return mcp.NewToolResultText(string(output)), nil
}
```

### Step 12: Wire Everything Together

Update `internal/server/server.go` to register the tool:

```go
func (s *MinikubeMCPServer) registerTools() {
	fmt.Println("Registering tools...")
	
	// Register cluster status tool
	s.mcpServer.AddTool(tools.ClusterStatusTool())
	
	fmt.Println("✓ Registered minikube_cluster_status")
}
```

And update `internal/cmd/root.go` to start the server:

```go
var rootCmd = &cobra.Command{
	Use:   "minikube-mcp",
	Short: "Minikube MCP Server",
	Long:  `A Model Context Protocol server for Minikube Kubernetes clusters.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🚀 Minikube MCP Server starting...")
		
		srv := server.NewMinikubeMCPServer()
		return srv.Start(transport)
	},
}
```

### Step 13: Build and Test!

```bash
# Build
make build

# Start the server (it will wait for input)
./bin/minikube-mcp --transport stdio
```

In another terminal, test it:

```bash
# Send a test request
echo '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"minikube_cluster_status","arguments":{}},"id":1}' | ./bin/minikube-mcp --transport stdio
```

**✅ Checkpoint:** You should see JSON output with your cluster status!

---

## 🎯 Next Steps

Congratulations! You've built your first MCP tool! 🎉

**What you've learned:**
- How to structure a Go project
- How to use Cobra for CLI
- How MCP protocol works
- How to execute system commands from Go
- How to create and register MCP tools

**Continue to Stage 3:** Add more cluster operations (start, stop, delete)

**Questions to explore:**
1. What happens if Minikube isn't installed?
2. How would you add error handling?
3. How can you format the output better?

**Update your LEARNING_LOG.md** with what you learned!

---

## 🆘 Troubleshooting

### "cannot find package"
```bash
go mod download
go mod tidy
```

### "command not found: minikube"
Make sure Minikube is installed and in your PATH

### "permission denied"
```bash
chmod +x bin/minikube-mcp
```

### Server doesn't respond
- Check if it's waiting for input (that's correct!)
- Make sure you're sending valid JSON
- Check for syntax errors in your JSON

---

**Ready for the next stage?** Reply with what you learned and any questions!
