package client

import (
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	invites "github.com/elevenlabs/elevenlabs-go/reception/team/invites"
	members "github.com/elevenlabs/elevenlabs-go/reception/team/members"
)

type Client struct {
	Members *members.Client
	Invites *invites.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		Members: members.NewClient(options),
		Invites: invites.NewClient(options),
		options: options,
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:         options.HTTPClient,
				MaxAttempts:    options.MaxAttempts,
				DisableRetries: options.DisableRetries,
			},
		),
	}
}
