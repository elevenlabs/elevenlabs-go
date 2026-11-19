package client

import (
	context "context"
	agents "github.com/elevenlabs/elevenlabs-go/agents"
	drafts "github.com/elevenlabs/elevenlabs-go/agents/procedures/drafts"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Drafts *drafts.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Drafts:          drafts.NewClient(options),
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

// List the procedures attached to this agent branch. By default, unpublished drafts take precedence over the latest committed version. Pass agent_version_id to list a published snapshot instead. has_draft is true when a procedure has unpublished draft changes on this branch. Procedure content is not included; use Get Procedure to read a procedure's body.
func (c *Client) List(ctx context.Context, agentId string, branchId string, request *agents.ListProceduresRequest, opts ...option.RequestOption) (*types.ListProceduresResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, agentId, branchId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new procedure for the agent on a branch.
func (c *Client) Create(ctx context.Context, agentId string, branchId string, request agents.CreateProcedureRequestModel, opts ...option.RequestOption) (*types.CreateProcedureResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, agentId, branchId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Legacy. Do not use. Saving an agent draft (`POST /v1/convai/agents/{agent_id}/drafts`) and publishing an agent (`PATCH /v1/convai/agents/{agent_id}`) compile structured procedures into workflow nodes and edges, save the compiled workflow with the draft or version, and return validation errors, so a separate compile call is no longer needed. This endpoint remains available for the time being so existing callers do not break, as a dry-run that compiles the current procedure drafts into a workflow without persisting anything. It will eventually be deprecated.
func (c *Client) Compile(ctx context.Context, agentId string, branchId string, opts ...option.RequestOption) (*types.CompileProceduresResponseModel, error) {
	response, err := c.WithRawResponse.Compile(ctx, agentId, branchId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Retrieve a procedure at a specific version or the current branch HEAD.
func (c *Client) Get(ctx context.Context, agentId string, branchId string, procedureId string, request *agents.GetProceduresRequest, opts ...option.RequestOption) (*types.ProcedureAtVersionResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, agentId, branchId, procedureId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Remove a procedure from the agent's draft working set. Removing a folder cascades to its entire subtree, rejected if any procedure outside the subtree hands off into it.
func (c *Client) Remove(ctx context.Context, agentId string, branchId string, procedureId string, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Remove(ctx, agentId, branchId, procedureId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
