package captcha

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	verification "github.com/elevenlabs/elevenlabs-go/voices/pvc/verification"
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

// Get captcha for PVC voice verification.
func (c *Client) Get(ctx context.Context, voiceId string, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Get(ctx, voiceId, opts...)
	return err
}

// Submit captcha verification for PVC voice.
func (c *Client) Verify(ctx context.Context, voiceId string, request *verification.BodyVerifyPvcVoiceCaptchaV1VoicesPvcVoiceIdCaptchaPost, opts ...option.RequestOption) (*types.VerifyPvcVoiceCaptchaResponseModel, error) {
	response, err := c.WithRawResponse.Verify(ctx, voiceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
