package client

import (
	req "github.com/levigross/grequests"
	"github.com/partitio/gonextcloud/client/types"
	"net/url"
)

type Client struct {
	baseURL      *url.URL
	username     string
	password     string
	session      *req.Session
	headers      map[string]string
	capabilities *types.Capabilities
}

func NewClient(hostname string) (*Client, error) {
	baseURL, err := url.Parse(hostname)
	if err != nil {
		return nil, err
	}
	c := Client{
		baseURL: baseURL,
		headers: map[string]string{
			"OCS-APIREQUEST": "true",
			"Accept":         "application/json",
		},
	}
	return &c, nil
}
