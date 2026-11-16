package invites

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	team "github.com/elevenlabs/elevenlabs-go/reception/team"
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

// List the workspace's pending invites.
func (c *Client) List(ctx context.Context, request *team.ListInvitesRequest, opts ...option.RequestOption) ([]*types.ReceptionTeamInviteResponse, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Email an invite to join the workspace. A pending invite takes a seat, so this is refused once members and invites reach the plan's team limit.
func (c *Client) Create(ctx context.Context, request *team.CreateTeamInviteRequest, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Cancel a pending invite, freeing its seat.
func (c *Client) Cancel(ctx context.Context, request *team.BodySmbCancelTeamInvite, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.Cancel(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
