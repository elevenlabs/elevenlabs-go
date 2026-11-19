package drafts

import (
	context "context"
	procedures "github.com/elevenlabs/elevenlabs-go/agents/procedures"
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

// Get user's draft for a procedure
func (c *Client) Get(ctx context.Context, agentId string, branchId string, procedureId string, opts ...option.RequestOption) (*types.ProcedureDraftResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, agentId, branchId, procedureId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create or update user's draft for a procedure
func (c *Client) Update(ctx context.Context, agentId string, branchId string, procedureId string, request *procedures.UpdateProcedureDraftRequestModel, opts ...option.RequestOption) (*types.ProcedureDraftResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, agentId, branchId, procedureId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete user's draft for a procedure, resetting to the committed version
func (c *Client) Delete(ctx context.Context, agentId string, branchId string, procedureId string, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Delete(ctx, agentId, branchId, procedureId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
