package texttosoundeffects

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	"io"
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

// Turn text into sound effects for your videos, voice-overs or video games using the most advanced sound effects models in the world.
func (c *Client) Convert(ctx context.Context, request *types.CreateSoundEffectRequest, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Convert(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
