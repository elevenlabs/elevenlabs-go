package batchcalls

import (
	context "context"
	agents "github.com/elevenlabs/elevenlabs-go/agents"
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

// Submit a batch call request to schedule calls for multiple recipients.
func (c *Client) Create(ctx context.Context, request *agents.BodySubmitABatchCallRequestV1ConvaiBatchCallingSubmitPost, opts ...option.RequestOption) (*types.BatchCallResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get all batch calls for the current workspace.
func (c *Client) List(ctx context.Context, request *agents.ListBatchCallsRequest, opts ...option.RequestOption) (*types.WorkspaceBatchCallsResponse, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get detailed information about a batch call including all recipients.
func (c *Client) Get(ctx context.Context, batchId string, opts ...option.RequestOption) (*types.BatchCallDetailedResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, batchId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Permanently delete a batch call and all recipient records. Conversations remain in history.
func (c *Client) Delete(ctx context.Context, batchId string, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Delete(ctx, batchId, opts...)
	return err
}

// Cancel a running batch call and set all recipients to cancelled status.
func (c *Client) Cancel(ctx context.Context, batchId string, opts ...option.RequestOption) (*types.BatchCallResponse, error) {
	response, err := c.WithRawResponse.Cancel(ctx, batchId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Retry a batch call, calling failed and no-response recipients again.
func (c *Client) Retry(ctx context.Context, batchId string, opts ...option.RequestOption) (*types.BatchCallResponse, error) {
	response, err := c.WithRawResponse.Retry(ctx, batchId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Download all recipients and conversation results for a batch call as CSV, including calls that are still in progress.
func (c *Client) Export(ctx context.Context, batchId string, request *agents.ExportBatchCallsRequest, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Export(ctx, batchId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
