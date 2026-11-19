package client

import (
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	singleuse "github.com/elevenlabs/elevenlabs-go/tokens/singleuse"
)

type Client struct {
	SingleUse *singleuse.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		SingleUse: singleuse.NewClient(options),
		options:   options,
		baseURL:   options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:         options.HTTPClient,
				MaxAttempts:    options.MaxAttempts,
				DisableRetries: options.DisableRetries,
			},
		),
	}
}
