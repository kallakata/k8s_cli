package parser

import (
	"k8s.io/client-go/tools/clientcmd"
	"fmt"
	"os"
	"path/filepath"
)

func getContext() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Getting user home dir failed: %v\n", err)
		os.Exit(1)
	}
	kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")
    config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
        &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath},
        &clientcmd.ConfigOverrides{
            CurrentContext: "",
        }).RawConfig()
	if err != nil {
		fmt.Printf("Getting default config failed: %v\n", err)
		os.Exit(1)
	}
    currentContext := config.CurrentContext
	fmt.Println(currentContext)
}