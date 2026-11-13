package tags

import (
	context "context"
	conversations "github.com/elevenlabs/elevenlabs-go/agents/conversations"
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

// Assign one or more conversation tags to a conversation. Tags that are already assigned are ignored. Tags must belong to the same workspace.
func (c *Client) Assign(ctx context.Context, conversationId string, request *conversations.AssignConversationTagsRequestModel, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Assign(ctx, conversationId, request, opts...)
	return err
}

// Remove a single conversation tag from a conversation.
func (c *Client) Unassign(ctx context.Context, conversationId string, tagId string, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Unassign(ctx, conversationId, tagId, opts...)
	return err
}

// List conversation tags for the workspace, ordered by most recently created first.
func (c *Client) List(ctx context.Context, request *conversations.ListTagsRequest, opts ...option.RequestOption) (*types.GetConversationTagsPageResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new conversation tag for the workspace.
func (c *Client) Create(ctx context.Context, request *conversations.CreateConversationTagRequestModel, opts ...option.RequestOption) (*types.ConversationTagResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a conversation tag by ID.
func (c *Client) Get(ctx context.Context, tagId string, opts ...option.RequestOption) (*types.ConversationTagResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, tagId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update a conversation tag's title and/or description. Restricted to the tag owner or a workspace admin.
func (c *Client) Update(ctx context.Context, tagId string, request *conversations.PatchConversationTagRequestModel, opts ...option.RequestOption) (*types.ConversationTagResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, tagId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a conversation tag. Restricted to the tag owner or a workspace admin.
func (c *Client) Delete(ctx context.Context, tagId string, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Delete(ctx, tagId, opts...)
	return err
}
