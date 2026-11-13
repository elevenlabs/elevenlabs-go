package settings

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

// Gets the default settings for voices. "similarity_boost" corresponds to"Clarity + Similarity Enhancement" in the web app and "stability" corresponds to "Stability" slider in the web app.
func (c *Client) GetDefault(ctx context.Context, opts ...option.RequestOption) (*types.VoiceSettingsResponseModel, error) {
	response, err := c.WithRawResponse.GetDefault(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Returns the settings for a specific voice. "similarity_boost" corresponds to"Clarity + Similarity Enhancement" in the web app and "stability" corresponds to "Stability" slider in the web app.
func (c *Client) Get(ctx context.Context, voiceId string, opts ...option.RequestOption) (*types.VoiceSettingsResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, voiceId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Edit your settings for a specific voice. "similarity_boost" corresponds to "Clarity + Similarity Enhancement" in the web app and "stability" corresponds to "Stability" slider in the web app.
func (c *Client) Update(ctx context.Context, voiceId string, request *types.VoiceSettingsResponseModel, opts ...option.RequestOption) (*types.EditVoiceSettingsResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, voiceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
