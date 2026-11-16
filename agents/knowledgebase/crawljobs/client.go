package crawljobs

import (
	context "context"
	knowledgebase "github.com/elevenlabs/elevenlabs-go/agents/knowledgebase"
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

// Get a list of ongoing and recent crawl jobs for the user.
func (c *Client) List(ctx context.Context, request *knowledgebase.ListCrawlJobsRequest, opts ...option.RequestOption) (*types.ListCrawlJobsResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a crawl job to crawl the given URL with specified depth and page limits.
func (c *Client) Create(ctx context.Context, request *knowledgebase.BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost, opts ...option.RequestOption) (*types.CreateCrawlJobResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get details about a specific crawl job including status and progress.
func (c *Client) Get(ctx context.Context, crawlJobId string, opts ...option.RequestOption) (*types.GetCrawlJobResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, crawlJobId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Cancel a crawl job and delete all associated documents and folders.
func (c *Client) Cancel(ctx context.Context, crawlJobId string, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Cancel(ctx, crawlJobId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
