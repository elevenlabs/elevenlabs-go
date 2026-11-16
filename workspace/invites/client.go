package invites

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

// Sends an email invitation to join your workspace to the provided email. If the user doesn't have an account they will be prompted to create one. If the user accepts this invite they will be added as a user to your workspace and your subscription using one of your seats. This endpoint may only be called by workspace members with the WORKSPACE_MEMBERS_INVITE permission. If the user is already in the workspace a 400 error will be returned.
func (c *Client) Create(ctx context.Context, request *workspace.InviteUserRequest, opts ...option.RequestOption) (*types.AddWorkspaceInviteResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Sends email invitations to join your workspace to the provided emails. Requires all email addresses to be part of a verified domain. If the users don't have an account they will be prompted to create one. If the users accept these invites they will be added as users to your workspace and your subscription using one of your seats. This endpoint may only be called by workspace members with the WORKSPACE_MEMBERS_INVITE permission.
func (c *Client) CreateBatch(ctx context.Context, request *workspace.BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost, opts ...option.RequestOption) (*types.AddWorkspaceInviteResponseModel, error) {
	response, err := c.WithRawResponse.CreateBatch(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Invalidates an existing email invitation. The invitation will still show up in the inbox it has been delivered to, but activating it to join the workspace won't work. This endpoint may only be called by workspace members with the WORKSPACE_MEMBERS_INVITE permission.
func (c *Client) Delete(ctx context.Context, request *workspace.BodyDeleteExistingInvitationV1WorkspaceInvitesDelete, opts ...option.RequestOption) (*types.DeleteWorkspaceInviteResponseModel, error) {
	response, err := c.WithRawResponse.Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
