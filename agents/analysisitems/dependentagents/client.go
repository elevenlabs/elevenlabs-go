package dependentagents

import (
	context "context"
	analysisitems "github.com/elevenlabs/elevenlabs-go/agents/analysisitems"
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

// Get a list of agents (and branches) that reference this analysis item.
func (c *Client) Get(ctx context.Context, analysisItemId string, request *analysisitems.GetDependentAgentsRequest, opts ...option.RequestOption) (*types.GetAnalysisItemDependentAgentsResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, analysisItemId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
