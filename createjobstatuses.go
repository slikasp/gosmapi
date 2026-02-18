package gosmapi

import (
	"context"
	"fmt"
)

// GET

func (c *Client) Createjobstatuses(ctx context.Context) ([]Createjobstatus, error) {
	var output multiCreatejobstatusOutput

	err := c.makeRequest(
		ctx, GetRequest,
		CreatejobstatusesEndpoint,
		nil,
		nil,
		nil, &output)
	if err != nil {
		return nil, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

func (c *Client) Createjobstatus(ctx context.Context, createjobstatusID string) (Createjobstatus, error) {
	var output singleCreatejobstatusOutput

	err := c.makeRequest(
		ctx, GetRequest,
		CreatejobstatusesEndpoint,
		[]string{createjobstatusID},
		nil,
		nil, &output)
	if err != nil {
		return Createjobstatus{}, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}
