package client

import (
	context "context"
	agents "github.com/elevenlabs/elevenlabs-go/agents"
	folders "github.com/elevenlabs/elevenlabs-go/agents/tests/folders"
	invocations "github.com/elevenlabs/elevenlabs-go/agents/tests/invocations"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Folders     *folders.Client
	Invocations *invocations.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Folders:         folders.NewClient(options),
		Invocations:     invocations.NewClient(options),
		options:         options,
		baseURL:         options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:         options.HTTPClient,
				MaxAttempts:    options.MaxAttempts,
				DisableRetries: options.DisableRetries,
			},
		),
	}
}

// Creates a new agent response test.
func (c *Client) Create(ctx context.Context, request *agents.CreateTestsRequestBody, opts ...option.RequestOption) (*types.CreateAgentTestResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Moves multiple tests or folders from one folder to another.
func (c *Client) Move(ctx context.Context, request *agents.BodyBulkMoveTestsToFolderV1ConvaiAgentTestingBulkMovePost, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Move(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Gets an agent response test by ID.
func (c *Client) Get(ctx context.Context, testId string, opts ...option.RequestOption) (*agents.GetTestsResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, testId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Updates an agent response test by ID.
func (c *Client) Update(ctx context.Context, testId string, request *agents.UpdateTestsRequestBody, opts ...option.RequestOption) (*agents.UpdateTestsResponse, error) {
	response, err := c.WithRawResponse.Update(ctx, testId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Deletes an agent response test by ID.
func (c *Client) Delete(ctx context.Context, testId string, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Delete(ctx, testId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Gets agent response test summaries for the requested test IDs.
func (c *Client) Summaries(ctx context.Context, request *agents.ListTestsByIdsRequestModel, opts ...option.RequestOption) (*types.GetTestsSummariesByIdsResponseModel, error) {
	response, err := c.WithRawResponse.Summaries(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Lists all agent response tests with pagination support and optional search filtering.
func (c *Client) List(ctx context.Context, request *agents.ListTestsRequest, opts ...option.RequestOption) (*types.GetTestsPageResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
