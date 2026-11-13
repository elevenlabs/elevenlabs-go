package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	voices "github.com/elevenlabs/elevenlabs-go/voices"
	samplesclient "github.com/elevenlabs/elevenlabs-go/voices/pvc/samples/client"
	verificationclient "github.com/elevenlabs/elevenlabs-go/voices/pvc/verification/client"
)

type Client struct {
	WithRawResponse *RawClient

	Samples      *samplesclient.Client
	Verification *verificationclient.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Samples:         samplesclient.NewClient(options),
		Verification:    verificationclient.NewClient(options),
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

// Creates a new PVC voice with metadata but no samples
func (c *Client) Create(ctx context.Context, request *voices.CreatePvcVoiceRequest, opts ...option.RequestOption) (*types.AddVoiceResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Edit PVC voice metadata
func (c *Client) Update(ctx context.Context, voiceId string, request *voices.BodyEditPvcVoiceV1VoicesPvcVoiceIdPost, opts ...option.RequestOption) (*types.AddVoiceResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, voiceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Start PVC training process for a voice.
func (c *Client) Train(ctx context.Context, voiceId string, request *voices.BodyRunPvcTrainingV1VoicesPvcVoiceIdTrainPost, opts ...option.RequestOption) (*types.StartPvcVoiceTrainingResponseModel, error) {
	response, err := c.WithRawResponse.Train(ctx, voiceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
