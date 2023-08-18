package parser

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/kallakata/k8s_cli/model"
	"github.com/kallakata/k8s_cli/pretty"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

func ListNamespaces(ctx string) ([]model.Ns, *kubernetes.Clientset, error) {
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
    namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
    if err != nil {
        err = fmt.Errorf("error getting namespaces: %v\n", err)
        return nil, nil, err
    }
	if err != nil {
		os.Exit(1)
	}
	var items []model.Ns
    for _, namespace := range namespaces.Items {
        item := model.Ns{
            Namespace:  namespace.Name,
        }
        items = append(items, item)
    }

	p := tea.NewProgram(pretty.NewNsModel(items, ctx))
	fmt.Printf("========== Getting namespaces ==========\n\n")
	time.Sleep(2 * time.Second)
	p.Run()

    return items, clientset, nil
}