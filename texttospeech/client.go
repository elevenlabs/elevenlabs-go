package texttospeech

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

// Converts text into speech using a voice of your choice and returns audio.
func (c *Client) Convert(ctx context.Context, voiceId string, request *types.BodyTextToSpeechFull, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Convert(ctx, voiceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Generate speech from text with precise character-level timing information for audio-text synchronization.
func (c *Client) ConvertWithTimestamps(ctx context.Context, voiceId string, request *types.BodyTextToSpeechFullWithTimestamps, opts ...option.RequestOption) (*types.AudioWithTimestampsResponse, error) {
	response, err := c.WithRawResponse.ConvertWithTimestamps(ctx, voiceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Converts text into speech using a voice of your choice and returns audio as an audio stream.
func (c *Client) Stream(ctx context.Context, voiceId string, request *types.StreamTextToSpeechRequest, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Stream(ctx, voiceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Converts text into speech using a voice of your choice and returns a stream of JSONs containing audio as a base64 encoded string together with information on when which character was spoken.
func (c *Client) StreamWithTimestamps(ctx context.Context, voiceId string, request *types.StreamTextToSpeechWithTimestampsRequest, opts ...option.RequestOption) (*core.Stream[types.StreamingAudioChunkWithTimestampsResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(options.BaseURL, internal.ResolveEnvironmentBaseURL(options.Environment, "Base"), c.baseURL, internal.ResolveEnvironmentBaseURL(c.options.Environment, "Base"), "https://api.elevenlabs.io")
	endpointURL := internal.EncodeURL(baseURL+"/v1/text-to-speech/%v/stream/with-timestamps", voiceId)
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
	streamer := internal.NewStreamer[types.StreamingAudioChunkWithTimestampsResponse](c.caller)
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
