package procedures

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	receptionists "github.com/elevenlabs/elevenlabs-go/reception/receptionists"
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

// List the procedures (SOPs) on the named receptionist, with the tools each one is missing.
func (c *Client) List(ctx context.Context, agentId string, request *receptionists.ListProceduresRequest, opts ...option.RequestOption) ([]*types.ReceptionProcedureResponse, error) {
	response, err := c.WithRawResponse.List(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Add a procedure (SOP) to the named receptionist.
func (c *Client) Create(ctx context.Context, agentId string, request *receptionists.ReceptionCreateProcedureRequest, opts ...option.RequestOption) (*types.ReceptionProcedure, error) {
	response, err := c.WithRawResponse.Create(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Replace an existing procedure on the named receptionist. Send every field: an omitted `body` resets to empty.
func (c *Client) Update(ctx context.Context, agentId string, procedureId string, request *receptionists.ReceptionUpdateProcedureRequest, opts ...option.RequestOption) (*types.ReceptionProcedure, error) {
	response, err := c.WithRawResponse.Update(ctx, agentId, procedureId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a procedure from the named receptionist.
func (c *Client) Delete(ctx context.Context, agentId string, procedureId string, request *receptionists.DeleteProceduresRequest, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.Delete(ctx, agentId, procedureId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
