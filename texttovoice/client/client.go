package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	preview "github.com/elevenlabs/elevenlabs-go/texttovoice/preview"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Preview *preview.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Preview:         preview.NewClient(options),
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

// Create a voice from previously generated voice preview. This endpoint should be called after you fetched a generated_voice_id using POST /v1/text-to-voice/design or POST /v1/text-to-voice/:voice_id/remix.
func (c *Client) Create(ctx context.Context, request *types.BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost, opts ...option.RequestOption) (*types.Voice, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Design a voice via a prompt. This method returns a list of voice previews. Each preview has a generated_voice_id and a sample of the voice as base64 encoded mp3 audio. To create a voice use the generated_voice_id of the preferred preview with the /v1/text-to-voice endpoint.
func (c *Client) Design(ctx context.Context, request *types.VoiceDesignRequestModel, opts ...option.RequestOption) (*types.VoiceDesignPreviewResponse, error) {
	response, err := c.WithRawResponse.Design(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Remix an existing voice via a prompt. This method returns a list of voice previews. Each preview has a generated_voice_id and a sample of the voice as base64 encoded mp3 audio. To create a voice use the generated_voice_id of the preferred preview with the /v1/text-to-voice endpoint.
func (c *Client) Remix(ctx context.Context, voiceId string, request *types.VoiceRemixRequestModel, opts ...option.RequestOption) (*types.VoiceDesignPreviewResponse, error) {
	response, err := c.WithRawResponse.Remix(ctx, voiceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
