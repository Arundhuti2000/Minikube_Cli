# Minikube MCP Server 🚀

A Model Context Protocol (MCP) server that enables AI assistants to interact with Minikube Kubernetes clusters. This project serves as a bridge between AI tools (like Claude, GitHub Copilot, and other MCP-compatible tools) and Minikube, translating natural language requests into Minikube operations similar to aks agentic cli.

## 🎯 Learning Project

This is a hands-on learning project built step-by-step to understand:

- How MCP (Model Context Protocol) works
- Building CLI tools in Go
- Integrating with Kubernetes/Minikube
- Creating AI-powered developer tools

Inspired by boot.dev's learning approach - we build to learn! 🎓

## 🌟 What This Does

The Minikube-MCP server allows AI assistants to:

- Manage Minikube clusters (start, stop, status, delete)
- Interact with Kubernetes resources (pods, services, deployments)
- Execute kubectl commands safely
- Debug and diagnose cluster issues
- Manage Minikube addons and configurations

**Example interactions with AI:**

```
"Start my minikube cluster"
"List all pods in the default namespace"
"What's the status of my cluster?"
"Enable the ingress addon"
"Show me the logs for the nginx pod"
```

## 🏗️ Architecture

```
┌─────────────────────┐
│   AI Assistant      │
│  (Claude, Copilot)  │
└──────────┬──────────┘
           │
           │ MCP Protocol
           │ (JSON-RPC)
           │
┌──────────▼──────────┐
│  Minikube MCP       │
│     Server          │
│                     │
│  ┌───────────────┐  │
│  │  MCP Tools    │  │
│  │  - cluster_*  │  │
│  │  - kubectl_*  │  │
│  │  - addon_*    │  │
│  └───────┬───────┘  │
└──────────┼──────────┘
           │
           │ CLI Commands
           │
┌──────────▼──────────┐
│   Minikube CLI      │
│   kubectl           │
└─────────────────────┘
           │
┌──────────▼──────────┐
│  Minikube Cluster   │
│    (Kubernetes)     │
└─────────────────────┘
```

## 📋 Prerequisites

Before you start building, make sure you have:

- **Go** 1.21 or higher ([Install Go](https://go.dev/doc/install))
- **Minikube** ([Install Minikube](https://minikube.sigs.k8s.io/docs/start/))
- **kubectl** ([Install kubectl](https://kubernetes.io/docs/tasks/tools/))
- **Basic Go knowledge** (or willingness to learn!)
- **Git** for version control

Verify your installations:

```bash
go version          # Should show 1.21+
minikube version    # Should show minikube version
kubectl version     # Should show kubectl version
```

## 🚀 Quick Start

### 1. Clone and Setup

```bash
git clone https://github.com/yourusername/minikube-mcp
cd minikube-mcp
go mod download
```

### 2. Build the Server

```bash
make build
# or
go build -o minikube-mcp ./cmd/minikube-mcp
```

### 3. Run the Server

```bash
./minikube-mcp --transport stdio
```

### 4. Configure with Your AI Assistant

#### For VS Code (GitHub Copilot)

Create `.vscode/mcp.json`:

```json
{
  "servers": {
    "minikube-mcp": {
      "type": "stdio",
      "command": "/path/to/minikube-mcp",
      "args": ["--transport", "stdio"]
    }
  }
}
```

#### For Claude Desktop

Add to `~/Library/Application Support/Claude/claude_desktop_config.json` (macOS):

```json
{
  "mcpServers": {
    "minikube": {
      "command": "/path/to/minikube-mcp",
      "args": ["--transport", "stdio"]
    }
  }
}
```

## 📁 Project Structure

```
minikube-mcp/
├── cmd/
│   └── minikube-mcp/
│       └── main.go              # Application entry point
├── internal/
│   ├── cmd/
│   │   └── root.go              # CLI command setup (Cobra)
│   ├── server/
│   │   ├── server.go            # MCP server implementation
│   │   └── transport.go         # Transport layer (stdio)
│   ├── tools/
│   │   ├── cluster.go           # Cluster management tools
│   │   ├── kubectl.go           # kubectl wrapper tools
│   │   ├── addons.go            # Addon management tools
│   │   └── diagnostics.go       # Diagnostic tools
│   └── executor/
│       └── command.go           # Command execution utilities
├── pkg/
│   └── minikube/
│       └── client.go            # Minikube API wrapper
├── go.mod
├── go.sum
├── Makefile                     # Build automation
├── README.md                    # This file!
└── LICENSE
```

## 🛠️ Available Tools

### Cluster Management

- `minikube_cluster_start` - Start a Minikube cluster
- `minikube_cluster_stop` - Stop the cluster
- `minikube_cluster_status` - Get cluster status
- `minikube_cluster_delete` - Delete the cluster
- `minikube_cluster_list` - List all profiles

### Kubernetes Resources

- `kubectl_get` - Get resources (pods, services, deployments, etc.)
- `kubectl_describe` - Describe resources in detail
- `kubectl_logs` - Get container logs
- `kubectl_exec` - Execute commands in containers
- `kubectl_apply` - Apply configurations
- `kubectl_delete` - Delete resources

### Addons

- `minikube_addon_list` - List available addons
- `minikube_addon_enable` - Enable an addon
- `minikube_addon_disable` - Disable an addon

### Diagnostics

- `minikube_diagnostics` - Run cluster diagnostics
- `minikube_ssh` - SSH into the cluster node
- `kubectl_cluster_info` - Get cluster information

## 🎓 Learning Path

This project is designed to be built in stages. Follow along:

### Stage 1: Foundation (You Are Here!)

- [ ] Set up project structure
- [ ] Initialize Go module
- [ ] Understand MCP protocol basics
- [ ] Create basic CLI with Cobra

### Stage 2: First MCP Tool

- [ ] Implement MCP server scaffolding
- [ ] Create your first tool: `minikube_cluster_status`
- [ ] Test with stdio transport
- [ ] Understand JSON-RPC communication

### Stage 3: Cluster Operations

- [ ] Implement cluster start/stop/delete
- [ ] Add error handling
- [ ] Add input validation
- [ ] Test each operation

### Stage 4: kubectl Integration

- [ ] Wrap kubectl commands
- [ ] Implement resource listing (get pods, services)
- [ ] Add namespace support
- [ ] Handle kubectl output parsing

### Stage 5: Advanced Features

- [ ] Add addon management
- [ ] Implement diagnostics
- [ ] Add configuration options
- [ ] Create comprehensive tests

### Stage 6: Polish & Distribution

- [ ] Add proper logging
- [ ] Create Makefile for builds
- [ ] Write documentation
- [ ] Create release binaries
- [ ] Docker containerization

## 🧪 Development

### Running Tests

```bash
make test
# or
go test ./...
```

### Running with Debug Logging

```bash
./minikube-mcp --transport stdio --log-level debug
```

### Testing Individual Tools

```bash
# Test cluster status
echo '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"minikube_cluster_status","arguments":{}},"id":1}' | ./minikube-mcp --transport stdio
```

## 🤝 Contributing

This is a learning project, but contributions are welcome! Whether you're:

- Learning alongside and want to share improvements
- Found a bug or issue
- Have ideas for new tools
- Want to improve documentation

Feel free to open issues and pull requests!

## 📚 Resources

### Understanding MCP

- [MCP Specification](https://spec.modelcontextprotocol.io/)
- [MCP Go SDK](https://github.com/mark3labs/mcp-go)
- [Anthropic MCP Docs](https://docs.anthropic.com/claude/docs/mcp)

### Go Resources

- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Cobra CLI](https://github.com/spf13/cobra)

### Kubernetes/Minikube

- [Minikube Documentation](https://minikube.sigs.k8s.io/docs/)
- [kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
- [Kubernetes API](https://kubernetes.io/docs/reference/kubernetes-api/)

## 🔧 Configuration Options

### Command Line Flags

```bash
--transport string      Transport mechanism (stdio, sse, http) [default: stdio]
--access-level string   Access level (readonly, readwrite) [default: readonly]
--log-level string      Log level (debug, info, warn, error) [default: info]
--timeout int           Command timeout in seconds [default: 300]
--profile string        Minikube profile to use [default: minikube]
```

### Environment Variables

```bash
MINIKUBE_HOME          # Minikube home directory
KUBECONFIG             # Path to kubeconfig file
MINIKUBE_PROFILE       # Default profile to use
```

## 🐛 Troubleshooting

### Server Won't Start

- Check if Minikube is installed: `minikube version`
- Check if kubectl is installed: `kubectl version`
- Verify the binary has execute permissions: `chmod +x minikube-mcp`

### Tools Not Showing in AI Assistant

- Restart your AI assistant/IDE
- Check MCP server logs for errors
- Verify the server is running: `ps aux | grep minikube-mcp`
- Check your MCP configuration file path

### Command Timeouts

- Increase timeout: `--timeout 600`
- Check if Minikube cluster is responsive: `minikube status`
- Check network connectivity

## 📝 License

MIT License - see LICENSE file for details

## 🙏 Acknowledgments

- Inspired by [Azure/aks-mcp](https://github.com/Azure/aks-mcp)
- Built with [mcp-go](https://github.com/mark3labs/mcp-go)
- Learning approach inspired by [boot.dev](https://boot.dev)

## 🗺️ Roadmap

- [x] Project setup and README
- [ ] Basic MCP server implementation
- [ ] Cluster management tools
- [ ] kubectl integration
- [ ] Addon management
- [ ] Advanced diagnostics
- [ ] Docker support
- [ ] Helm integration
- [ ] Multi-cluster support
- [ ] CI/CD pipeline

---

**Ready to start building?** Follow the learning path above and let's create something awesome! 🚀

**Questions or stuck?** Open an issue or check the Resources section above.

**Learning tip:** Build one tool at a time, test it thoroughly, then move to the next. Don't try to build everything at once!
