package collections

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	voices "github.com/elevenlabs/elevenlabs-go/voices"
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

// Returns all the voice collections the user has access to, sorted by title.
func (c *Client) List(ctx context.Context, opts ...option.RequestOption) (*types.GetUserVoiceCollectionsResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Creates a new voice collection. The caller becomes its admin.
func (c *Client) Create(ctx context.Context, request *voices.BodyCreateVoiceCollectionV1VoicesCollectionsPost, opts ...option.RequestOption) (*types.UserVoiceCollectionResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Returns a voice collection the user has access to.
func (c *Client) Get(ctx context.Context, collectionId string, opts ...option.RequestOption) (*types.UserVoiceCollectionResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, collectionId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Updates the title and icon of a voice collection. Fields that are omitted are left unchanged. Requires editor access to the collection.
func (c *Client) Update(ctx context.Context, collectionId string, request *voices.BodyUpdateVoiceCollectionV1VoicesCollectionsCollectionIdPatch, opts ...option.RequestOption) (*types.UserVoiceCollectionResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, collectionId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Deletes a voice collection. The voices in it are not deleted. Requires admin access to the collection.
func (c *Client) Delete(ctx context.Context, collectionId string, opts ...option.RequestOption) (*types.DeleteUserVoiceCollectionResponseModel, error) {
	response, err := c.WithRawResponse.Delete(ctx, collectionId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Adds a voice to a voice collection. Adding a voice that is already in the collection succeeds without changes. Requires editor access to the collection.
func (c *Client) AddVoice(ctx context.Context, collectionId string, voiceId string, opts ...option.RequestOption) (*types.AddVoiceToUserVoiceCollectionResponseModel, error) {
	response, err := c.WithRawResponse.AddVoice(ctx, collectionId, voiceId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Removes a voice from a voice collection. The voice itself is not deleted. Requires editor access to the collection.
func (c *Client) RemoveVoice(ctx context.Context, collectionId string, voiceId string, opts ...option.RequestOption) (*types.RemoveVoiceFromUserVoiceCollectionResponseModel, error) {
	response, err := c.WithRawResponse.RemoveVoice(ctx, collectionId, voiceId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
