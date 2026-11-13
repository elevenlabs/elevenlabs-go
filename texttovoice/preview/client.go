package preview

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
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

// Stream a voice preview that was created via the /v1/text-to-voice/design endpoint.
func (c *Client) Stream(ctx context.Context, generatedVoiceId string, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Stream(ctx, generatedVoiceId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
