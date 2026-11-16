package members

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	groups "github.com/elevenlabs/elevenlabs-go/workspace/groups"
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

// Removes a member from the specified group. Requires `group_members_manage` permission.
func (c *Client) Remove(ctx context.Context, groupId string, request *groups.BodyDeleteMemberFromUserGroupV1WorkspaceGroupsGroupIdMembersRemovePost, opts ...option.RequestOption) (*types.DeleteWorkspaceGroupMemberResponseModel, error) {
	response, err := c.WithRawResponse.Remove(ctx, groupId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Adds a member of your workspace to the specified group. Requires `group_members_manage` permission.
func (c *Client) Add(ctx context.Context, groupId string, request *groups.AddMemberToGroupRequest, opts ...option.RequestOption) (*types.AddWorkspaceGroupMemberResponseModel, error) {
	response, err := c.WithRawResponse.Add(ctx, groupId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
