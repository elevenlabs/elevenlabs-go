package compositionplan

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	music "github.com/elevenlabs/elevenlabs-go/music"
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

// Create a composition plan for music generation. Usage of this endpoint does not cost any credits but is subject to rate limiting depending on your tier.
func (c *Client) Create(ctx context.Context, request *music.BodyGenerateCompositionPlanV1MusicPlanPost, opts ...option.RequestOption) (*music.CreateCompositionPlanResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
