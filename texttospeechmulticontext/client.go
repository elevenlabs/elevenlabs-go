package texttospeechmulticontext

import (
	core "github.com/elevenlabs/elevenlabs-go/core"
)

// Client opens connections to the text_to_speech_multi_context channel.
type Client struct {
	options *core.RequestOptions
}

// NewClient returns a client that shares the given request options.
func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		options: options,
	}
}
