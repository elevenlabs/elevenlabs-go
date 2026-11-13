package audio

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	samples "github.com/elevenlabs/elevenlabs-go/voices/pvc/samples"
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

// Retrieve the first 30 seconds of voice sample audio with or without noise removal.
func (c *Client) Get(ctx context.Context, voiceId string, sampleId string, request *samples.GetAudioRequest, opts ...option.RequestOption) (*types.VoiceSamplePreviewResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, voiceId, sampleId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
