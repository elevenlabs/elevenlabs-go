package summaries

import (
	context "context"
	agents "github.com/elevenlabs/elevenlabs-go/agents"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
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

// Returns summaries for the specified agents.
func (c *Client) Get(ctx context.Context, request *agents.GetSummariesRequest, opts ...option.RequestOption) (map[string]*agents.GetSummariesResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
