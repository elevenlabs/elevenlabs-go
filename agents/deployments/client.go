package deployments

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

// List the traffic split history of an agent, newest first
func (c *Client) List(ctx context.Context, agentId string, request *agents.ListDeploymentsRequest, opts ...option.RequestOption) (*types.ListResponseAgentDeploymentHistoryItem, error) {
	response, err := c.WithRawResponse.List(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new deployment for an agent
func (c *Client) Create(ctx context.Context, agentId string, request *agents.BodyCreateOrUpdateDeploymentsV1ConvaiAgentsAgentIdDeploymentsPost, opts ...option.RequestOption) (*types.AgentDeploymentResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
