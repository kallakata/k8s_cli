package parser

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/containerservice/mgmt/containerservice"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kallakata/k8s_cli/model"
	"github.com/kallakata/k8s_cli/pretty/pretty_azs_clusters"
)

func ListAzsClusters(subscription string) ([]model.AzureCluster, error) {
	ctx := context.Background()

	// Use Azure CLI authentication to get credentials
	authorizer, err := auth.NewAuthorizerFromCLI()
	if err != nil {
		fmt.Printf("Failed to authenticate: %v\n", err)
		return nil, err
	}

	client := containerservice.NewManagedClustersClient(subscription)
	client.Authorizer = authorizer

	clusters, err := client.List(ctx)
	if err != nil {
		fmt.Printf("Failed to list AKS clusters: %v\n", err)
		return nil, err
	}
	var azsClusters []model.AzureCluster

	for _, cluster := range clusters.Values() {
		if clusters.Values() == nil {
			fmt.Printf("No clusters found!")
		} else {
			currentStatus := string(*cluster.ManagedClusterProperties.ProvisioningState)
			location := string(*cluster.Location)
			identity := string(*cluster.ID)
			cluster := model.AzureCluster{
				Cluster:  *cluster.Name,
				Version:  *cluster.CurrentKubernetesVersion,
				Status:   currentStatus,
				Location: location,
				Identity: identity,
			}
			azsClusters = append(azsClusters, cluster)
		}
	}
	p := tea.NewProgram(pretty_azs_clusters.NewModel(azsClusters))
	fmt.Printf("\n\n========== | Getting clusters... | ==========\n\n")
	time.Sleep(1 * time.Second)
	p.Run()
	return azsClusters, nil
}
