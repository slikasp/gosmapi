package gosmapi

import (
	"context"
	"fmt"
)

// TODO: maybe use the below for setting up proxies?
// not a fan of the constructor
// Maybe just not expose the type via funtions?

// type Proxy struct {
//     Object  // All Object fields promoted
//     // Add proxy-specific fields
// }

// func NewProxy(id string) Proxy {
//     return Proxy{
//         Object: Object{ID: id, Type: "proxies"},
//     }
// }

// GET

func (c *Client) Proxies(ctx context.Context) ([]Proxy, error) {
	var output multipleProxyOutput

	err := c.makeRequest(
		ctx, GetRequest,
		ProxiesEndpoint,
		nil,
		nil,
		nil, &output)
	if err != nil {
		return nil, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

func (c *Client) Proxy(ctx context.Context, proxyID string) (Proxy, error) {
	var output singleProxyOutput

	err := c.makeRequest(
		ctx, GetRequest,
		ProxiesEndpoint,
		[]string{proxyID},
		nil,
		nil, &output)
	if err != nil {
		return Proxy{}, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

func (c *Client) ProxySubservers(
	ctx context.Context,
	proxyID string,
) ([]Assigned, error) {
	var output AssignedOutput

	// This one for some reason works without /relationships
	err := c.makeRequest(
		ctx, GetRequest,
		ProxiesEndpoint,
		[]string{proxyID, string(AssignedSubserversElement)},
		nil,
		nil, &output)
	if err != nil {
		return nil, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

// POST

// PATCH

// offline proxies?

// DELETE
