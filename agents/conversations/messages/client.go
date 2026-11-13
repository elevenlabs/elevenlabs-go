package messages

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

// Search through conversation transcript messages by full-text and fuzzy search
func (c *Client) TextSearch(ctx context.Context, request *conversations.TextSearchMessagesRequest, opts ...option.RequestOption) (*types.MessagesSearchResponse, error) {
	response, err := c.WithRawResponse.TextSearch(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Search conversation transcripts by semantic similarity to surface relevant messages based on meaning and intent, rather than exact keyword matches
func (c *Client) Search(ctx context.Context, request *conversations.SearchMessagesRequest, opts ...option.RequestOption) (*types.MessagesSearchResponse, error) {
	response, err := c.WithRawResponse.Search(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
