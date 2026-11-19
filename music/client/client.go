package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	compositionplan "github.com/elevenlabs/elevenlabs-go/music/compositionplan"
	finetunes "github.com/elevenlabs/elevenlabs-go/music/finetunes"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	"io"
)

type Client struct {
	WithRawResponse *RawClient

	CompositionPlan *compositionplan.Client
	Finetunes       *finetunes.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		CompositionPlan: compositionplan.NewClient(options),
		Finetunes:       finetunes.NewClient(options),
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

// Generate background music from one or more video files. Videos are combined in order. Optional description and style tags influence the generated music.
func (c *Client) VideoToMusic(ctx context.Context, bodyVideoToMusicV1MusicVideoToMusicPost *types.BodyVideoToMusicV1MusicVideoToMusicPost, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.VideoToMusic(ctx, bodyVideoToMusicV1MusicVideoToMusicPost, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Compose a song from a prompt or a composition plan.
func (c *Client) Compose(ctx context.Context, request *types.BodyComposeMusicV1MusicPost, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Compose(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Compose a song from a prompt or a composition plan.
func (c *Client) ComposeDetailed(ctx context.Context, request *types.BodyComposeMusicWithADetailedResponseV1MusicDetailedPost, opts ...option.RequestOption) (*types.ComposeDetailedMusicResponse, error) {
	response, err := c.WithRawResponse.ComposeDetailed(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Stream a song and its detailed metadata using Server-Sent Events (SSE).
func (c *Client) ComposeDetailedStream(ctx context.Context, request *types.BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost, opts ...option.RequestOption) (*core.Stream[string], error) {
	options := core.NewRequestOptions(opts...)
	endpointURL := internal.ResolveBaseURL(options.BaseURL, c.baseURL, "https://api.elevenlabs.io") + "/v1/music/detailed/stream"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Set("Accept", "text/event-stream")
	streamer := internal.NewStreamer[string](c.caller)
	return streamer.StreamWithReconnect(ctx, &internal.StreamParams{
		URL:                        endpointURL,
		Method:                     "POST",
		Headers:                    headers,
		MaxAttempts:                options.MaxAttempts,
		DisableRetries:             options.DisableRetries,
		BodyProperties:             options.BodyProperties,
		QueryParameters:            options.QueryParameters,
		Client:                     options.HTTPClient,
		MaxBufSize:                 options.MaxBufSize,
		MaxStreamReconnectAttempts: options.MaxStreamReconnectAttempts,
		DisableStreamReconnection:  options.DisableStreamReconnection,
		Prefix:                     internal.DefaultSSEDataPrefix,
		Delimiter:                  "",
		Terminator:                 "",
		EventDiscriminator:         "",
		Resumable:                  false,
		Format:                     core.StreamFormatSSE,
		RawText:                    true,
		Request:                    request,
		ErrorDecoder:               nil,
	})
}

// Stream a composed song from a prompt or a composition plan.
func (c *Client) Stream(ctx context.Context, request *types.BodyStreamComposedMusicV1MusicStreamPost, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Stream(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Upload a music file to be later used for inpainting. Price for uploading is the same as the one for song generation. All uploaded content gets inspected for copyright infringement. If copyrighted content is detected, half of the request cost is still charged.
func (c *Client) Upload(ctx context.Context, request *types.BodyUploadMusicV1MusicUploadPost, opts ...option.RequestOption) (*types.MusicUploadResponse, error) {
	response, err := c.WithRawResponse.Upload(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Separate an audio file into individual stems. This endpoint might have high latency, depending on the length of the audio file.
func (c *Client) SeparateStems(ctx context.Context, bodyStemSeparationV1MusicStemSeparationPost *types.BodyStemSeparationV1MusicStemSeparationPost, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.SeparateStems(ctx, bodyStemSeparationV1MusicStemSeparationPost, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
