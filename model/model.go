package model

type Ns struct {
	Namespace string
}

type Pod struct {
	Pod       string
	Status    string
	Namespace string
	Requests string
	Limits  string
	Context   string
}

type ClusterStatus int

type NodepoolStatus struct {
	message string
}

const (
    StatusUnspecified ClusterStatus = iota
    Provisioning
    Running
    Reconciling
    Stopping
    Error
    Degraded
)

type Cluster struct {
	Cluster  string
	Status   ClusterStatus
	Version  string
	Endpoint string
}

type Nodepool struct {
	Nodepool    string
	Status      string
	Version     string
	NodeCount  int
}
