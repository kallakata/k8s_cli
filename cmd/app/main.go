package main

import (
	"github.com/kallakata/k8s_cli/cmd/cli"
	// "github.com/kallakata/k8s_cli/auth"
	"fmt"
	"os"
	"path/filepath"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Getting user home dir failed: %v\n", err)
		os.Exit(1)
	}
	kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")

	var context string
	fmt.Println("Please enter context: ")
	fmt.Scanln(&context)

	configLoadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath}
	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: context}

	kubeConfig, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(configLoadingRules, configOverrides).ClientConfig()
    if err != nil {
        fmt.Printf("Error getting kubernetes config: %v\n", err)
        os.Exit(1)
    }

	clientset, err := kubernetes.NewForConfig(kubeConfig)

    if err != nil {
        fmt.Printf("Error getting kubernetes config: %v\n", err)
        os.Exit(1)
    }
	// An empty string returns all namespaces
	var namespace string
	fmt.Println("Enter desired namespace: ")
	fmt.Scanln(&namespace)

	pods, err := cli.ListPods(namespace, clientset)
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	for _, pod := range pods.Items {
		fmt.Printf("Pod name: %v\n", pod.Name)
		fmt.Printf("Pod status: %v\n", pod.Status.Phase)
	}
	var message string
	if namespace == "" {
		message = "Total Pods in all namespaces"
	} else {
		message = fmt.Sprintf("Total Pods in namespace `%s`", namespace)
	}
	fmt.Printf("%s %d\n", message, len(pods.Items))

	//ListNamespaces function call returns a list of namespaces in the kubernetes cluster
	namespaces, err := cli.ListNamespaces(clientset)
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}
	for _, namespace := range namespaces.Items {
		fmt.Println(namespace.Name)
	}
	fmt.Printf("Total namespaces: %d\n", len(namespaces.Items))
}