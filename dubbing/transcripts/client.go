package transcripts

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	dubbing "github.com/elevenlabs/elevenlabs-go/dubbing"
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

// Fetch the transcript for one of the languages in a dub.
func (c *Client) Get(ctx context.Context, dubbingId string, languageCode string, formatType *dubbing.GetTranscriptsRequestFormatType, opts ...option.RequestOption) (*types.DubbingTranscriptsResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, dubbingId, languageCode, formatType, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
