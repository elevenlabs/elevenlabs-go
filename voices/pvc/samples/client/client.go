package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	pvc "github.com/elevenlabs/elevenlabs-go/voices/pvc"
	audio "github.com/elevenlabs/elevenlabs-go/voices/pvc/samples/audio"
	speakersclient "github.com/elevenlabs/elevenlabs-go/voices/pvc/samples/speakers/client"
	waveform "github.com/elevenlabs/elevenlabs-go/voices/pvc/samples/waveform"
)

type Client struct {
	WithRawResponse *RawClient

	Audio    *audio.Client
	Waveform *waveform.Client
	Speakers *speakersclient.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Audio:           audio.NewClient(options),
		Waveform:        waveform.NewClient(options),
		Speakers:        speakersclient.NewClient(options),
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

// Add audio samples to a PVC voice
func (c *Client) Create(ctx context.Context, voiceId string, request *pvc.BodyAddSamplesToPvcVoiceV1VoicesPvcVoiceIdSamplesPost, opts ...option.RequestOption) ([]*types.VoiceSample, error) {
	response, err := c.WithRawResponse.Create(ctx, voiceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update a PVC voice sample - apply noise removal, select speaker, change trim times or file name.
func (c *Client) Update(ctx context.Context, voiceId string, sampleId string, request *pvc.BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost, opts ...option.RequestOption) (*types.AddVoiceResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, voiceId, sampleId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a sample from a PVC voice.
func (c *Client) Delete(ctx context.Context, voiceId string, sampleId string, opts ...option.RequestOption) (*types.DeleteVoiceSampleResponseModel, error) {
	response, err := c.WithRawResponse.Delete(ctx, voiceId, sampleId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
