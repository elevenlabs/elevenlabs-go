package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	audio "github.com/elevenlabs/elevenlabs-go/voices/pvc/samples/speakers/audio"
)

type Client struct {
	WithRawResponse *RawClient

	Audio *audio.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Audio:           audio.NewClient(options),
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

// Retrieve the status of the speaker separation process and the list of detected speakers if complete.
func (c *Client) Get(ctx context.Context, voiceId string, sampleId string, opts ...option.RequestOption) (*types.SpeakerSeparationResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, voiceId, sampleId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Start speaker separation process for a sample
func (c *Client) Separate(ctx context.Context, voiceId string, sampleId string, opts ...option.RequestOption) (*types.StartSpeakerSeparationResponseModel, error) {
	response, err := c.WithRawResponse.Separate(ctx, voiceId, sampleId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
