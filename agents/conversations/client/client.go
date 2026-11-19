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

// Create a text conversation with an agent. This does not run the agent, and the agent sends no first message. This endpoint is in alpha and may change.
func (c *Client) Create(ctx context.Context, request *agents.CreateAgentConversationRequest, opts ...option.RequestOption) (*types.CreateAgentConversationResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
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

// Send input to a conversation, and wait until the agent finishes or waits on client tool results. With `background` set to true, the endpoint returns 202 when the server admits the turn. Then stream the turn's events with `Last-Event-ID` set to the `stream_cursor`. This endpoint is in alpha and may change.
func (c *Client) SendMessage(ctx context.Context, conversationId string, request *agents.SendConversationMessageRequest, opts ...option.RequestOption) (*types.ConversationTurn, error) {
	response, err := c.WithRawResponse.SendMessage(ctx, conversationId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Send input to a conversation, and stream the events of the request as server-sent events. The last event is `stream_ended`, which has the outcome of the turn. This endpoint is in alpha and may change.
func (c *Client) SendMessageStream(ctx context.Context, conversationId string, request *agents.SendConversationMessageStreamRequest, opts ...option.RequestOption) (*core.Stream[types.StreamEnvelope], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(options.BaseURL, c.baseURL, "https://api.elevenlabs.io")
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/conversations/%v/messages/stream", conversationId)
	headers := internal.MergeHeaders(
		c.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Set("Accept", "text/event-stream")
	streamer := internal.NewStreamer[types.StreamEnvelope](c.caller)
	return streamer.StreamWithReconnect(ctx, &internal.StreamParams{
		URL:                        endpointURL,
		Method:                     "POST",
		Headers:                    headers,
		MaxAttempts:                options.MaxAttempts,
		DisableRetries:             options.DisableRetries,
		BodyProperties:             options.BodyProperties,
		QueryParameters:            options.QueryParameters,
		Client:                     options.HTTPClient,
		MaxBufSize:                 options.MaxBufSize,
		MaxStreamReconnectAttempts: options.MaxStreamReconnectAttempts,
		DisableStreamReconnection:  options.DisableStreamReconnection,
		Prefix:                     internal.DefaultSSEDataPrefix,
		Delimiter:                  "",
		Terminator:                 "",
		EventDiscriminator:         "",
		Resumable:                  false,
		Format:                     core.StreamFormatSSE,
		RawText:                    false,
		Request:                    request,
		ErrorDecoder:               nil,
	})
}

// Get the status and output of a turn. The status comes from the stored transcript, so a turn that failed after the agent started can show as completed. Only the `stream_ended` event has the real outcome. This endpoint is in alpha and may change.
func (c *Client) GetTurn(ctx context.Context, conversationId string, turnId string, opts ...option.RequestOption) (*types.ConversationTurn, error) {
	response, err := c.WithRawResponse.GetTurn(ctx, conversationId, turnId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Stream the events of a turn as server-sent events. Send `Last-Event-ID` to resume after an event. The server keeps events for about an hour. This endpoint is in alpha and may change.
func (c *Client) StreamTurnEvents(ctx context.Context, conversationId string, turnId string, request *agents.StreamTurnEventsConversationsRequest, opts ...option.RequestOption) (*core.Stream[types.StreamEnvelope], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(options.BaseURL, c.baseURL, "https://api.elevenlabs.io")
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/conversations/%v/turns/%v/events", conversationId, turnId)
	headers := internal.MergeHeaders(
		c.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Set("Accept", "text/event-stream")
	streamer := internal.NewStreamer[types.StreamEnvelope](c.caller)
	return streamer.StreamWithReconnect(ctx, &internal.StreamParams{
		URL:                        endpointURL,
		Method:                     "GET",
		Headers:                    headers,
		MaxAttempts:                options.MaxAttempts,
		DisableRetries:             options.DisableRetries,
		BodyProperties:             options.BodyProperties,
		QueryParameters:            options.QueryParameters,
		Client:                     options.HTTPClient,
		MaxBufSize:                 options.MaxBufSize,
		MaxStreamReconnectAttempts: options.MaxStreamReconnectAttempts,
		DisableStreamReconnection:  options.DisableStreamReconnection,
		Prefix:                     internal.DefaultSSEDataPrefix,
		Delimiter:                  "",
		Terminator:                 "",
		EventDiscriminator:         "",
		Resumable:                  false,
		Format:                     core.StreamFormatSSE,
		RawText:                    false,
		Request:                    request,
		ErrorDecoder:               nil,
	})
}
