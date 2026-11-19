package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	project "github.com/elevenlabs/elevenlabs-go/dubbing/project"
	transcript "github.com/elevenlabs/elevenlabs-go/dubbing/project/language/transcript"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Transcript *transcript.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
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

// List a project's language targets, cursor-paginated, each with signed output URLs once it has produced an output.
func (c *Client) List(ctx context.Context, projectId string, request *project.ListLanguageRequest, opts ...option.RequestOption) (*types.DubbingLanguageListResponse, error) {
	response, err := c.WithRawResponse.List(ctx, projectId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Add a language to dub a project into, and queue the dub.
//
// This is the call that produces dubbed audio, and it is billed per generation. The target is created `queued` and starts as soon as the project is `ready`, so it can be added at any point after the project is created. It inherits the project's dubbing model and cannot pick another.
//
// A project created with `webhook_ids` sends a `dubbing_language_completed` event carrying the output download URLs, so we recommend subscribing rather than polling this target to completion.
func (c *Client) Create(ctx context.Context, projectId string, request *project.BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost, opts ...option.RequestOption) (*types.DubbingLanguageResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, projectId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Full language-target detail. Once the target reports `completed`, `outputs` carries the signed download URLs. To learn when that happens, we recommend the project's `webhook_ids` subscription rather than polling this endpoint; fetch here when a delivered URL has expired, or to reconcile after an edit.
func (c *Client) Get(ctx context.Context, projectId string, languageId string, opts ...option.RequestOption) (*types.DubbingLanguageResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, projectId, languageId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a language target and its outputs, leaving the project and its other languages intact. This cannot be undone, and a dub already running is still billed.
func (c *Client) Delete(ctx context.Context, projectId string, languageId string, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Delete(ctx, projectId, languageId, opts...)
	return err
}
