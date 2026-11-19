package client

import (
	context "context"
	agents "github.com/elevenlabs/elevenlabs-go/agents"
	dependentagents "github.com/elevenlabs/elevenlabs-go/agents/analysisitems/dependentagents"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	DependentAgents *dependentagents.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		DependentAgents: dependentagents.NewClient(options),
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

// List analysis items available in the workspace.
func (c *Client) List(ctx context.Context, request *agents.ListAnalysisItemsRequest, opts ...option.RequestOption) (*types.AnalysisItemsResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new analysis item (evaluation criterion or data collection) in the workspace.
func (c *Client) Create(ctx context.Context, request *agents.CreateAnalysisItemsRequestBody, opts ...option.RequestOption) (*types.AnalysisItemResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Attach or detach one analysis item on multiple agent branches.
func (c *Client) UpdateAttachments(ctx context.Context, analysisItemId string, request *agents.UpdateAnalysisItemAttachmentsRequestModel, opts ...option.RequestOption) (*types.UpdateAnalysisItemAttachmentsResponseModel, error) {
	response, err := c.WithRawResponse.UpdateAttachments(ctx, analysisItemId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get an analysis item available in the workspace.
func (c *Client) Get(ctx context.Context, analysisItemId string, opts ...option.RequestOption) (*types.AnalysisItemResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, analysisItemId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update an analysis item. Publishes a new version when the content changes.
func (c *Client) Update(ctx context.Context, analysisItemId string, request *agents.UpdateAnalysisItemsRequestBody, opts ...option.RequestOption) (*types.AnalysisItemResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, analysisItemId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete an analysis item and all of its versions from the workspace.
func (c *Client) Delete(ctx context.Context, analysisItemId string, request *agents.DeleteAnalysisItemsRequest, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Delete(ctx, analysisItemId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// List the published versions of an analysis item, newest first.
func (c *Client) ListVersions(ctx context.Context, analysisItemId string, request *agents.ListVersionsAnalysisItemsRequest, opts ...option.RequestOption) (*types.AnalysisItemVersionsResponseModel, error) {
	response, err := c.WithRawResponse.ListVersions(ctx, analysisItemId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new analysis item from this item's current definition.
func (c *Client) Copy(ctx context.Context, analysisItemId string, opts ...option.RequestOption) (*types.AnalysisItemResponseModel, error) {
	response, err := c.WithRawResponse.Copy(ctx, analysisItemId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
