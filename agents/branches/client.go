package branches

import (
	context "context"
	agents "github.com/elevenlabs/elevenlabs-go/agents"
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

// Returns a list of branches an agent has
func (c *Client) List(ctx context.Context, agentId string, request *agents.ListBranchesRequest, opts ...option.RequestOption) (*types.ListResponseAgentBranchSummary, error) {
	response, err := c.WithRawResponse.List(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new branch from a given version of any branch
func (c *Client) Create(ctx context.Context, agentId string, request *agents.BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost, opts ...option.RequestOption) (*types.CreateAgentBranchResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get information about a single agent branch
func (c *Client) Get(ctx context.Context, agentId string, branchId string, opts ...option.RequestOption) (*types.AgentBranchResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, agentId, branchId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update agent branch properties such as name, description, archiving status, protection level, and auto-rebase onto main
func (c *Client) Update(ctx context.Context, agentId string, branchId string, request *agents.BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch, opts ...option.RequestOption) (*types.AgentBranchResponse, error) {
	response, err := c.WithRawResponse.Update(ctx, agentId, branchId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Returns the result of merging the source branch into the target branch without performing the merge. Useful for showing an accurate diff before confirming.
func (c *Client) PreviewMerge(ctx context.Context, agentId string, sourceBranchId string, request *agents.PreviewMergeBranchesRequest, opts ...option.RequestOption) (*types.MergePreviewResponseModel, error) {
	response, err := c.WithRawResponse.PreviewMerge(ctx, agentId, sourceBranchId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Merge a branch into a target branch
func (c *Client) Merge(ctx context.Context, agentId string, sourceBranchId string, request *agents.BodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Merge(ctx, agentId, sourceBranchId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Returns the result of pulling changes from a source branch (defaults to main) into the given branch without performing it, preserving the given branch's own changes. Useful for showing an accurate diff before confirming.
func (c *Client) PreviewRebase(ctx context.Context, agentId string, branchId string, request *agents.PreviewRebaseBranchesRequest, opts ...option.RequestOption) (*types.MergePreviewResponseModel, error) {
	response, err := c.WithRawResponse.PreviewRebase(ctx, agentId, branchId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Pull changes from a source branch (defaults to main) into this branch as a new version, preserving this branch's own changes on conflicts.
func (c *Client) Rebase(ctx context.Context, agentId string, branchId string, request *agents.BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Rebase(ctx, agentId, branchId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
