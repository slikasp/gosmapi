package gosmapi

import (
	"context"
	"fmt"
)

// GET

// Returns a list of file servers already configured on the core in Servers struct
//
// TODO: add support for fields, limit, sort
func (c *Client) Subservers(ctx context.Context) ([]Subserver, error) {
	var output MultiFileSubserverOutput

	err := c.makeRequest(
		ctx, GetRequest,
		SubserversEndpoint,
		nil,
		nil,
		nil, &output)
	if err != nil {
		return nil, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

func (c *Client) Subserver(ctx context.Context, subserverID string) (Subserver, error) {
	var output SingleFileSubserverOutput

	err := c.makeRequest(
		ctx, GetRequest,
		SubserversEndpoint,
		[]string{subserverID},
		nil,
		nil, &output)
	if err != nil {
		return Subserver{}, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

func (c *Client) SubserverParent(ctx context.Context, subserverID string) (FileServer, error) {
	var output SingleFileServerOutput

	err := c.makeRequest(
		ctx, GetRequest,
		SubserversEndpoint,
		[]string{subserverID, "server"},
		nil,
		nil, &output)
	if err != nil {
		return FileServer{}, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

func (c *Client) SubserverProxies(
	ctx context.Context,
	subserverID string,
) ([]Assigned, error) {
	var output AssignedOutput

	// This one for some reason works without /relationships
	err := c.makeRequest(
		ctx, GetRequest,
		SubserversEndpoint,
		[]string{subserverID, string(AssignedProxiesElement)},
		nil,
		nil, &output)
	if err != nil {
		return nil, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

// POST

// PATCH

// TODO: all shares/exports are set to Automatic by default, test if they are modified
func (c *Client) EditSubserverConnection(
	ctx context.Context,
	subserverID string,
	connection DataConnection,
) (Subserver, error) {
	var input editSubserverConnectionInput
	input.Data.Attributes.ConnectionConfig = connection
	var output SingleFileSubserverOutput

	err := c.makeRequest(
		ctx,
		PatchRequest,
		SubserversEndpoint,
		[]string{subserverID},
		nil,
		input, &output)
	if err != nil {
		return Subserver{}, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

// TODO: create lower level struct
// TODO: check if configured shares are modified
// TODO: how to set all to none (use the list provided after creation)
func (c *Client) EditDataAccess(
	ctx context.Context,
	subserverID string,
	dataAccess EditSubserverDataAccess,
) (Subserver, error) {
	var input editSubserverDataAccessInput
	input.Data.Attributes.DataAccess = dataAccess
	var output SingleFileSubserverOutput

	err := c.makeRequest(
		ctx, PatchRequest,
		SubserversEndpoint,
		[]string{subserverID},
		nil,
		input, &output)
	if err != nil {
		return Subserver{}, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}
