package parser

import (
	"context"
	"fmt"
	"os"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"path/filepath"
	"k8s.io/client-go/tools/clientcmd"
)

// var globalClient kubernetes.Interface


func ListNamespaces(ctx string) (*v1.NamespaceList, *kubernetes.Clientset, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Getting user home dir failed: %v\n", err)
		os.Exit(1)
	}
	kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")

	configLoadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath}
	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: ctx}

	kubeConfig, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(configLoadingRules, configOverrides).ClientConfig()
    if err != nil {
        fmt.Printf("Error getting kubernetes config: %v\n", err)
        os.Exit(1)
    }

	clientset, err := kubernetes.NewForConfig(kubeConfig)
    fmt.Println("Getting namespaces...")
    namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
    if err != nil {
        err = fmt.Errorf("error getting namespaces: %v\n", err)
        return nil, nil, err
    }
	if err != nil {
		os.Exit(1)
	}
	for _, namespace := range namespaces.Items {
		fmt.Println(namespace.Name)
	}
	fmt.Printf("Total namespaces: %d\n", len(namespaces.Items))
    return namespaces, clientset, nil
}