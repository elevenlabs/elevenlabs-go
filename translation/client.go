package translation

import (
	core "github.com/elevenlabs/elevenlabs-go/core"
)

// Client opens connections to the translation channel.
type Client struct {
	options *core.RequestOptions
}

// NewClient returns a client that shares the given request options.
func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		options: options,
	}
}
