package parser

import (
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kallakata/k8s_cli/model"
	"github.com/kallakata/k8s_cli/pretty/pretty_pods"
	"github.com/kallakata/k8s_cli/prompt/prompt_pods"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"time"
)

func ListPods(ns string, ctx string) ([]model.Pod, *kubernetes.Clientset, error) {
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
			Pod:       pod.Name,
			Status:    string(pod.Status.Phase),
			Namespace: ns,
			CPUReq: pod.Spec.Containers[0].Resources.Requests.Cpu().String(),
			CPULim: pod.Spec.Containers[0].Resources.Limits.Cpu().String(),
			MemReq: pod.Spec.Containers[0].Resources.Requests.Memory().String(),
			MemLim: pod.Spec.Containers[0].Resources.Limits.Memory().String(),
			Image: pod.Spec.Containers[0].Image,
			Context:   ctx,
		}
		items = append(items, item)
	}

	p := tea.NewProgram(pretty_pods.NewModel(items, ctx, ns))
	fmt.Printf("\n\n========== | Getting pods | ==========\n\n")
	time.Sleep(2 * time.Second)
	p.Run()

	return items, clientset, nil
}

func ListPodsUsingPrompt(ctx string) (error) {

	// Create a new prompt model and run it
	promptModel := prompt_pods.InitialModel()
	promptProgram := tea.NewProgram(promptModel)
	resultMsg, _ := promptProgram.Run()
	
	// Check if the result message is a model with a GetNamespace method
	if resultMsg != nil {
		if namespaceModel, ok := resultMsg.(interface{ GetNamespace() string }); ok {
			ns := namespaceModel.GetNamespace()
			ctx := ctx // Replace with the appropriate context
			checkNs := ListNamespacesShort(ctx, ns)
			if checkNs {
				ListPods(ns, ctx)
			} else {
				os.Exit(1)
			}
		}
	}

	return fmt.Errorf("failed to get namespace from prompt")
}
