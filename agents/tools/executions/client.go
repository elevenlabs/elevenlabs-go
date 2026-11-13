package executions

import (
	context "context"
	tools "github.com/elevenlabs/elevenlabs-go/agents/tools"
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

// Get paginated list of tool executions for a specific tool.
func (c *Client) Get(ctx context.Context, toolId string, request *tools.GetExecutionsRequest, opts ...option.RequestOption) (*types.GetToolExecutionsPageResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, toolId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
