package gosmapi

import (
	"context"
)

// GET

// Returns a list of file servers already configured on the core in Servers struct
//
// TODO: add support for fields, limit, sort
func (c *Client) Subservers(ctx context.Context) ([]Subserver, error) {
	var output MultiFileSubserverOutput

	elems := []string{"subservers"}
	params := map[string]string{}
	path := buildRequestPath(elems, params)
	err := c.makeRequest(ctx, "GET", path, nil, &output)

	return output.Data, err
}

func (c *Client) Subserver(ctx context.Context, subserverID string) (Subserver, error) {
	var output SingleFileSubserverOutput

	elems := []string{"subservers", subserverID}
	params := map[string]string{}
	path := buildRequestPath(elems, params)
	err := c.makeRequest(ctx, "GET", path, nil, &output)

	return output.Data, err
}

func (c *Client) SubserverParent(ctx context.Context, subserverID string) (FileServer, error) {
	var output SingleFileServerOutput

	elems := []string{"subservers", subserverID, "server"}
	params := map[string]string{}
	path := buildRequestPath(elems, params)
	err := c.makeRequest(ctx, "GET", path, nil, &output)

	return output.Data, err
}

// POST

// PATCH

// TODO: all shares/exports are set to Automatic by default, test they are modified
// TODO: Need my own Isilon for testing
func (c *Client) EditSubserverConnection(
	ctx context.Context,
	connection ConnectionConfig,
	subserverID string,
) (Subserver, error) {
	var input editSubserverConnectionInput
	input.Data.Attributes.ConnectionConfig = connection
	var output SingleFileSubserverOutput

	elems := []string{"subservers", subserverID}
	params := map[string]string{}
	path := buildRequestPath(elems, params)
	err := c.makeRequest(ctx, "PATCH", path, input, &output)

	return output.Data, err
}

// TODO: create lower level struct
// TODO: check if configured shares are modified
// TODO: how to set all to none (use the list provided after creation)
func (c *Client) EditDataAccess(
	ctx context.Context,
	dataAccess EditSubserverDataAccess,
	subserverID string,
) (Subserver, error) {
	var input editSubserverDataAccessInput
	input.Data.Attributes.DataAccess = dataAccess
	var output SingleFileSubserverOutput

	elems := []string{"subservers", subserverID}
	params := map[string]string{}
	path := buildRequestPath(elems, params)
	err := c.makeRequest(ctx, "PATCH", path, input, &output)

	return output.Data, err
}

// ASSIGNED PROXIES

// Wrapper for outer proxy assignment functions
// Currently has no output
func assignedProxies(
	c *Client,
	ctx context.Context,
	subserverID string,
	proxyIDs []string,
	request RequestType,
) error {
	var input AssignedProxiesInput
	input.Data = make([]AssignedProxy, 0, len(proxyIDs))
	for _, id := range proxyIDs {
		input.Data = append(input.Data, AssignedProxy{
			ID:   id,
			Type: ProxyObject,
		})
	}
	var output struct{}

	elems := []string{"subservers", subserverID, "assignedProxies"}
	params := map[string]string{}
	path := buildRequestPath(elems, params)
	err := c.makeRequest(ctx, string(request), path, input, &output)

	return err
}

// GET

func (c *Client) SubserverProxies(
	ctx context.Context,
	subserverID string,
) ([]AssignedProxy, error) {
	var output AssignedProxiesOutput

	elems := []string{"subservers", subserverID, "assignedProxies"}
	params := map[string]string{}
	path := buildRequestPath(elems, params)
	err := c.makeRequest(ctx, "GET", path, nil, &output)

	return output.Data, err
}

// POST

func (c *Client) SubserverAddProxies(
	ctx context.Context,
	subserverID string,
	proxyIDs []string,
) error {
	return assignedProxies(c, ctx, subserverID, proxyIDs, PostRequest)
}

// PATCH

func (c *Client) SubserverSetProxies(
	ctx context.Context,
	subserverID string,
	proxyIDs []string,
) error {
	return assignedProxies(c, ctx, subserverID, proxyIDs, PatchRequest)
}

// DELETE

func (c *Client) SubserverRemoveProxies(
	ctx context.Context,
	subserverID string,
	proxyIDs []string,
) error {
	return assignedProxies(c, ctx, subserverID, proxyIDs, PostRequest)
}
