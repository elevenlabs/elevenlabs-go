package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	dubbing "github.com/elevenlabs/elevenlabs-go/dubbing"
	languageclient "github.com/elevenlabs/elevenlabs-go/dubbing/project/language/client"
	transcript "github.com/elevenlabs/elevenlabs-go/dubbing/project/transcript"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Language   *languageclient.Client
	Transcript *transcript.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Language:        languageclient.NewClient(options),
		Transcript:      transcript.NewClient(options),
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

// List the dubbing projects in your workspace that you can access, newest first, cursor-paginated. Listed projects carry no `language_ids`; fetch a project, or list its language targets, to see them.
func (c *Client) List(ctx context.Context, request *dubbing.ListProjectRequest, opts ...option.RequestOption) (*types.DubbingProjectListResponse, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a dubbing project from an uploaded file (`file`) or a source URL (`source_url`).
//
// Returns as soon as the project record exists, before the source has been fetched: the project starts `queued` and reaches `ready` once its source has been transcribed. Add a language target to it for each language you want, or pass `target_language` to queue the first one here.
//
// Creating a project incurs a minimum charge of one language: you are charged for at least one language's dubbing when the project is created, before any output exists. This charge prepays the first language target — the first target you add, or the one queued via `target_language`, consumes it — and each additional language is charged separately.
//
// Preparation can take minutes on a long source, so we recommend passing `webhook_ids` to be notified when the project turns `ready` or `failed`, rather than polling for it.
func (c *Client) Create(ctx context.Context, request *dubbing.BodyCreateDubbingProjectV1DubbingProjectPost, opts ...option.RequestOption) (*types.DubbingProjectResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Full project detail, including the IDs of every language target under it. To follow a project to `ready`, we recommend a `webhook_ids` subscription rather than polling this endpoint.
func (c *Client) Get(ctx context.Context, projectId string, opts ...option.RequestOption) (*types.DubbingProjectResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, projectId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a project, every language target under it, and their stored media and outputs. This cannot be undone, and a dub already running is still billed.
func (c *Client) Delete(ctx context.Context, projectId string, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Delete(ctx, projectId, opts...)
	return err
}
