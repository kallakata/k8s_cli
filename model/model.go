package model

type Ns struct {
    Namespace string
}

type Pod struct {
    Pod string
	Status string
    Namespace string
    Context string
}

type Cluster struct {
    Cluster string
	Status string
    Version string
    Endpoint string
}