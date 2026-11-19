package files

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

// Download a file the user attached to a conversation, including while the conversation is still active.
func (c *Client) Get(ctx context.Context, conversationId string, fileId string, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Get(ctx, conversationId, fileId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Remove a file upload from a conversation. Only possible if the file hasn't already been used in the conversation.
func (c *Client) Delete(ctx context.Context, fileId string, conversationId string, opts ...option.RequestOption) (*types.ConvAiFileUploadResponseModel, error) {
	response, err := c.WithRawResponse.Delete(ctx, fileId, conversationId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Upload an image, PDF, or text file for a conversation. Returns a unique file ID that can be used to reference the file in the conversation.
func (c *Client) Create(ctx context.Context, conversationId string, request *conversations.BodyUploadFileV1ConvaiConversationsConversationIdFilesPost, opts ...option.RequestOption) (*types.ConvAiFileUploadResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, conversationId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
