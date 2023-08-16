package parser

import (
	"context"
	"fmt"
	"os"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	tea "github.com/charmbracelet/bubbletea"
	"k8s.io/client-go/kubernetes"
	"path/filepath"
	"github.com/kallakata/k8s_cli/model"
	"github.com/kallakata/k8s_cli/pretty"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

func ListPods(ns string, ctx string) ([]model.Pod, *kubernetes.Clientset, error) {
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

    if err != nil {
        fmt.Printf("Error getting kubernetes config: %v\n", err)
        os.Exit(1)
    }

	pods, err := clientset.CoreV1().Pods(ns).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		err = fmt.Errorf("error getting pods: %v\n", err)
		return nil, nil, err
	}
	if err != nil {
		os.Exit(1)
	}

	var items []model.Pod
    for _, pod := range pods.Items {
        item := model.Pod{
            Pod:  pod.Name,
			Status: string(pod.Status.Phase),
			Namespace: ns,
			Context: ctx,
        }
        items = append(items, item)
    }

	// var message string
	// if ns == "" {
	// 	t.AppendFooter(table.Row{"Total pods in namespaces: ", len(pods.Items)})
	// } else {
	// 	message = fmt.Sprintf("Total Pods in namespace `%s`", ns)
	// 	t.AppendFooter(table.Row{message, len(pods.Items)})
	// }

	p := tea.NewProgram(pretty.NewPodsModel(items, ctx, ns))
	fmt.Printf("========== Getting pods ==========\n\n")
	time.Sleep(2 * time.Second)
	p.Run()

	return items, clientset, nil
}