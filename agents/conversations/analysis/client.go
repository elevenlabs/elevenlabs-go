package analysis

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

// Run the analysis for a conversation using the agent's current evaluation criteria and data collection settings.
func (c *Client) Run(ctx context.Context, conversationId string, opts ...option.RequestOption) (*types.GetConversationResponseModel, error) {
	response, err := c.WithRawResponse.Run(ctx, conversationId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Rerun a specific evaluation for a conversation.
func (c *Client) RunEvaluation(ctx context.Context, conversationId string, request *conversations.RunConversationEvaluationsRequest, opts ...option.RequestOption) (*types.GetConversationResponseModel, error) {
	response, err := c.WithRawResponse.RunEvaluation(ctx, conversationId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
