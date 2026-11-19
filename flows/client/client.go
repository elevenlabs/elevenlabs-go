package client

import (
	core "github.com/elevenlabs/elevenlabs-go/core"
	image "github.com/elevenlabs/elevenlabs-go/flows/image"
	templatesclient "github.com/elevenlabs/elevenlabs-go/flows/templates/client"
	texttospeech "github.com/elevenlabs/elevenlabs-go/flows/texttospeech"
	video "github.com/elevenlabs/elevenlabs-go/flows/video"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
)

type Client struct {
	Video        *video.Client
	Image        *image.Client
	TextToSpeech *texttospeech.Client
	Templates    *templatesclient.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		Video:        video.NewClient(options),
		Image:        image.NewClient(options),
		TextToSpeech: texttospeech.NewClient(options),
		Templates:    templatesclient.NewClient(options),
		options:      options,
		baseURL:      options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:         options.HTTPClient,
				MaxAttempts:    options.MaxAttempts,
				DisableRetries: options.DisableRetries,
			},
		),
	}
}
