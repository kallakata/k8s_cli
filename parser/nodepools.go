package parser

import (
	container "cloud.google.com/go/container/apiv1"
	containerpb "cloud.google.com/go/container/apiv1/containerpb"
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kallakata/k8s_cli/model"
	// "github.com/kallakata/k8s_cli/pretty/pretty_nodepools"
	"github.com/kallakata/k8s_cli/prompt/prompt_nodepools"
	// "log"
	// "time"
)

func ListNodepools(project, zone, cluster string) ([]model.Nodepool, error) {
	ctx := context.Background()

	c, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("Could not initialize gke client: %v", err)
	}

	defer c.Close()

	var nodepools []model.Nodepool

	req := &containerpb.ListNodePoolsRequest{
		ProjectId: project,
		Zone:      zone,
		ClusterId: cluster,
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
		fmt.Printf("  -> Pool %q (%s) machineType=%s node_version=v%s autoscaling=%v", np.Name, np.Status,
			np.Config.MachineType, np.Version, np.Autoscaling != nil && np.Autoscaling.Enabled)
	}

	// fmt.Println(nodepools)

	// p := tea.NewProgram(pretty_nodepools.NewModel(nodepools))
	// fmt.Printf("========== Getting nodepools ==========\n\n")
	// time.Sleep(2 * time.Second)
	// defer func() {
	// 	if _, err := p.Run(); err != nil {
	// 		log.Printf("Error running Bubble Tea program: %v", err)
	// 	}
	// }()

	return nodepools, nil
}

func ListPoolsUsingPrompt(project, zone string) ([]model.Nodepool, error) {

	// Create a new prompt model and run it
	promptModel := prompt_nodepools.InitialModel()
	promptProgram := tea.NewProgram(promptModel)
	resultMsg, _ := promptProgram.Run()

	// Check if the result message is a model with a GetCluster method
	if clusterModel, ok := resultMsg.(interface{ GetCluster() string }); ok {
		cluster := clusterModel.GetCluster()
		project := project
		zone := zone
		fmt.Printf(cluster, project, zone)
		nodepools, err := ListNodepools(project, zone, cluster)
		if err != nil {
			return nil, err
		}
		return nodepools, nil
	}

	return nil, fmt.Errorf("failed to get cluster from prompt")
}

