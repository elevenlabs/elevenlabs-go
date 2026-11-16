package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	analyticsclient "github.com/elevenlabs/elevenlabs-go/workspace/analytics/client"
	auditlogs "github.com/elevenlabs/elevenlabs-go/workspace/auditlogs"
	authconnections "github.com/elevenlabs/elevenlabs-go/workspace/authconnections"
	groupsclient "github.com/elevenlabs/elevenlabs-go/workspace/groups/client"
	invites "github.com/elevenlabs/elevenlabs-go/workspace/invites"
	members "github.com/elevenlabs/elevenlabs-go/workspace/members"
	resources "github.com/elevenlabs/elevenlabs-go/workspace/resources"
	usage "github.com/elevenlabs/elevenlabs-go/workspace/usage"
)

type Client struct {
	WithRawResponse *RawClient

	AuditLogs       *auditlogs.Client
	AuthConnections *authconnections.Client
	Groups          *groupsclient.Client
	Invites         *invites.Client
	Members         *members.Client
	Resources       *resources.Client
	Usage           *usage.Client
	Analytics       *analyticsclient.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		AuditLogs:       auditlogs.NewClient(options),
		AuthConnections: authconnections.NewClient(options),
		Groups:          groupsclient.NewClient(options),
		Invites:         invites.NewClient(options),
		Members:         members.NewClient(options),
		Resources:       resources.NewClient(options),
		Usage:           usage.NewClient(options),
		Analytics:       analyticsclient.NewClient(options),
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

// Set the workspace-wide Third-Party Disabling policy. When set, it forces, for every API key in the workspace, whether the holder of a key (potentially a third party who found it) may disable it via the self-disable endpoint or when it leaks publicly — overriding each key's own setting. Pass `true` to allow it for all keys, `false` to forbid it for all keys, or `null` to clear the override so per-key values and the plan default apply again. Workspace admins only.
func (c *Client) SetThirdPartyDisablingPolicy(ctx context.Context, request *types.BodySetWorkspaceThirdPartyDisablingPolicyV1WorkspacesApiKeysThirdPartyDisablingPost, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.SetThirdPartyDisablingPolicy(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
