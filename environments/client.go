package environments

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

// List the workspace environments, including 'production'.
func (c *Client) List(ctx context.Context, opts ...option.RequestOption) (*types.EnvironmentsListResponse, error) {
	response, err := c.WithRawResponse.List(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a workspace environment. Variables without a value for it fall back to their production value.
func (c *Client) Create(ctx context.Context, request *types.CreateEnvironmentRequest, opts ...option.RequestOption) (*types.EnvironmentResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a workspace environment and remove its values from all environment variables. 'production' cannot be deleted.
func (c *Client) Delete(ctx context.Context, environment string, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Delete(ctx, environment, opts...)
	return err
}
