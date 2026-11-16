package members

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	team "github.com/elevenlabs/elevenlabs-go/reception/team"
	types "github.com/elevenlabs/elevenlabs-go/types"
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

// List the workspace's members, suspended members last.
func (c *Client) List(ctx context.Context, request *team.ListMembersRequest, opts ...option.RequestOption) ([]*types.ReceptionTeamMemberResponse, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Promote a member to admin or demote an admin to member.
func (c *Client) UpdateRole(ctx context.Context, userId string, request *team.BodySmbUpdateTeamMemberRole, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.UpdateRole(ctx, userId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Remove a member from the workspace. Their seat goes to a suspended member, if any.
func (c *Client) Remove(ctx context.Context, userId string, request *team.RemoveMembersRequest, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.Remove(ctx, userId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
