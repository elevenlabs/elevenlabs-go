package client

import (
	context "context"
	agents "github.com/elevenlabs/elevenlabs-go/agents"
	crawljobs "github.com/elevenlabs/elevenlabs-go/agents/knowledgebase/crawljobs"
	document "github.com/elevenlabs/elevenlabs-go/agents/knowledgebase/document"
	documentsclient "github.com/elevenlabs/elevenlabs-go/agents/knowledgebase/documents/client"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Documents *documentsclient.Client
	CrawlJobs *crawljobs.Client
	Document  *document.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Documents:       documentsclient.NewClient(options),
		CrawlJobs:       crawljobs.NewClient(options),
		Document:        document.NewClient(options),
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

// Returns the number of pages in the agent's knowledge base.
func (c *Client) Size(ctx context.Context, agentId string, opts ...option.RequestOption) (*types.GetAgentKnowledgebaseSizeResponseModel, error) {
	response, err := c.WithRawResponse.Size(ctx, agentId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a list of available knowledge base documents
func (c *Client) List(ctx context.Context, request *agents.ListKnowledgeBaseRequest, opts ...option.RequestOption) (*types.GetKnowledgeBaseListResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Retrieves and/or creates RAG indexes for multiple knowledge base documents in a single request. Maximum 100 items per request.
func (c *Client) GetOrCreateRagIndexes(ctx context.Context, request *agents.BodyComputeRagIndexesInBatchV1ConvaiKnowledgeBaseRagIndexPost, opts ...option.RequestOption) (map[string]*agents.GetOrCreateRagIndexesKnowledgeBaseResponse, error) {
	response, err := c.WithRawResponse.GetOrCreateRagIndexes(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Fuzzy text search over knowledge base document content
func (c *Client) Search(ctx context.Context, request *agents.SearchKnowledgeBaseRequest, opts ...option.RequestOption) (*types.KnowledgeBaseContentSearchResponseModel, error) {
	response, err := c.WithRawResponse.Search(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
