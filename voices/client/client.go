package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	accents "github.com/elevenlabs/elevenlabs-go/voices/accents"
	collections "github.com/elevenlabs/elevenlabs-go/voices/collections"
	ivc "github.com/elevenlabs/elevenlabs-go/voices/ivc"
	pvcclient "github.com/elevenlabs/elevenlabs-go/voices/pvc/client"
	samplesclient "github.com/elevenlabs/elevenlabs-go/voices/samples/client"
	settings "github.com/elevenlabs/elevenlabs-go/voices/settings"
)

type Client struct {
	WithRawResponse *RawClient

	Samples     *samplesclient.Client
	Collections *collections.Client
	Settings    *settings.Client
	Accents     *accents.Client
	Ivc         *ivc.Client
	Pvc         *pvcclient.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Samples:         samplesclient.NewClient(options),
		Collections:     collections.NewClient(options),
		Settings:        settings.NewClient(options),
		Accents:         accents.NewClient(options),
		Ivc:             ivc.NewClient(options),
		Pvc:             pvcclient.NewClient(options),
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

// Returns metadata about a specific voice.
func (c *Client) Get(ctx context.Context, voiceId string, request *types.GetVoicesRequest, opts ...option.RequestOption) (*types.Voice, error) {
	response, err := c.WithRawResponse.Get(ctx, voiceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Deletes a voice by its ID.
func (c *Client) Delete(ctx context.Context, voiceId string, opts ...option.RequestOption) (*types.DeleteVoiceResponseModel, error) {
	response, err := c.WithRawResponse.Delete(ctx, voiceId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Edit a voice created by you.
func (c *Client) Update(ctx context.Context, voiceId string, request *types.EditVoiceRequest, opts ...option.RequestOption) (*types.EditVoiceResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, voiceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Gets a list of all available voices for a user with search, filtering and pagination.
func (c *Client) Search(ctx context.Context, request *types.SearchVoicesRequest, opts ...option.RequestOption) (*types.GetVoicesV2Response, error) {
	response, err := c.WithRawResponse.Search(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Replicates an Instant Voice Clone or Voice Design voice to a workspace in a different data residency. The target workspace must belong to the same consolidated billing group. The user must have VOICES_WRITE in the source workspace, and be an admin on the source voice. Human users (i.e. not service accounts) must also have VOICES_WRITE in the target workspace. This endpoint is available on the central environment only.
func (c *Client) ReplicateToIsolatedEnvironment(ctx context.Context, voiceId string, request *types.ReplicateVoiceToIsolatedEnvironmentRequestModel, opts ...option.RequestOption) (*types.ReplicateVoiceToIsolatedEnvironmentResponseModel, error) {
	response, err := c.WithRawResponse.ReplicateToIsolatedEnvironment(ctx, voiceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Add a shared voice to your collection of Voices
func (c *Client) Share(ctx context.Context, publicUserId string, voiceId string, request *types.BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost, opts ...option.RequestOption) (*types.AddVoiceResponseModel, error) {
	response, err := c.WithRawResponse.Share(ctx, publicUserId, voiceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Returns a list of shared voices similar to the provided audio sample. If neither similarity_threshold nor top_k is provided, we will apply default values.
func (c *Client) FindSimilarVoices(ctx context.Context, listSimilarVoicesRequest *types.ListSimilarVoicesRequest, opts ...option.RequestOption) (*types.GetLibraryVoicesResponse, error) {
	response, err := c.WithRawResponse.FindSimilarVoices(ctx, listSimilarVoicesRequest, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Retrieves a list of shared voices.
func (c *Client) GetShared(ctx context.Context, request *types.GetSharedVoicesRequest, opts ...option.RequestOption) (*types.GetLibraryVoicesResponse, error) {
	response, err := c.WithRawResponse.GetShared(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
