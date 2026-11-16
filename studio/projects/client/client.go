package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	studio "github.com/elevenlabs/elevenlabs-go/studio"
	chaptersclient "github.com/elevenlabs/elevenlabs-go/studio/projects/chapters/client"
	content "github.com/elevenlabs/elevenlabs-go/studio/projects/content"
	pronunciationdictionaries "github.com/elevenlabs/elevenlabs-go/studio/projects/pronunciationdictionaries"
	snapshots "github.com/elevenlabs/elevenlabs-go/studio/projects/snapshots"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	PronunciationDictionaries *pronunciationdictionaries.Client
	Content                   *content.Client
	Snapshots                 *snapshots.Client
	Chapters                  *chaptersclient.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse:           NewRawClient(options),
		PronunciationDictionaries: pronunciationdictionaries.NewClient(options),
		Content:                   content.NewClient(options),
		Snapshots:                 snapshots.NewClient(options),
		Chapters:                  chaptersclient.NewClient(options),
		options:                   options,
		baseURL:                   options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:         options.HTTPClient,
				MaxAttempts:    options.MaxAttempts,
				DisableRetries: options.DisableRetries,
			},
		),
	}
}

// Returns a list of your Studio projects with metadata.
func (c *Client) List(ctx context.Context, opts ...option.RequestOption) (*types.GetProjectsResponse, error) {
	response, err := c.WithRawResponse.List(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Creates a new Studio project, it can be either initialized as blank, from a document or from a URL.
func (c *Client) Create(ctx context.Context, bodyCreateStudioProjectV1StudioProjectsPost *studio.BodyCreateStudioProjectV1StudioProjectsPost, opts ...option.RequestOption) (*types.AddProjectResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, bodyCreateStudioProjectV1StudioProjectsPost, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Returns information about a specific Studio project. This endpoint returns more detailed information about a project than `GET /v1/studio`.
func (c *Client) Get(ctx context.Context, projectId string, request *studio.GetProjectsRequest, opts ...option.RequestOption) (*types.ProjectExtendedResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, projectId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Updates the specified Studio project by setting the values of the parameters passed.
func (c *Client) Update(ctx context.Context, projectId string, request *studio.BodyUpdateStudioProjectV1StudioProjectsProjectIdPost, opts ...option.RequestOption) (*types.EditProjectResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, projectId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Deletes a Studio project.
func (c *Client) Delete(ctx context.Context, projectId string, opts ...option.RequestOption) (*types.DeleteProjectResponseModel, error) {
	response, err := c.WithRawResponse.Delete(ctx, projectId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Starts conversion of a Studio project and all of its chapters.
func (c *Client) Convert(ctx context.Context, projectId string, opts ...option.RequestOption) (*types.ConvertProjectResponseModel, error) {
	response, err := c.WithRawResponse.Convert(ctx, projectId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Returns a list of chapter IDs that have muted tracks in a project.
func (c *Client) GetMutedTracks(ctx context.Context, projectId string, opts ...option.RequestOption) (*types.ProjectMutedTracksResponseModel, error) {
	response, err := c.WithRawResponse.GetMutedTracks(ctx, projectId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
