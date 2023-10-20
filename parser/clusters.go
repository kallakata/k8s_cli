// parser/clusters.go
package parser

import (
	"cloud.google.com/go/container/apiv1"
	"cloud.google.com/go/container/apiv1/containerpb"
	"context"
	"fmt"
	"github.com/kallakata/k8s_cli/model"
	"os"
	"time"
)

func ListClusters(projectID, zone string) ([]model.Cluster, error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	c, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("could not initialize gke client: %v", err)
	}
	defer c.Close()
	var parent string
	if len(zone) == 0 {
		parent = fmt.Sprintf("projects/%s/locations/%s", projectID, zone)
	} else {
		parent = fmt.Sprintf("projects/%s/locations/%s", projectID, "-")
	}

	req := &containerpb.ListClustersRequest{
		Parent: parent,
	}
	resp, err := c.ListClusters(ctx, req)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("error listing clusters: %v", err)
	}

	var clusters []model.Cluster

	if resp.Clusters != nil {

		for _, c := range resp.Clusters {
			var status model.ClusterStatus

			switch c.Status {
			case 0:
				status = model.StatusUnspecified
			case 1:
				status = model.Provisioning
			case 2:
				status = model.Running
			case 3:
				status = model.Reconciling
			case 4:
				status = model.Stopping
			case 5:
				status = model.Error
			case 6:
				status = model.Degraded
			default:
				status = model.StatusUnspecified // Default to unspecified if not recognized
			}

			cluster := model.Cluster{
				Cluster:  c.Name,
				Status:   status,
				Version:  c.CurrentMasterVersion,
				Endpoint: c.Endpoint,
			}

			clusters = append(clusters, cluster)
		}
	} else {
		fmt.Printf("No clusters found!")
		os.Exit(1)
	}

	fmt.Printf("========== Getting clusters ==========\n\n")
	time.Sleep(2 * time.Second)
	return clusters, nil
}
