package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	workspace "github.com/elevenlabs/elevenlabs-go/workspace"
	members "github.com/elevenlabs/elevenlabs-go/workspace/groups/members"
)

type Client struct {
	WithRawResponse *RawClient

	Members *members.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Members:         members.NewClient(options),
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

// Get all groups in the workspace
func (c *Client) List(ctx context.Context, opts ...option.RequestOption) (map[string]*types.WorkspaceGroupResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Searches for user groups in the workspace. Multiple or no groups may be returned.
func (c *Client) Search(ctx context.Context, request *workspace.SearchGroupsRequest, opts ...option.RequestOption) ([]*types.WorkspaceGroupByNameResponseModel, error) {
	response, err := c.WithRawResponse.Search(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
