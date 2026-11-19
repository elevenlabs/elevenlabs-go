package contentprovenance

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

// Check an audio file for an embedded ElevenLabs watermark. This endpoint is in beta and may change.
func (c *Client) Watermark(ctx context.Context, request *types.BodyCheckWatermarkV1ContentProvenanceWatermarkPost, opts ...option.RequestOption) (*types.WatermarkCheck, error) {
	response, err := c.WithRawResponse.Watermark(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
