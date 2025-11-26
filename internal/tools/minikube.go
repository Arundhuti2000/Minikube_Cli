package tools

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// MinikubeStatus represents the JSON output of minikube status
type MinikubeStatus struct {
	Name       string `json:"Name"`
	Host       string `json:"Host"`
	Kubelet    string `json:"Kubelet"`
	APIServer  string `json:"APIServer"`
	Kubeconfig string `json:"Kubeconfig"`
}

// StartCluster starts the minikube cluster
func StartCluster() (string, error) {
	cmd := exec.Command("minikube", "start")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to start minikube: %s, output: %s", err, string(output))
	}
	return string(output), nil
}

// StopCluster stops the minikube cluster
func StopCluster() (string, error) {
	cmd := exec.Command("minikube", "stop")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to stop minikube: %s, output: %s", err, string(output))
	}
	return string(output), nil
}

// GetStatus returns the status of the minikube cluster
func GetStatus() (*MinikubeStatus, error) {
	cmd := exec.Command("minikube", "status", "-o", "json")
	output, err := cmd.Output()
	if err != nil {
		// If minikube is not running, it might return an error code.
		// We can try to parse the output anyway if it's JSON, or just return the error.
		// For simplicity, we'll return the error for now.
		return nil, fmt.Errorf("failed to get status: %s", err)
	}

	var status MinikubeStatus
	if err := json.Unmarshal(output, &status); err != nil {
		return nil, fmt.Errorf("failed to parse status json: %s", err)
	}
	return &status, nil
}
