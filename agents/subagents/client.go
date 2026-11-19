package subagents

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

// Get a list of agents that use this agent as a subagent.
func (c *Client) GetDependents(ctx context.Context, agentId string, request *agents.GetDependentsSubagentsRequest, opts ...option.RequestOption) (*types.GetAgentSubagentDependentsResponseModel, error) {
	response, err := c.WithRawResponse.GetDependents(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
