package cmd

import (
	"fmt"
	"os"

	"github.com/Arundhuti2000/Minikube_Cli/internal/server"
	"github.com/spf13/cobra"
)

var (
	transport string
	logLevel string 
	accessLevel string
)
var rootCmd =&cobra.Command{
	Use: "minicube-mcp",
	Short: "Minikube MCP Server",
	Long: "A Model Context Protocol server for Minikube Kubernetes clusters.",
	Run: func(cmd *cobra.Command, args []string){
		s := server.NewMinikubeServer("minikube-mcp", "0.1.0")
		s.RegisterTools()

		if err := s.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init(){
	rootCmd.Flags().StringVar(&transport,"transport","stdio","Transport Mechanism (stdio, sse, http)")
	rootCmd.Flags().StringVar(&logLevel, "log-level", "info", "Log level (debug, info, warn, error)")
	rootCmd.Flags().StringVar(&accessLevel, "access-level", "readonly", "Access level (readonly, readwrite)")
}


func Execute() error {
	return rootCmd.Execute()
}