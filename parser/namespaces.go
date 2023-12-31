package parser

import (
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kallakata/k8s_cli/model"
	"github.com/kallakata/k8s_cli/pretty/pretty_ns"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"time"
)

func ListNamespaces(ctx string) ([]model.Ns, *kubernetes.Clientset, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Getting user home dir failed: %v\n", err)
		os.Exit(1)
	}
	kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")

	var clientset *kubernetes.Clientset
	if ctx == "" {
		// If ctx is not provided as an argument, get the current Kubernetes context
		kubeConfig, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath},
			&clientcmd.ConfigOverrides{}).RawConfig()
		if err != nil {
			fmt.Printf("Error getting kubernetes config: %v\n", err)
			os.Exit(1)
		}
		ctx = kubeConfig.CurrentContext
	}

	configLoadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath}
	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: ctx}

	kubeConfig, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(configLoadingRules, configOverrides).ClientConfig()
	if err != nil {
		fmt.Printf("Error getting kubernetes config: %v\n", err)
		os.Exit(1)
	}

	clientset, err = kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		fmt.Printf("Error preparing new clientset: %v\n", err)
	}
	namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error getting namespaces: %v\n", err)
		return nil, nil, err
	}
	if err != nil {
		os.Exit(1)
	}
	var items []model.Ns
	for _, namespace := range namespaces.Items {
		pods, err := clientset.CoreV1().Pods(namespace.Name).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			fmt.Printf("Error getting pods in namespace %s: %v\n", namespace.Name, err)
			continue
		}
		item := model.Ns{
			Namespace: namespace.Name,
			Pods:      len(pods.Items),
		}
		items = append(items, item)
	}

	p := tea.NewProgram(pretty_ns.NewModel(items, ctx))
	fmt.Printf("========== | Getting namespaces... | ==========\n\n")
	time.Sleep(1 * time.Second)
	p.Run()

	return items, clientset, nil
}

func ListNamespacesShort(ctx, ns string) bool {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Getting user home dir failed: %v\n", err)
		os.Exit(1)
	}
	kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")

	var clientset *kubernetes.Clientset
	if ctx == "" {
		// If ctx is not provided as an argument, get the current Kubernetes context
		kubeConfig, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath},
			&clientcmd.ConfigOverrides{}).RawConfig()
		if err != nil {
			fmt.Printf("Error getting kubernetes config: %v\n", err)
			os.Exit(1)
		}
		ctx = kubeConfig.CurrentContext
	}

	configLoadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath}
	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: ctx}

	kubeConfig, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(configLoadingRules, configOverrides).ClientConfig()
	if err != nil {
		fmt.Printf("Error getting Kubernetes config: %v\n", err)
		os.Exit(1)
	}
	// Get clientset
	clientset, err = kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		fmt.Printf("Error creating Kubernetes client: %v\n", err)
		os.Exit(1)
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error getting namespaces: %v\n", err)
		os.Exit(1)
	}

	for _, namespace := range namespaces.Items {
		if namespace.Name == ns || ns == "" {
			return true
		}
	}

	fmt.Printf("Namespace %s doesn't exist in the cluster.\n", ns)
	return false
}
