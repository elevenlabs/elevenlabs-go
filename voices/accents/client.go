package accents

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	voices "github.com/elevenlabs/elevenlabs-go/voices"
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

// Gets the list of available accents in the shared voice library.
func (c *Client) Get(ctx context.Context, request *voices.GetAccentsRequest, opts ...option.RequestOption) (*types.GetVoiceAccentsResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
