package model

type Ns struct {
	Namespace string
	Pods      int
}

type Pod struct {
	Pod       string
	Status    string
	Namespace string
	CPUReq    string
	MemReq    string
	CPULim    string
	MemLim    string
	Image     string
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
	Nodepool  string
	Status    string
	Version   string
	NodeCount int
}

type AzureCluster struct {
	Cluster  string `json:"cluster,omitempty"`
	Version  string `json:"version,omitempty"`
	Status   string `json:"status,omitempty"`
	Location string `json:"location,omitempty"`
	Identity string `json:"identity,omitempty"`
}
