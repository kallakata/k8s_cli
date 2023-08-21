package parser

import (
	container "cloud.google.com/go/container/apiv1"
	containerpb "cloud.google.com/go/container/apiv1/containerpb"
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kallakata/k8s_cli/model"
	"github.com/kallakata/k8s_cli/pretty/pretty_clusters"
	"time"
)

func ListClusters(projectID, zone string) ([]model.Cluster, error) {
	ctx := context.Background()

	c, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("Could not initialize gke client: %v", err)
	}

	defer c.Close()

	req := &containerpb.ListClustersRequest{
		ProjectId: projectID,
		Zone:      zone,
	}
	resp, err := c.ListClusters(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Error listing clusters: %v", err)
	}

	var clusters []model.Cluster
	for _, c := range resp.Clusters {
		cluster := model.Cluster{
			Cluster:  c.Name,
			Status:   string(c.Status),
			Version:  c.CurrentMasterVersion,
			Endpoint: c.Endpoint,
		}
		clusters = append(clusters, cluster)
	}

	p := tea.NewProgram(pretty_clusters.NewModel(clusters))
	fmt.Printf("========== Getting clusters ==========\n\n")
	time.Sleep(2 * time.Second)
	p.Run()

	// np := tea.NewProgram(pretty.NewNodepoolsModel(nodepools))
	// fmt.Printf("========== Getting nodepools ==========\n\n")
	// time.Sleep(2 * time.Second)
	// np.Run()

	for _, cluster := range resp.Clusters {
		fmt.Printf("Cluster Name: %s\n", cluster.Name)
		fmt.Printf("Cluster Status: %s\n", cluster.Status)
		fmt.Printf("Cluster Version: %s\n", cluster.CurrentMasterVersion)
		fmt.Printf("Cluster Endpoint: %s\n", cluster.Endpoint)
		fmt.Println("----------------------")
	}
	return clusters, nil
}
