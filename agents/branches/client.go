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

// Returns the result of merging the source branch into the target branch without performing the merge. When `conflicts` is not empty, the merge itself is rejected unless `conflict_resolutions` settles every conflict or `force` is set; until then the returned configuration shows one side of each conflict only as a placeholder.
func (c *Client) PreviewMerge(ctx context.Context, agentId string, sourceBranchId string, request *agents.PreviewMergeBranchesRequest, opts ...option.RequestOption) (*types.MergePreviewResponseModel, error) {
	response, err := c.WithRawResponse.PreviewMerge(ctx, agentId, sourceBranchId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Merge a branch into a target branch. If both branches changed the same field, the merge returns 409 unless `conflict_resolutions` settles every conflict the merge preview reports, or `force` is set.
func (c *Client) Merge(ctx context.Context, agentId string, sourceBranchId string, request *agents.BodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Merge(ctx, agentId, sourceBranchId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Returns the result of pulling changes from a source branch (defaults to main) into the given branch without performing it. When `conflicts` is not empty, the rebase itself is rejected unless `conflict_resolutions` settles every conflict; until then the returned configuration keeps the given branch's side of each conflict.
func (c *Client) PreviewRebase(ctx context.Context, agentId string, branchId string, request *agents.PreviewRebaseBranchesRequest, opts ...option.RequestOption) (*types.MergePreviewResponseModel, error) {
	response, err := c.WithRawResponse.PreviewRebase(ctx, agentId, branchId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Pull changes from a source branch (defaults to main) into this branch as a new version. If both branches changed the same field, the rebase returns 409 unless `conflict_resolutions` settles every conflict the rebase preview reports.
func (c *Client) Rebase(ctx context.Context, agentId string, branchId string, request *agents.BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Rebase(ctx, agentId, branchId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
