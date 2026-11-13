package history

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

// Returns a list of your generated audio (e.g. text to speech, speech to speech, Studio, dubbing). Music and SFX generations are not included and cannot currently be retrieved via the API.
func (c *Client) List(ctx context.Context, request *types.ListHistoryRequest, opts ...option.RequestOption) (*types.GetSpeechHistoryResponse, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Retrieves a history item.
func (c *Client) Get(ctx context.Context, historyItemId string, opts ...option.RequestOption) (*types.SpeechHistoryItemResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, historyItemId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a history item by its ID
func (c *Client) Delete(ctx context.Context, historyItemId string, opts ...option.RequestOption) (*types.DeleteHistoryItemResponse, error) {
	response, err := c.WithRawResponse.Delete(ctx, historyItemId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Returns the audio of an history item.
func (c *Client) GetAudio(ctx context.Context, historyItemId string, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.GetAudio(ctx, historyItemId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Download one or more history items. If one history item ID is provided, we will return a single audio file. If more than one history item IDs are provided, we will provide the history items packed into a .zip file.
func (c *Client) Download(ctx context.Context, request *types.DownloadHistoryRequest, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Download(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
