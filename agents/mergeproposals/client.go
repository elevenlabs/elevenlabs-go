package mergeproposals

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

// List the proposals for an agent, newest first.
func (c *Client) List(ctx context.Context, agentId string, request *agents.ListMergeProposalsRequest, opts ...option.RequestOption) (*types.PaginatedResultAgentMergeProposalResponse, error) {
	response, err := c.WithRawResponse.List(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Record a request to merge a source branch into a target branch. Anyone with edit access can open one; merging it later is gated on write access to the target branch, so this is how a change reaches a protected branch the author cannot merge into themselves.
func (c *Client) Create(ctx context.Context, agentId string, request *agents.BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost, opts ...option.RequestOption) (*types.CreateAgentMergeProposalResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a single merge_proposal.
func (c *Client) Get(ctx context.Context, agentId string, mergeProposalId string, opts ...option.RequestOption) (*types.AgentMergeProposalResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, agentId, mergeProposalId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Edit an open merge_proposal's title or description, or close it. The author closing it is recorded as withdrawn; anyone else as rejected.
func (c *Client) Update(ctx context.Context, agentId string, mergeProposalId string, request *agents.BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch, opts ...option.RequestOption) (*types.AgentMergeProposalResponse, error) {
	response, err := c.WithRawResponse.Update(ctx, agentId, mergeProposalId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Approve a merge_proposal or request changes on it. A user's latest review replaces their previous one. Non-admins need an approval from another user before the merge is allowed.
func (c *Client) SubmitReview(ctx context.Context, agentId string, mergeProposalId string, request *agents.BodyReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost, opts ...option.RequestOption) (*types.AgentMergeProposalResponse, error) {
	response, err := c.WithRawResponse.SubmitReview(ctx, agentId, mergeProposalId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Leave a comment on a merge_proposal without recording a review verdict. Unlike reviews, comments accumulate and can still be added once the merge_proposal is merged or closed.
func (c *Client) AddComment(ctx context.Context, agentId string, mergeProposalId string, request *agents.BodyCommentOnAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdCommentsPost, opts ...option.RequestOption) (*types.AgentMergeProposalResponse, error) {
	response, err := c.WithRawResponse.AddComment(ctx, agentId, mergeProposalId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Execute the merge. The caller must have write access to the target branch (admins only, for a protected branch), so this is where a reviewer approves and merges a request opened by someone who could not.
func (c *Client) Merge(ctx context.Context, agentId string, mergeProposalId string, request *agents.BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost, opts ...option.RequestOption) (*types.AgentMergeProposalResponse, error) {
	response, err := c.WithRawResponse.Merge(ctx, agentId, mergeProposalId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
