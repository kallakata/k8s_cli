package parser

import (
	"context"
	"fmt"
	"os"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"k8s.io/client-go/kubernetes"
	"path/filepath"
	"k8s.io/client-go/tools/clientcmd"
	"time"
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

	tm := table.NewWriter()
	tm.SetOutputMirror(os.Stdout)
	tm.AppendRows([]table.Row{
		{"Getting pods..."},
		})
	tm.Style().Format.Row = text.FormatLower
	tm.Render()
	time.Sleep(2 * time.Second)
	pods, err := clientset.CoreV1().Pods(ns).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		err = fmt.Errorf("error getting pods: %v\n", err)
		return nil, nil, err
	}
	if err != nil {
		os.Exit(1)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Pod", "Status"})

	for _, pod := range pods.Items {
		t.AppendRows([]table.Row{
		{pod.Name, ""},
		})
		t.AppendRows([]table.Row{
		{"", pod.Status.Phase},
		})
	}
	var message string
	if ns == "" {
		t.AppendFooter(table.Row{"Total pods in namespaces: ", len(pods.Items)})
	} else {
		message = fmt.Sprintf("Total Pods in namespace `%s`", ns)
		t.AppendFooter(table.Row{message, len(pods.Items)})
	}

	t.SetStyle(table.StyleColoredDark)
	t.Style().Color.Header = text.Colors{text.BgBlack, text.FgRed}
    t.Style().Color.Footer = text.Colors{text.BgBlack, text.FgRed}
	t.Render()
	return pods, clientset, nil
}