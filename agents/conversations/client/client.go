package client

import (
	context "context"
	agents "github.com/elevenlabs/elevenlabs-go/agents"
	analysis "github.com/elevenlabs/elevenlabs-go/agents/conversations/analysis"
	audio "github.com/elevenlabs/elevenlabs-go/agents/conversations/audio"
	feedback "github.com/elevenlabs/elevenlabs-go/agents/conversations/feedback"
	files "github.com/elevenlabs/elevenlabs-go/agents/conversations/files"
	messages "github.com/elevenlabs/elevenlabs-go/agents/conversations/messages"
	tags "github.com/elevenlabs/elevenlabs-go/agents/conversations/tags"
	topics "github.com/elevenlabs/elevenlabs-go/agents/conversations/topics"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Audio    *audio.Client
	Files    *files.Client
	Feedback *feedback.Client
	Messages *messages.Client
	Tags     *tags.Client
	Topics   *topics.Client
	Analysis *analysis.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Audio:           audio.NewClient(options),
		Files:           files.NewClient(options),
		Feedback:        feedback.NewClient(options),
		Messages:        messages.NewClient(options),
		Tags:            tags.NewClient(options),
		Topics:          topics.NewClient(options),
		Analysis:        analysis.NewClient(options),
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

// Get a signed url to start a conversation with an agent with an agent that requires authorization
func (c *Client) GetSignedUrl(ctx context.Context, request *agents.GetSignedUrlConversationsRequest, opts ...option.RequestOption) (*types.ConversationSignedUrlResponseModel, error) {
	response, err := c.WithRawResponse.GetSignedUrl(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a WebRTC session token for real-time communication.
func (c *Client) GetWebrtcToken(ctx context.Context, request *agents.GetWebrtcTokenConversationsRequest, opts ...option.RequestOption) (*types.TokenResponseModel, error) {
	response, err := c.WithRawResponse.GetWebrtcToken(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get all conversations of agents that user owns. With option to restrict to a specific agent.
func (c *Client) List(ctx context.Context, request *agents.ListConversationsRequest, opts ...option.RequestOption) (*types.GetConversationsPageResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Resolve a conversation URL (a Slack message URL or a Zendesk ticket URL) to the deterministic conversation ID for the given agent, then confirm the conversation exists.
func (c *Client) Resolve(ctx context.Context, request *agents.ResolveConversationsRequest, opts ...option.RequestOption) (*types.GetConversationResponseModel, error) {
	response, err := c.WithRawResponse.Resolve(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get the details of a particular conversation
func (c *Client) Get(ctx context.Context, conversationId string, request *agents.GetConversationsRequest, opts ...option.RequestOption) (*types.GetConversationResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, conversationId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a particular conversation
func (c *Client) Delete(ctx context.Context, conversationId string, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Delete(ctx, conversationId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a lightweight summary of a conversation: its title, the generated transcript summary, whether the call was successful, and — only when the conversation is short — the plain chat messages. Tool calls, tool results, and contextual updates are omitted so the response stays small. Use this instead of the full conversation endpoint when you only need the gist (e.g. an agent reading many conversations); use GET /v1/convai/conversations/{conversation_id} when you need the full transcript with tool calls and contextual updates.
func (c *Client) GetSummary(ctx context.Context, conversationId string, request *agents.GetSummaryConversationsRequest, opts ...option.RequestOption) (*types.GetConversationSummaryResponseModel, error) {
	response, err := c.WithRawResponse.GetSummary(ctx, conversationId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get SIP messages associated with a conversation's phone call
func (c *Client) GetSipMessages(ctx context.Context, conversationId string, request *agents.GetSipMessagesConversationsRequest, opts ...option.RequestOption) (*types.GetSipLogMessagesResponse, error) {
	response, err := c.WithRawResponse.GetSipMessages(ctx, conversationId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
