package parser

import (
	"context"
	"fmt"
	"log"

	// "github.com/jedib0t/go-pretty/v6/table"
	// "github.com/jedib0t/go-pretty/v6/text"
	// "golang.org/x/oauth2/google"
	container "google.golang.org/api/container/v1"
)

// var globalClient kubernetes.Interface


func ListClusters(projectID, zone string) error {
	ctx := context.Background()
	// hc, err := google.DefaultClient(ctx, container.CloudPlatformScope)
	// if err != nil {
	// 	log.Fatalf("Could not get authenticated client: %v", err)
	// }

	svc, err := container.NewService(ctx)
	if err != nil {
		log.Fatalf("Could not initialize gke client: %v", err)
	}

	list, err := svc.Projects.Zones.Clusters.List(projectID, zone).Do()
	if err != nil {
		return fmt.Errorf("failed to list clusters: %w", err)
	}
	for _, v := range list.Clusters {
		fmt.Printf("Cluster %q (%s) master_version: v%s", v.Name, v.Status, v.CurrentMasterVersion)

		poolList, err := svc.Projects.Zones.Clusters.NodePools.List(projectID, zone, v.Name).Do()
		if err != nil {
			return fmt.Errorf("failed to list node pools for cluster %q: %w", v.Name, err)
		}
		for _, np := range poolList.NodePools {
			fmt.Printf("  -> Pool %q (%s) machineType=%s node_version=v%s autoscaling=%v", np.Name, np.Status,
				np.Config.MachineType, np.Version, np.Autoscaling != nil && np.Autoscaling.Enabled)
		}
	}
	return nil
}