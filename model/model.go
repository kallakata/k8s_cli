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

type Cluster struct {
	Cluster  string
	Status   string
	Version  string
	Endpoint string
}

type Nodepool struct {
	Nodepool    string
	Status      string
	Version     string
	MinNode     int32
	MaxNode     int32
	Autoscaling bool
}
