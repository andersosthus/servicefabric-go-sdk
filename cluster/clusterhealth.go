package cluster

import (
	"encoding/json"
)

func GetClusterHealth(auth Auther, options CallOptions) (health *Health, err error) {
	client := &SfClient{}
	resp, err := client.Call(auth, options)
	if err != nil {
		return nil, err
	}

	var result *Health
	err = json.Unmarshal(resp, &result)

	return result, err
}

type Health struct {
	HealthState int
	NodeHealthStateChunks NodeHealthStateChunks
	ApplicationHealthStateChunks ApplicationHealthStateChunks
	ServiceHealthStateChunks ServiceHealthStateChunks
	PartitionHealthStateChunks PartitionHealthStateChunks
	ReplicaHealthStateChunks ReplicaHealthStateChunks
	DeployedApplicationHealthStateChunks DeployedApplicationHealthStateChunks
	DeployedServicePackageHealthStateChunks DeployedServicePackageHealthStateChunks

}

type NodeHealthStateChunks struct {
	TotalCount int
	Items []NodeHealthChunk
}

type ApplicationHealthStateChunks struct {
	TotalCount int
	Items []ApplicationHealthChunk
}

type ServiceHealthStateChunks struct {
	TotalCount int
	Items []ServiceHealthHealthChunk
}

type PartitionHealthStateChunks struct {
	TotalCount int
	Items []PartitionHealthHealthChunk
}

type ReplicaHealthStateChunks struct {
	TotalCount int
	Items []ReplicaHealthHealthChunk
}

type DeployedApplicationHealthStateChunks struct {
	TotalCount int
	Items []DeployedApplicationHealthHealthChunk
}

type DeployedServicePackageHealthStateChunks struct {
	TotalCount int
	Items []DeployedServicePackageHealthHealthChunk
}


type ApplicationHealthChunk struct {
	ApplicationName string
	HealthState int
}

type NodeHealthChunk struct {
	NodeName string
	HealthState int
}

type ServiceHealthHealthChunk struct {
	ServiceName string
	HealthState int
}

type PartitionHealthHealthChunk struct {
	PartitionId string
	HealthState int
}

type ReplicaHealthHealthChunk struct {
	ReplicaOrInstanceId string
	HealthState int
}

type DeployedApplicationHealthHealthChunk struct {
	NodeName string
	HealthState int
}

type DeployedServicePackageHealthHealthChunk struct {
	ServiceManifestName string
	HealthState int
}