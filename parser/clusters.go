// parser/clusters.go
package parser

import (
    "context"
    "fmt"
    "time"
    "github.com/kallakata/k8s_cli/internal/interfaces"
    "github.com/kallakata/k8s_cli/model"
    "cloud.google.com/go/container/apiv1"
    "cloud.google.com/go/container/apiv1/containerpb"
)

func ListClusters(projectID, zone string, npf interfaces.NodePoolsFetcher) ([]model.Cluster, error) {
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

    fmt.Printf("========== Getting clusters ==========\n\n")
    time.Sleep(2 * time.Second)
    // Instead of directly running the program, return the clusters
    return clusters, nil
}
