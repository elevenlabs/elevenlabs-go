package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	transcripts "github.com/elevenlabs/elevenlabs-go/speechtotext/transcripts"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Transcripts *transcripts.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Transcripts:     transcripts.NewClient(options),
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

// Transcribe an audio or video file. If webhook is set to true, the request will be processed asynchronously and results sent to configured webhooks. When use_multi_channel is true and the provided audio has multiple channels, a 'transcripts' object with separate transcripts for each channel is returned; set multichannel_output_style='combined' to instead receive a single transcript with all channels merged and sorted by time. Otherwise, returns a single transcript. The optional webhook_metadata parameter allows you to attach custom data that will be included in webhook responses for request correlation and tracking.
func (c *Client) Convert(ctx context.Context, bodySpeechToTextV1SpeechToTextPost *types.BodySpeechToTextV1SpeechToTextPost, opts ...option.RequestOption) (*types.ConvertSpeechToTextResponse, error) {
	response, err := c.WithRawResponse.Convert(ctx, bodySpeechToTextV1SpeechToTextPost, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
