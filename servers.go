package gosmapi

import (
	"context"
)

func (c *Client) FileServersGet(ctx context.Context, serverID string) (Server, error) {
	var response Server

	elems := []string{"servers", serverID}
	params := map[string]string{"include": "subServers"}
	path := buildRequestPath(elems, params)
	err := c.makeRequest(ctx, "GET", path, nil, &response)

	return response, err
}

// Returns a list of file servers already configured on the core in Servers struct
//
// TODO: add support for fields, limit, sort
func (c *Client) FileServersGetAll(ctx context.Context) (Servers, error) {
	var response Servers

	elems := []string{"servers"}
	params := map[string]string{}
	path := buildRequestPath(elems, params)
	err := c.makeRequest(ctx, "GET", path, nil, &response)

	return response, err
}

func (c *Client) FileServerAddOther(ctx context.Context, input FileServerAddOtherRequest) (Server, error) {
	var response Server

	elems := []string{"servers"}
	// even OtherNas servers which don't have subServers in the UI, have a fake subserver which allows share/export/proxy config
	params := map[string]string{"include": "subServers"}
	path := buildRequestPath(elems, params)
	err := c.makeRequest(ctx, "POST", path, input, &response)

	return response, err
}

func (c *Client) FileServerEditOther(ctx context.Context, input FileServerEditOtherRequest, serverID string) (Server, error) {
	var response Server

	elems := []string{"servers", serverID}
	params := map[string]string{}
	path := buildRequestPath(elems, params)
	err := c.makeRequest(ctx, "PATCH", path, input, &response)

	return response, err
}

func (c *Client) FileServerAddIntegrated(ctx context.Context, input FileServerAddIntegratedRequest) (Server, error) {
	var response Server

	elems := []string{"servers"}
	params := map[string]string{"include": "subServers"}
	path := buildRequestPath(elems, params)
	err := c.makeRequest(ctx, "POST", path, input, &response)

	return response, err
}

func (c *Client) FileServerEditIntegrated(ctx context.Context, input FileServerEditIntegratedRequest, serverID string) (Server, error) {
	var response Server

	elems := []string{"servers", serverID}
	params := map[string]string{}
	path := buildRequestPath(elems, params)
	err := c.makeRequest(ctx, "PATCH", path, input, &response)

	return response, err
}
