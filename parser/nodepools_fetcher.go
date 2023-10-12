package parser

import (
	"fmt"
    "github.com/kallakata/k8s_cli/model"
	container "cloud.google.com/go/container/apiv1"
	containerpb "cloud.google.com/go/container/apiv1/containerpb"
	"github.com/kallakata/k8s_cli/internal/interfaces"
	"context"
)

type nodePoolsFetcher struct{}

func (npf nodePoolsFetcher) FetchNodePoolsForCluster(projectID, zone, clusterName string) ([]model.Nodepool, error) {
    ctx := context.Background()

    c, err := container.NewClusterManagerClient(ctx)
    if err != nil {
        return nil, fmt.Errorf("could not initialize gke client: %v", err)
    }
    defer c.Close()

    req := &containerpb.ListNodePoolsRequest{
        Parent: fmt.Sprintf("projects/%s/locations/%s/clusters/%s", projectID, zone, clusterName),
    }
    resp, err := c.ListNodePools(ctx, req)
    if err != nil {
        return nil, fmt.Errorf("error listing node pools: %v", err)
    }

    var nodePools []model.Nodepool
    for _, np := range resp.NodePools {
        nodePool := model.Nodepool{
            Nodepool:   np.Name,
            // Status: np.Conditions{Message: message},
			Version: np.Version,
        }
        nodePools = append(nodePools, nodePool)
    }

    return nodePools, nil
}

// Create an instance of the interface for external use
var NodePoolsFetcherInstance interfaces.NodePoolsFetcher = nodePoolsFetcher{}