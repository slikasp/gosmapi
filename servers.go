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

	elems := []string{"servers"}
	params := map[string]string{}
	path := buildRequestPath(elems, params)
	err := c.makeRequest(ctx, "GET", path, nil, &output)
	if err != nil {
		return []FileServer{}, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

func (c *Client) FileServer(ctx context.Context, serverID string) (FileServer, error) {
	var output SingleFileServerOutput

	elems := []string{"servers", serverID}
	params := map[string]string{"include": "subServers"}
	path := buildRequestPath(elems, params)
	err := c.makeRequest(ctx, "GET", path, nil, &output)
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

	elems := []string{"servers"}
	params := map[string]string{"include": "subServers"}
	path := buildRequestPath(elems, params)

	err := c.makeRequest(ctx, "POST", path, input, &output)
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
	attributes T,
	serverID string,
) (FileServer, error) {

	var input editFileServerInput[T]
	input.Data.Attributes = attributes

	var output SingleFileServerOutput

	//request
	elems := []string{"servers", serverID}
	params := map[string]string{}
	path := buildRequestPath(elems, params)
	err := c.makeRequest(ctx, "PATCH", path, input, &output)
	if err != nil {
		return FileServer{}, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

func (c *Client) EditOtherFileServer(
	ctx context.Context,
	attributes EditOtherFileServerAttributes,
	serverID string,
) (FileServer, error) {
	return editFileServer(c, ctx, attributes, serverID)
}

func (c *Client) EditIntegratedFileServer(
	ctx context.Context,
	attributes EditIntegratedFileServerAttributes,
	serverID string,
) (FileServer, error) {
	return editFileServer(c, ctx, attributes, serverID)
}

// DELETE
