// package auth

// import (
// 	"fmt"
// 	"os"
	// "path/filepath"
	// "k8s.io/client-go/kubernetes"
	// "k8s.io/client-go/tools/clientcmd"
// )

// func auth() {
// 	userHomeDir, err := os.UserHomeDir()
// 	if err != nil {
// 		fmt.Printf("Getting user home dir failed: %v\n", err)
// 		os.Exit(1)
// 	}
// 	kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")

// 	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
//     if err != nil {
//         fmt.Printf("Error getting kubernetes config: %v\n", err)
//         os.Exit(1)
//     }

// 	Clientset, err := kubernetes.NewForConfig(kubeConfig)

//     if err != nil {
//         fmt.Printf("Error getting kubernetes config: %v\n", err)
//         os.Exit(1)
//     }
// }