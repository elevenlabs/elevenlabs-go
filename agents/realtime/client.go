package realtime

import (
	context "context"
	agents "github.com/elevenlabs/elevenlabs-go/agents"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
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

// List the realtime conversation groups the current user can view.
func (c *Client) ListConversationGroups(ctx context.Context, request *agents.ListConversationGroupsRealtimeRequest, opts ...option.RequestOption) (*types.RealtimeConversationGroupsResponseModel, error) {
	response, err := c.WithRawResponse.ListConversationGroups(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a realtime conversation group owned by the current user.
func (c *Client) CreateConversationGroup(ctx context.Context, request *agents.CreateRealtimeConversationGroupRequestModel, opts ...option.RequestOption) (*types.RealtimeConversationGroupResponseModel, error) {
	response, err := c.WithRawResponse.CreateConversationGroup(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a realtime conversation group by id.
func (c *Client) GetConversationGroup(ctx context.Context, conversationGroupId string, opts ...option.RequestOption) (*types.RealtimeConversationGroupResponseModel, error) {
	response, err := c.WithRawResponse.GetConversationGroup(ctx, conversationGroupId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Rename or re-describe a realtime conversation group.
func (c *Client) UpdateConversationGroup(ctx context.Context, conversationGroupId string, request *agents.UpdateRealtimeConversationGroupRequestModel, opts ...option.RequestOption) (*types.RealtimeConversationGroupResponseModel, error) {
	response, err := c.WithRawResponse.UpdateConversationGroup(ctx, conversationGroupId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a realtime conversation group. Every conversation in it becomes ungrouped, so access granted through the group is revoked.
func (c *Client) DeleteConversationGroup(ctx context.Context, conversationGroupId string, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.DeleteConversationGroup(ctx, conversationGroupId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Place a realtime conversation in a conversation group, moving it out of its current group if it has one.
func (c *Client) SetConversationGroup(ctx context.Context, conversationId string, request *agents.SetConversationGroupRequestModel, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.SetConversationGroup(ctx, conversationId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Remove a realtime conversation from its conversation group.
func (c *Client) ClearConversationGroup(ctx context.Context, conversationId string, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.ClearConversationGroup(ctx, conversationId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create an ephemeral client secret for authenticating browser-based Realtime API sessions.
func (c *Client) CreateClientSecret(ctx context.Context, request *agents.CreateClientSecretRequest, opts ...option.RequestOption) (*types.ClientSecretResponse, error) {
	response, err := c.WithRawResponse.CreateClientSecret(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
