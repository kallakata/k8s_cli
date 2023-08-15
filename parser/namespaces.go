package parser

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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
	tm := table.NewWriter()
	tm.SetOutputMirror(os.Stdout)
	tm.AppendRows([]table.Row{
		{"Getting namespaces..."},
		})
	tm.Style().Format.Row = text.FormatLower
	tm.Render()
    namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
    if err != nil {
        err = fmt.Errorf("error getting namespaces: %v\n", err)
        return nil, nil, err
    }
	if err != nil {
		os.Exit(1)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Namespace"})

	for _, namespace := range namespaces.Items {
		// fmt.Println(namespace.Name)
		t.AppendRows([]table.Row{
			{namespace.Name},
		})
	}
	// fmt.Printf("Total namespaces: %d\n", len(namespaces.Items))
	t.AppendFooter(table.Row{"Total: ", len(namespaces.Items)})
	t.SetStyle(table.StyleColoredDark)
	t.Style().Color.Header = text.Colors{text.BgBlack, text.FgRed}
    t.Style().Color.Footer = text.Colors{text.BgBlack, text.FgRed}
	t.Render()
    return namespaces, clientset, nil
}