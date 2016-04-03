package cluster

import (
	"crypto/x509"
	"errors"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"launchpad.net/gwacl/fork/http"
	"launchpad.net/gwacl/fork/tls"
)

type PfxFileAuth struct {
	PfxFile     string
	PfxPasscode string
}

func (c *PfxFileAuth) Auth() (*http.Transport, error) {
	if c.PfxFile == "" {
		return &http.Transport{}, errors.New("No PFX file specified")
	}

	data, err := ioutil.ReadFile(c.PfxFile)
	if err != nil {
		return &http.Transport{}, err
	}

	privateKey, cert, err := pkcs12.Decode(data, c.PfxPasscode)
	if err != nil {
		return &http.Transport{}, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AddCert(cert)

	tlsCert := tls.Certificate{
		Certificate: [][]byte{cert.Raw},
		PrivateKey:  privateKey,
		Leaf:        cert,
	}

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{tlsCert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: true,
	}

	tlsConfig.BuildNameToCertificate()
	return &http.Transport{TLSClientConfig: tlsConfig}, nil
}
