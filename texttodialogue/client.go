package texttodialogue

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	"io"
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

// Converts a list of text and voice ID pairs into speech (dialogue) and returns audio.
func (c *Client) Convert(ctx context.Context, request *types.BodyTextToDialogueMultiVoiceV1TextToDialoguePost, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Convert(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Converts a list of text and voice ID pairs into speech (dialogue) and returns an audio stream.
func (c *Client) Stream(ctx context.Context, request *types.BodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPost, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Stream(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Converts a list of text and voice ID pairs into speech (dialogue) and returns a stream of JSON blobs containing audio as a base64 encoded string and timestamps
func (c *Client) StreamWithTimestamps(ctx context.Context, request *types.BodyTextToDialogueStreamWithTimestamps, opts ...option.RequestOption) (*core.Stream[types.StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel], error) {
	options := core.NewRequestOptions(opts...)
	endpointURL := internal.ResolveBaseURL(options.BaseURL, c.baseURL, "https://api.elevenlabs.io") + "/v1/text-to-dialogue/stream/with-timestamps"
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
	streamer := internal.NewStreamer[types.StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel](c.caller)
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
		Prefix:                     "",
		Delimiter:                  "\n",
		Terminator:                 "",
		EventDiscriminator:         "",
		Resumable:                  false,
		Format:                     core.StreamFormatEmpty,
		RawText:                    false,
		Request:                    request,
		ErrorDecoder:               nil,
	})
}

// Generate dialogue from text with precise character-level timing information for audio-text synchronization.
func (c *Client) FullWithTimestamps(ctx context.Context, request *types.BodyTextToDialogueFullWithTimestamps, opts ...option.RequestOption) (*types.AudioWithTimestampsAndVoiceSegmentsResponseModel, error) {
	response, err := c.WithRawResponse.FullWithTimestamps(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
