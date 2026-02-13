package gosmapi

import (
	"context"
	"fmt"
)

// GET

// Returns a list of file servers already configured on the core in Servers struct
//
// TODO: add support for fields, limit, sort
func (c *Client) FileServers(ctx context.Context) ([]FileServer, error) {
	var output MultiFileServerOutput

	err := c.makeRequest(
		ctx, GetRequest,
		ServersEndpoint,
		nil,
		nil,
		nil, &output)
	if err != nil {
		return []FileServer{}, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

func (c *Client) FileServer(ctx context.Context, serverID string) (FileServer, error) {
	var output SingleFileServerOutput

	err := c.makeRequest(
		ctx, GetRequest,
		ServersEndpoint,
		[]string{serverID},
		map[string]string{"include": "subServers"},
		nil, &output)
	if err != nil {
		return FileServer{}, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

// POST

// Wrapper for repeating exposed functions
func addFileServer[T AddFileServerAttributes](
	c *Client,
	ctx context.Context,
	attributes T,
) (FileServer, error) {
	var input addFileServerInput[T]
	input.Data.Attributes = attributes
	var output SingleFileServerOutput

	err := c.makeRequest(
		ctx, PostRequest,
		ServersEndpoint,
		nil,
		map[string]string{"include": "subServers"},
		input, &output)
	if err != nil {
		return FileServer{}, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, nil
}

func (c *Client) AddOtherFileServer(
	ctx context.Context,
	attributes AddOtherFileServerAttributes,
) (FileServer, error) {
	return addFileServer(c, ctx, attributes)
}

func (c *Client) AddIntegratedFileServer(
	ctx context.Context,
	attributes AddIntegratedFileServerAttributes,
) (FileServer, error) {
	return addFileServer(c, ctx, attributes)
}

// PATCH
func editFileServer[T EditFileServerAttributes](
	c *Client,
	ctx context.Context,
	serverID string,
	attributes T,
) (FileServer, error) {
	var input editFileServerInput[T]
	input.Data.Attributes = attributes
	var output SingleFileServerOutput

	err := c.makeRequest(
		ctx, PatchRequest,
		ServersEndpoint,
		[]string{serverID},
		nil,
		input, &output)
	if err != nil {
		return FileServer{}, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

func (c *Client) EditOtherFileServer(
	ctx context.Context,
	serverID string,
	attributes EditOtherFileServerAttributes,
) (FileServer, error) {
	return editFileServer(c, ctx, serverID, attributes)
}

func (c *Client) EditIntegratedFileServer(
	ctx context.Context,
	serverID string,
	attributes EditIntegratedFileServerAttributes,
) (FileServer, error) {
	return editFileServer(c, ctx, serverID, attributes)
}

// DELETE
