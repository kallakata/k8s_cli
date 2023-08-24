package parser

import (
	"os"
	// "log"
	// "time"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/fatih/color"
	container "cloud.google.com/go/container/apiv1"
	containerpb "cloud.google.com/go/container/apiv1/containerpb"
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kallakata/k8s_cli/model"
	// "github.com/kallakata/k8s_cli/pretty/pretty_nodepools"
	"github.com/kallakata/k8s_cli/prompt/prompt_nodepools"
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
		return nil, fmt.Errorf("Could not initialize gke client: %v", err)
	}

	defer c.Close()

	var nodepools []model.Nodepool

	parent := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", project, zone, cluster)

	req := &containerpb.ListNodePoolsRequest{
		Parent: parent,
	}

	resp, err := c.ListNodePools(ctx, req)

	// for _, np := range resp.NodePools {
	// 	nodepool := model.Nodepool{
	// 		Nodepool:    np.Name,
	// 		Status:      string(np.Status),
	// 		Version:     np.Version,
	// 		Autoscaling: np.Autoscaling.Enabled,
	// 		MinNode:     np.Autoscaling.MinNodeCount,
	// 		MaxNode:     np.Autoscaling.MaxNodeCount,
	// 	}
	// 	nodepools = append(nodepools, nodepool)
	// }

	for _, np := range resp.NodePools {
		color.Magenta("\n-> Pool %q\n", np.Name)
		color.Magenta("  | version: v%s\n", np.Version)
		color.Green("  | status: %s\n", np.Status)
		color.Yellow("  | machineType: %s\n", np.Config.MachineType)
		color.Yellow("  | autoscaling: %v\n", np.Autoscaling != nil && np.Autoscaling.Enabled)
	}

	// fmt.Println(nodepools)

	// p := tea.NewProgram(pretty_nodepools.NewModel(nodepools))
	// fmt.Printf("========== Getting nodepools ==========\n\n")
	// time.Sleep(2 * time.Second)
	// p.Run()

	return nodepools, nil
}

func ListPoolsUsingPrompt(project, zone string) ([]model.Nodepool, error) {

	// Create a new prompt model and run it
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
        return false, fmt.Errorf("Could not initialize gke client: %v", err)
    }
    defer c.Close()

    clusterName := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", project, location, cluster)
    req := &containerpb.GetClusterRequest{
        Name: clusterName,
    }

    _, err = c.GetCluster(ctx, req)
    if err != nil {
        // Handle the error, cluster doesn't exist
        if status.Code(err) == codes.NotFound {
            return false, nil
        }
        return false, fmt.Errorf("Error getting cluster: %v", err)
    }

    return true, nil
}

