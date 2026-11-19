package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	pvc "github.com/elevenlabs/elevenlabs-go/voices/pvc"
	captcha "github.com/elevenlabs/elevenlabs-go/voices/pvc/verification/captcha"
)

type Client struct {
	WithRawResponse *RawClient

	Captcha *captcha.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Captcha:         captcha.NewClient(options),
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

// Request manual verification for a PVC voice.
func (c *Client) Request(ctx context.Context, voiceId string, request *pvc.BodyRequestManualVerificationV1VoicesPvcVoiceIdVerificationPost, opts ...option.RequestOption) (*types.RequestPvcManualVerificationResponseModel, error) {
	response, err := c.WithRawResponse.Request(ctx, voiceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
