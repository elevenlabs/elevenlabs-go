package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	projects "github.com/elevenlabs/elevenlabs-go/studio/projects"
	snapshots "github.com/elevenlabs/elevenlabs-go/studio/projects/chapters/snapshots"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Snapshots *snapshots.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Snapshots:       snapshots.NewClient(options),
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

// Returns a list of a Studio project's chapters.
func (c *Client) List(ctx context.Context, projectId string, opts ...option.RequestOption) (*types.GetChaptersResponse, error) {
	response, err := c.WithRawResponse.List(ctx, projectId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Creates a new chapter either as blank or from a URL.
func (c *Client) Create(ctx context.Context, projectId string, request *projects.BodyCreateChapterV1StudioProjectsProjectIdChaptersPost, opts ...option.RequestOption) (*types.AddChapterResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, projectId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Returns information about a specific chapter.
func (c *Client) Get(ctx context.Context, projectId string, chapterId string, opts ...option.RequestOption) (*types.ChapterWithContentResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, projectId, chapterId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Updates a chapter.
func (c *Client) Update(ctx context.Context, projectId string, chapterId string, request *projects.BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost, opts ...option.RequestOption) (*types.EditChapterResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, projectId, chapterId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Deletes a chapter.
func (c *Client) Delete(ctx context.Context, projectId string, chapterId string, opts ...option.RequestOption) (*types.DeleteChapterResponseModel, error) {
	response, err := c.WithRawResponse.Delete(ctx, projectId, chapterId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Starts conversion of a specific chapter.
func (c *Client) Convert(ctx context.Context, projectId string, chapterId string, opts ...option.RequestOption) (*types.ConvertChapterResponseModel, error) {
	response, err := c.WithRawResponse.Convert(ctx, projectId, chapterId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
