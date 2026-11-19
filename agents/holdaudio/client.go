package holdaudio

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

// Sets the custom hold audio played on loop to callers waiting in the agent's concurrency wait queue. Replaces any previously uploaded clip.
func (c *Client) Create(ctx context.Context, agentId string, request *agents.BodyPostAgentHoldAudioV1ConvaiAgentsAgentIdHoldAudioPost, opts ...option.RequestOption) (*types.PostAgentHoldAudioResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Removes the agent's custom hold audio; queued callers hear the default hold tone again.
func (c *Client) Delete(ctx context.Context, agentId string, opts ...option.RequestOption) (*types.DeleteAgentHoldAudioResponseModel, error) {
	response, err := c.WithRawResponse.Delete(ctx, agentId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
