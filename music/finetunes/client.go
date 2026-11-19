package finetunes

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	music "github.com/elevenlabs/elevenlabs-go/music"
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

// List music finetunes accessible to you (your own, workspace-shared, and ElevenLabs-curated), with optional filtering, sorting, and cursor pagination.
func (c *Client) List(ctx context.Context, request *music.ListFinetunesRequest, opts ...option.RequestOption) (*types.MusicFinetunePageResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new music finetune
func (c *Client) Create(ctx context.Context, request *music.BodyCreateMusicFinetuneV1MusicFinetunesPost, opts ...option.RequestOption) (*types.MusicFinetuneResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a music finetune.
func (c *Client) Get(ctx context.Context, finetuneId string, opts ...option.RequestOption) (*types.MusicFinetuneResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, finetuneId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update a music finetune.
func (c *Client) Update(ctx context.Context, finetuneId string, request *music.UpdateMusicFinetuneRequestModel, opts ...option.RequestOption) (*types.MusicFinetuneResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, finetuneId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a music finetune
func (c *Client) Delete(ctx context.Context, finetuneId string, opts ...option.RequestOption) (*types.MusicFinetuneResponseModel, error) {
	response, err := c.WithRawResponse.Delete(ctx, finetuneId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
