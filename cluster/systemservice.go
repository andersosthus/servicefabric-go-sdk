package cluster

import (
	"encoding/json"
)

func GetSystemServices(auth Auther, options CallOptions) (services *[]SystemService, err error) {
	client := &SfClient{}
	resp, err := client.Call(auth, options)
	if err != nil {
		return nil, err
	}

	var result *[]SystemService
	err = json.Unmarshal(resp, &result)

	return result, err
}

type SystemService struct {
	Id                string
	ServiceKind       int
	Name              string
	TypeName          string
	ManifestVersion   string
	HasPersistedState bool
	HealthState       int
	ServiceStatus     int
	IsServiceGroup    bool
}
