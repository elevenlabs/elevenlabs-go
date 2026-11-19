package summaries

import (
	context "context"
	documents "github.com/elevenlabs/elevenlabs-go/agents/knowledgebase/documents"
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

// Gets multiple knowledge base document summaries by their IDs.
func (c *Client) Get(ctx context.Context, request *documents.GetSummariesRequest, opts ...option.RequestOption) (map[string]*documents.GetSummariesResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
