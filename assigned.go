package gosmapi

import (
	"context"
	"fmt"
)

// GENERIC

// Generic wrapper for outer assignment functions (proxies, subservers)
// Currently these requests have no body (HTTP 204)
func assign(
	c *Client, ctx context.Context, request RequestType,
	endpoint Endpoint,
	endpointID string,
	assign AssignElement,
	input AssignedInput,
) error {
	err := c.makeRequest(
		ctx, request,
		endpoint,
		[]string{endpointID, "relationships", string(assign)},
		nil,
		input, nil)
	if err != nil {
		return fmt.Errorf("Request failed: %w", err)
	}

	return err
}

// PROXIES

func assigProxies(
	c *Client,
	ctx context.Context,
	subserverID string,
	proxyIDs []string,
	request RequestType,
) error {
	var input AssignedInput
	input.Data = make([]Assigned, 0, len(proxyIDs))
	for _, id := range proxyIDs {
		input.Data = append(input.Data, Assigned{
			ID:   id,
			Type: ProxiesObject,
		})
	}

	// Assign Proxies(element, input) to Subservers(endpoint, ID)
	return assign(
		c, ctx, request,
		SubserversEndpoint, subserverID,
		AssignedProxiesElement, input,
	)
}

// SUBSERVERS

func assigSubservers(
	c *Client,
	ctx context.Context,
	proxyID string,
	subserverIDs []string,
	request RequestType,
) error {
	var input AssignedInput
	input.Data = make([]Assigned, 0, len(subserverIDs))
	for _, id := range subserverIDs {
		input.Data = append(input.Data, Assigned{
			ID:   id,
			Type: SubserversObject,
		})
	}

	// Assign Subservers(element, input) to Proxies(endpoint, ID)
	return assign(
		c, ctx, request,
		ProxiesEndpoint, proxyID,
		AssignedSubserversElement, input,
	)
}

// POST

func (c *Client) SubserverAddProxies(
	ctx context.Context,
	subserverID string,
	proxyIDs []string,
) error {
	return assigProxies(c, ctx, subserverID, proxyIDs, PostRequest)
}

func (c *Client) ProxyAddSubservers(
	ctx context.Context,
	proxyID string,
	subserverIDs []string,
) error {
	return assigSubservers(c, ctx, proxyID, subserverIDs, PostRequest)
}

// PATCH

func (c *Client) SubserverSetProxies(
	ctx context.Context,
	subserverID string,
	proxyIDs []string,
) error {
	return assigProxies(c, ctx, subserverID, proxyIDs, PatchRequest)
}

func (c *Client) ProxySetSubservers(
	ctx context.Context,
	proxyID string,
	subserverIDs []string,
) error {
	return assigSubservers(c, ctx, proxyID, subserverIDs, PatchRequest)
}

// DELETE

func (c *Client) SubserverRemoveProxies(
	ctx context.Context,
	subserverID string,
	proxyIDs []string,
) error {
	return assigProxies(c, ctx, subserverID, proxyIDs, DeleteRequest)
}

func (c *Client) ProxyRemoveSubservers(
	ctx context.Context,
	proxyID string,
	subserverIDs []string,
) error {
	return assigSubservers(c, ctx, proxyID, subserverIDs, DeleteRequest)
}
