package chunk

import (
	context "context"
	documents "github.com/elevenlabs/elevenlabs-go/agents/knowledgebase/documents"
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

// Get details about a specific documentation part used by RAG.
func (c *Client) Get(ctx context.Context, documentationId string, chunkId string, request *documents.GetChunkRequest, opts ...option.RequestOption) (*types.KnowledgeBaseDocumentChunkResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, documentationId, chunkId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
