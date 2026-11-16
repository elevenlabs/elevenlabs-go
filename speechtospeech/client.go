package speechtospeech

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

// Transform audio from one voice to another. Maintain full control over emotion, timing and delivery.
func (c *Client) Convert(ctx context.Context, voiceId string, bodySpeechToSpeechV1SpeechToSpeechVoiceIdPost *types.BodySpeechToSpeechV1SpeechToSpeechVoiceIdPost, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Convert(ctx, voiceId, bodySpeechToSpeechV1SpeechToSpeechVoiceIdPost, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Stream audio from one voice to another. Maintain full control over emotion, timing and delivery.
func (c *Client) Stream(ctx context.Context, voiceId string, bodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost *types.BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Stream(ctx, voiceId, bodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
