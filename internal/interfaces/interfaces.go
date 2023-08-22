package interfaces

import "github.com/kallakata/k8s_cli/model"

type NodePoolsFetcher interface {
    FetchNodePoolsForCluster(projectID, zone, clusterName string) ([]model.Nodepool, error)
}

type nodePoolsFetcher struct {}