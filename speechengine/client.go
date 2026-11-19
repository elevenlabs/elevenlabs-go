package speechengine

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

// Returns a paginated list of Speech Engine resources.
func (c *Client) List(ctx context.Context, request *types.ListSpeechEngineRequest, opts ...option.RequestOption) (*types.ListSpeechEnginesResponse, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new Speech Engine resource
func (c *Client) Create(ctx context.Context, request *types.CreateSpeechEngineRequest, opts ...option.RequestOption) (*types.SpeechEngineResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Retrieve a Speech Engine resource
func (c *Client) Get(ctx context.Context, speechEngineId string, opts ...option.RequestOption) (*types.SpeechEngineResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, speechEngineId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update a Speech Engine resource (partial update)
func (c *Client) Update(ctx context.Context, speechEngineId string, request *types.UpdateSpeechEngineRequest, opts ...option.RequestOption) (*types.SpeechEngineResponse, error) {
	response, err := c.WithRawResponse.Update(ctx, speechEngineId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a Speech Engine resource
func (c *Client) Delete(ctx context.Context, speechEngineId string, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Delete(ctx, speechEngineId, opts...)
	return err
}

// Create a new Speech Engine resource by duplicating an existing one
func (c *Client) Duplicate(ctx context.Context, speechEngineId string, request types.DuplicateSpeechEngineRequest, opts ...option.RequestOption) (*types.SpeechEngineResponse, error) {
	response, err := c.WithRawResponse.Duplicate(ctx, speechEngineId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
