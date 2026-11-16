package transcripts

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	speechtotext "github.com/elevenlabs/elevenlabs-go/speechtotext"
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

// Retrieve a previously generated transcript by its ID.
func (c *Client) Get(ctx context.Context, transcriptionId string, opts ...option.RequestOption) (*speechtotext.GetTranscriptsResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, transcriptionId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a previously generated transcript by its ID.
func (c *Client) Delete(ctx context.Context, transcriptionId string, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Delete(ctx, transcriptionId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
