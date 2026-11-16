package resources

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

// Gets the metadata of a resource by ID.
func (c *Client) Get(ctx context.Context, resourceId string, request *workspace.GetResourcesRequest, opts ...option.RequestOption) (*types.ResourceMetadataResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, resourceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Grants a role (one of 'admin', 'editor', 'commenter', or 'viewer') on a workspace resource to a user, group, or workspace (service account) API key. This overrides any existing role the target has on the resource. To target a user or service account, pass only the user email; the user must be in your workspace. To target a group, pass only the group id. To target a workspace (service account) API key, pass the api key id; the resource will be shared with the service account associated with that key. You must have admin access to the resource to share it.
func (c *Client) Share(ctx context.Context, resourceId string, request *workspace.BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Share(ctx, resourceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Removes any existing role on a workspace resource from a user, group, or workspace (service account) API key. To target a user or service account, pass only the user email; the user must be in your workspace. To target a group, pass only the group id. To target a workspace (service account) API key, pass the api key id; the resource will be unshared from the service account associated with that key. You must have admin access to the resource to unshare it. You cannot remove permissions from the user who created the resource.
func (c *Client) Unshare(ctx context.Context, resourceId string, request *workspace.BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Unshare(ctx, resourceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
