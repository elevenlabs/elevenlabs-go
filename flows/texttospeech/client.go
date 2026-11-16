package texttospeech

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	flows "github.com/elevenlabs/elevenlabs-go/flows"
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

// List the speech generations created through this API, newest first.
func (c *Client) List(ctx context.Context, request *flows.ListTextToSpeechRequest, opts ...option.RequestOption) (*types.MediaGenerationListResponse, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Start a speech generation with the selected model. Charged per character via text-to-speech billing. Use this over `/v1/text-to-speech` for the asynchronous generation lifecycle or for models not offered there; for direct, synchronous speech synthesis, prefer `/v1/text-to-speech`.
func (c *Client) Create(ctx context.Context, textToSpeechGenerationRequest *types.TextToSpeechGenerationRequest, opts ...option.RequestOption) (*types.MediaGenerationCreateResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, textToSpeechGenerationRequest, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Retrieve the status of a speech generation, and retrieve its output URL once completed.
func (c *Client) Get(ctx context.Context, generationId string, opts ...option.RequestOption) (*types.MediaGenerationResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, generationId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
