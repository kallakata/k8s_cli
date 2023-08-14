package cli

import (
	"fmt"
	"os"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
)


func ListPods(namespace string) (*v1.PodList, error) {
	fmt.Println("Getting pods...")
	pods, err := globalClient.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		err = fmt.Errorf("error getting pods: %v\n", err)
		return nil, err
	}
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	for _, pod := range pods.Items {
		fmt.Printf("Pod name: %v\n", pod.Name)
		fmt.Printf("Pod status: %v\n", pod.Status.Phase)
	}
	var message string
	if namespace == "" {
		message = "Total Pods in all namespaces"
	} else {
		message = fmt.Sprintf("Total Pods in namespace `%s`", namespace)
	}
	fmt.Printf("%s %d\n", message, len(pods.Items))
	return pods, nil
}