package audioeffects

import (
	context "context"
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

// Lists the filter presets, reverb environments, and background noises that `audio_effects` can reference in Text to Speech, Text to Dialogue, and agent configurations.
func (c *Client) GetCatalog(ctx context.Context, opts ...option.RequestOption) (*types.AudioEffectsCatalogResponseModel, error) {
	response, err := c.WithRawResponse.GetCatalog(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
