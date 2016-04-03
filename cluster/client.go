package cluster

import (
	"launchpad.net/gwacl/fork/http"
	"io/ioutil"
)

type SfClient struct {
}

func (c *SfClient) Call(auth Auther, options CallOptions) (responseData []byte, err error) {
	test := make([]byte, 0)
	transport, err := auth.Auth()
	if err != nil {
		return test, err
	}

	client := &http.Client{Transport: transport}

	url := options.Endpoint() + "?api-version=1.0"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return test, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return body, err
}
