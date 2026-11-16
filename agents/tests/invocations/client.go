package invocations

import (
	context "context"
	tests "github.com/elevenlabs/elevenlabs-go/agents/tests"
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

// Lists all test invocations with pagination support and optional search filtering.
func (c *Client) List(ctx context.Context, request *tests.ListInvocationsRequest, opts ...option.RequestOption) (*types.GetTestInvocationsPageResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Gets a test invocation by ID.
func (c *Client) Get(ctx context.Context, testInvocationId string, opts ...option.RequestOption) (*types.GetTestSuiteInvocationResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, testInvocationId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Updates the user-provided name of a test invocation.
func (c *Client) Update(ctx context.Context, testInvocationId string, request *tests.UpdateTestSuiteInvocationRequestModel, opts ...option.RequestOption) (*types.GetTestSuiteInvocationResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, testInvocationId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Cancels all pending runs in a test invocation.
func (c *Client) Cancel(ctx context.Context, testInvocationId string, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Cancel(ctx, testInvocationId, opts...)
	return err
}

// Resubmits specific test runs from a test invocation.
func (c *Client) Resubmit(ctx context.Context, testInvocationId string, request *tests.ResubmitTestsRequestModel, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Resubmit(ctx, testInvocationId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
