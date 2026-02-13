package gosmapi

import (
	"context"
	"fmt"
)

// GET

func (c *Client) Jobs(ctx context.Context) ([]Job, error) {
	var output multipleJobOutput

	err := c.makeRequest(
		ctx, GetRequest,
		JobsEndpoint,
		nil,
		nil,
		nil, &output)
	if err != nil {
		return nil, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

func (c *Client) Job(ctx context.Context, jobID string) (Job, error) {
	var output singleJobOutput

	err := c.makeRequest(
		ctx, GetRequest,
		JobsEndpoint,
		[]string{jobID},
		nil,
		nil, &output)
	if err != nil {
		return Job{}, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

// POST

func (c *Client) CreateJob(
	ctx context.Context,
	sourceID, destinationID string,
	attributes CreateJobAttributes,
) (Createjobstatus, error) {
	var input createJob
	input.Data.Type = JobsObject
	input.Data.Attributes = attributes
	input.Data.Relationships.SourceSubserver.Data.Type = SubserversObject
	input.Data.Relationships.SourceSubserver.Data.ID = sourceID
	input.Data.Relationships.DestinationSubserver.Data.Type = SubserversObject
	input.Data.Relationships.DestinationSubserver.Data.ID = destinationID
	var output singleCreatejobstatusOutput

	err := c.makeRequest(
		ctx,
		PostRequest,
		JobsEndpoint,
		nil,
		nil,
		input, &output)
	if err != nil {
		return Createjobstatus{}, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

// creating a job returns createjobstatuses JSON
// using /api/createjobstatuses returns createjobstatus JSON if it fails, job JSON if it works, why?
// can verify by checking type

// PATCH

func (c *Client) EditJob(
	ctx context.Context,
	jobID string,
	options JobOptions,
) (Job, error) {
	var input editJob
	input.Data.Attributes.Options = options
	var output singleJobOutput

	err := c.makeRequest(
		ctx,
		PatchRequest,
		JobsEndpoint,
		nil,
		nil,
		input, &output)
	if err != nil {
		return Job{}, fmt.Errorf("Request failed: %w", err)
	}

	return output.Data, err
}

// DELETE
