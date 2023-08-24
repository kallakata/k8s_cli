package model

type Ns struct {
	Namespace string
}

type Pod struct {
	Pod       string
	Status    string
	Namespace string
	Context   string
}

type ClusterStatus int

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
	Autoscaling bool
	MinNode     int32
	MaxNode     int32
}
