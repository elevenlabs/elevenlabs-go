package audioisolation

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
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

// Removes background noise from audio.
func (c *Client) Convert(ctx context.Context, audioIsolationRequest *types.AudioIsolationRequest, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Convert(ctx, audioIsolationRequest, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Returns a list of all your audio isolation generations.
func (c *Client) List(ctx context.Context, request *types.ListAudioIsolationRequest, opts ...option.RequestOption) (*types.GetAudioIsolationHistoryResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Deletes a specific audio isolation history item and the associated media files.
func (c *Client) Delete(ctx context.Context, historyItemId string, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Delete(ctx, historyItemId, opts...)
	return err
}

// Removes background noise from audio.
func (c *Client) Stream(ctx context.Context, audioIsolationStreamRequest *types.AudioIsolationStreamRequest, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Stream(ctx, audioIsolationStreamRequest, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
