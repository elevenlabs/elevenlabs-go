package audio

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

// Retrieve the separated audio for a specific speaker.
func (c *Client) Get(ctx context.Context, voiceId string, sampleId string, speakerId string, opts ...option.RequestOption) (*types.SpeakerAudioResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, voiceId, sampleId, speakerId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
