package cli

import (
	"fmt"
	"context"
    "os"
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
)

var globalClient kubernetes.Interface


func ListNamespaces(namespace string) (*v1.NamespaceList, error) {
    fmt.Println("Getting namespaces...")
    namespaces, err := globalClient.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
    if err != nil {
        err = fmt.Errorf("error getting namespaces: %v\n", err)
        return nil, err
    }
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}
	for _, namespace := range namespaces.Items {
		fmt.Println(namespace.Name)
	}
	fmt.Printf("Total namespaces: %d\n", len(namespaces.Items))
    return namespaces, nil
}