package parser

import (
	container "cloud.google.com/go/container/apiv1"
	containerpb "cloud.google.com/go/container/apiv1/containerpb"
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/kallakata/k8s_cli/model"
	"github.com/kallakata/k8s_cli/pretty/pretty_nodepools"
	"github.com/kallakata/k8s_cli/prompt/prompt_nodepools"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"time"
)

func ListNodepools(project, zone, cluster string) ([]model.Nodepool, error) {
	exists, err := ClusterExists(project, zone, cluster)
	if err != nil {
		return nil, err
	}

	if !exists {
		color.Red("Cluster '%s' does not exist. Please choose a different location, or cluster name", cluster)
		os.Exit(1)
	}

	ctx := context.Background()

	c, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not initialize gke client: %v", err)
	}

	defer c.Close()

	var nodepools []model.Nodepool

	parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", project, zone, cluster)

	req := &containerpb.ListNodePoolsRequest{
		Parent: parent,
	}

	resp, err := c.ListNodePools(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not list nodepools: %v", err)
	}

	for _, np := range resp.NodePools {
		nodepool := model.Nodepool{
			Nodepool:  np.Name,
			Status:    string(np.StatusMessage),
			Version:   np.Version,
			NodeCount: int(np.InitialNodeCount),
		}
		nodepools = append(nodepools, nodepool)
	}

	p := tea.NewProgram(pretty_nodepools.NewModel(nodepools))
	fmt.Printf("========== Getting nodepools ==========\n\n")
	time.Sleep(2 * time.Second)
	p.Run()

	return nodepools, nil
}

func ListPoolsUsingPrompt(project, zone string) ([]model.Nodepool, error) {

	promptModel := prompt_nodepools.InitialModel()
	promptProgram := tea.NewProgram(promptModel)
	resultMsg, _ := promptProgram.Run()

	// Check if the result message is a model with a GetCluster method
	if clusterModel, ok := resultMsg.(interface{ GetCluster() string }); ok && resultMsg.(interface{ GetCluster() string }) != nil {
		cluster := clusterModel.GetCluster()
		project := project
		zone := zone
		nodepools, err := ListNodepools(project, zone, cluster)
		if err != nil {
			return nil, err
		}
		return nodepools, nil
	}

	return nil, fmt.Errorf("failed to get cluster from prompt")
}

func ClusterExists(project, location, cluster string) (bool, error) {
	ctx := context.Background()

	c, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		return false, fmt.Errorf("could not initialize gke client: %v", err)
	}
	defer c.Close()

	clusterName := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", project, location, cluster)
	req := &containerpb.GetClusterRequest{
		Name: clusterName,
	}

	_, err = c.GetCluster(ctx, req)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return false, nil
		}
		return false, fmt.Errorf("error getting cluster: %v", err)
	}

	return true, nil
}
