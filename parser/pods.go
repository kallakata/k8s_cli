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

func ListPods(ns string, ctx string) (*v1.PodList, *kubernetes.Clientset, error) {
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
	fmt.Println("Getting pods...")
	pods, err := clientset.CoreV1().Pods(ns).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		err = fmt.Errorf("error getting pods: %v\n", err)
		return nil, nil, err
	}
	if err != nil {
		os.Exit(1)
	}

	for _, pod := range pods.Items {
		fmt.Printf("Pod name: %v\n", pod.Name)
		fmt.Printf("Pod status: %v\n", pod.Status.Phase)
	}
	var message string
	if ns == "" {
		message = "Total Pods in all namespaces"
	} else {
		message = fmt.Sprintf("Total Pods in namespace `%s`", ns)
	}
	fmt.Printf("%s %d\n", message, len(pods.Items))
	return pods, clientset, nil
}