package audio

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

// Returns dub as a streamed MP3 or MP4 file. If this dub has been edited using Dubbing Studio you need to use the resource render endpoint as this endpoint only returns the original automatic dub result.
func (c *Client) Get(ctx context.Context, dubbingId string, languageCode string, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Get(ctx, dubbingId, languageCode, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
