package audionative

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

// Creates Audio Native enabled project, optionally starts conversion and returns project ID and embeddable HTML snippet.
func (c *Client) Create(ctx context.Context, request *types.BodyCreatesAudioNativeEnabledProjectV1AudioNativePost, opts ...option.RequestOption) (*types.AudioNativeCreateProjectResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get player settings for the specific project.
func (c *Client) GetSettings(ctx context.Context, projectId string, opts ...option.RequestOption) (*types.GetAudioNativeProjectSettingsResponseModel, error) {
	response, err := c.WithRawResponse.GetSettings(ctx, projectId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Updates content for the specific AudioNative Project.
func (c *Client) Update(ctx context.Context, projectId string, request *types.BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost, opts ...option.RequestOption) (*types.AudioNativeEditContentResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, projectId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Finds an AudioNative project matching the provided URL, extracts content from the URL, updates the project content, and queues it for conversion and auto-publishing.
func (c *Client) UpdateContentFromUrl(ctx context.Context, request *types.BodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPost, opts ...option.RequestOption) (*types.AudioNativeEditContentResponseModel, error) {
	response, err := c.WithRawResponse.UpdateContentFromUrl(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
