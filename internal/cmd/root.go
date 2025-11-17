package cmd

import (
	"fmt"

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
		fmt.Println("Minikube MCP Server starting...")
		fmt.Printf("Transport: %s\n", transport)
		fmt.Printf("Log Level: %s\n", logLevel)
		fmt.Printf("Access Level: %s\n", accessLevel)
		
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