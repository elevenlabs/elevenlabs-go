package knowledgesources

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	reception "github.com/elevenlabs/elevenlabs-go/reception"
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

// List all knowledge sources for the current workspace
func (c *Client) List(ctx context.Context, request *reception.ListKnowledgeSourcesRequest, opts ...option.RequestOption) ([]*types.ReceptionKnowledgeSource, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Upload a file as a knowledge source
func (c *Client) Upload(ctx context.Context, bodySmbUploadKnowledgeSource *reception.BodySmbUploadKnowledgeSource, opts ...option.RequestOption) (*types.ReceptionKnowledgeSource, error) {
	response, err := c.WithRawResponse.Upload(ctx, bodySmbUploadKnowledgeSource, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a specific knowledge source by ID
func (c *Client) Get(ctx context.Context, sourceId string, request *reception.GetKnowledgeSourcesRequest, opts ...option.RequestOption) (*types.ReceptionKnowledgeSource, error) {
	response, err := c.WithRawResponse.Get(ctx, sourceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a knowledge source
func (c *Client) Delete(ctx context.Context, sourceId string, request *reception.DeleteKnowledgeSourcesRequest, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.Delete(ctx, sourceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Add a URL as a knowledge source
func (c *Client) AddUrl(ctx context.Context, request *reception.AddUrlRequest, opts ...option.RequestOption) (*types.ReceptionKnowledgeSource, error) {
	response, err := c.WithRawResponse.AddUrl(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
