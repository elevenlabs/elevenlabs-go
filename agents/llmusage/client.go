package llmusage

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

// Calculates expected number of LLM tokens needed for the specified agent.
func (c *Client) CalculateForAgent(ctx context.Context, agentId string, request *agents.LlmUsageCalculatorRequestModel, opts ...option.RequestOption) (*types.LlmUsageCalculatorResponseModel, error) {
	response, err := c.WithRawResponse.CalculateForAgent(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Returns a list of LLM models and the expected cost for using them based on the provided values.
func (c *Client) Calculate(ctx context.Context, request *agents.LlmUsageCalculatorPublicRequestModel, opts ...option.RequestOption) (*types.LlmUsageCalculatorResponseModel, error) {
	response, err := c.WithRawResponse.Calculate(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
