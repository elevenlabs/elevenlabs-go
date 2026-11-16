package triagetickets

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

// List an agent's issues, ordered by most recently created first unless sorted by priority. These are issues about the agent's own performance on a conversation (to work on with Architect), not tickets an agent opens for end users.
func (c *Client) List(ctx context.Context, agentId string, request *agents.ListTriageTicketsRequest, opts ...option.RequestOption) (*types.GetAgentConversationTicketsPageResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Manually raise a follow-up issue against an agent, not tied to any conversation (for example a task like 'add the KB about X'). Without a title, one is derived from the comment. Requires viewer access to the agent.
func (c *Client) CreateManual(ctx context.Context, agentId string, request *agents.CreateManualTicketRequestModel, opts ...option.RequestOption) (*types.AgentConversationTicketResponseModel, error) {
	response, err := c.WithRawResponse.CreateManual(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// List agent issues across every agent in the workspace, ordered by most recently created first. Use this to build a workspace-wide view (for example, issues assigned to the caller). Supports the same filters, sorting and search as the per-agent endpoint. Issues for agents the caller cannot access are omitted.
func (c *Client) ListForWorkspace(ctx context.Context, request *agents.ListForWorkspaceTriageTicketsRequest, opts ...option.RequestOption) (*types.GetAgentConversationTicketsPageResponseModel, error) {
	response, err := c.WithRawResponse.ListForWorkspace(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Raise an issue about an agent's performance on a conversation, to work on with Architect. Provide an overall comment and/or turn-level comments describing what went wrong.
func (c *Client) Create(ctx context.Context, request *agents.CreateAgentConversationTicketRequestModel, opts ...option.RequestOption) (*types.AgentConversationTicketResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// All non-service-account workspace members, each flagged with whether they currently have at least viewer access to the agent. Members without access are included (not filtered out) so the UI can offer them as an assignee and prompt to grant access first.
func (c *Client) ListAssignableUsers(ctx context.Context, agentId string, opts ...option.RequestOption) ([]*types.AssignableUserResponseModel, error) {
	response, err := c.WithRawResponse.ListAssignableUsers(ctx, agentId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get an agent issue by ID.
func (c *Client) Get(ctx context.Context, agentqaTicketId string, opts ...option.RequestOption) (*types.AgentConversationTicketResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, agentqaTicketId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update an agent issue's title, comment, status, priority, and/or assignee. Requires viewer access to the issue's agent.
func (c *Client) Update(ctx context.Context, agentqaTicketId string, request *agents.PatchAgentConversationTicketRequestModel, opts ...option.RequestOption) (*types.AgentConversationTicketResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, agentqaTicketId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete an agent issue. Restricted to the issue creator or a workspace admin.
func (c *Client) Delete(ctx context.Context, agentqaTicketId string, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Delete(ctx, agentqaTicketId, opts...)
	return err
}

// Append a comment discussing how to resolve the agent issue. Requires viewer access to the issue's agent.
func (c *Client) AddComment(ctx context.Context, agentqaTicketId string, request *agents.AddTicketCommentRequestModel, opts ...option.RequestOption) (*types.AgentConversationTicketResponseModel, error) {
	response, err := c.WithRawResponse.AddComment(ctx, agentqaTicketId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Merge other open issues of the same agent into this pending, in-progress, or resolved one. Their conversations move onto this issue, its priority is raised to the highest among them, and they are kept with status 'merged'. Requires viewer access to the issue's agent.
func (c *Client) Merge(ctx context.Context, agentqaTicketId string, request *agents.MergeAgentConversationTicketsRequestModel, opts ...option.RequestOption) (*types.AgentConversationTicketResponseModel, error) {
	response, err := c.WithRawResponse.Merge(ctx, agentqaTicketId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Append a turn-level comment to an agent issue. Requires viewer access to the issue's agent.
func (c *Client) AddTurnComment(ctx context.Context, agentqaTicketId string, request *agents.AddTurnCommentRequestModel, opts ...option.RequestOption) (*types.AgentConversationTicketResponseModel, error) {
	response, err := c.WithRawResponse.AddTurnComment(ctx, agentqaTicketId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
