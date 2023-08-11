package cli

import (
	"fmt"
	"context"
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
)

func ListNamespaces(client kubernetes.Interface) (*v1.NamespaceList, error) {
    fmt.Println("Get Kubernetes Namespaces")
    namespaces, err := client.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
    if err != nil {
        err = fmt.Errorf("error getting namespaces: %v\n", err)
        return nil, err
    }
    return namespaces, nil
}