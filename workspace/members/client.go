package members

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	workspace "github.com/elevenlabs/elevenlabs-go/workspace"
)

type Client struct {
	WithRawResponse *RawClient

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
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

// Gets a list of all members of the workspace, including locked members. Service accounts are excluded. Requires the workspace_members_read permission.
func (c *Client) List(ctx context.Context, opts ...option.RequestOption) ([]*types.WorkspaceMemberResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Updates attributes of a workspace member. Apart from the email identifier, all parameters will remain unchanged unless specified. This endpoint may only be called by workspace administrators.
func (c *Client) Update(ctx context.Context, request *workspace.UpdateMemberRequest, opts ...option.RequestOption) (*types.UpdateWorkspaceMemberResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Sets or clears the monthly credit usage limit of a workspace member or service account. This endpoint may only be called by workspace administrators.
func (c *Client) UpdateUsageLimit(ctx context.Context, userId string, request *workspace.UpdateMemberUsageLimitRequestModel, opts ...option.RequestOption) (*types.UpdateMemberUsageLimitResponseModel, error) {
	response, err := c.WithRawResponse.UpdateUsageLimit(ctx, userId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
