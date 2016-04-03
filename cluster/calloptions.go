package cluster

import (
	"launchpad.net/gwacl/fork/http"
)

type Endpointer interface {
	Endpoint() string
}

type Auther interface {
	Auth() (*http.Transport, error)
}

type CallOptions struct {
	BaseUrl     string
	EndpointUrl string
}

func (c *CallOptions) Endpoint() string {
	return c.BaseUrl + c.EndpointUrl
}
