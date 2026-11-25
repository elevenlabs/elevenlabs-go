package snapshots

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	projects "github.com/elevenlabs/elevenlabs-go/studio/projects"
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

// Retrieves a list of snapshots for a Studio project.
func (c *Client) List(ctx context.Context, projectId string, opts ...option.RequestOption) (*types.ProjectSnapshotsResponse, error) {
	response, err := c.WithRawResponse.List(ctx, projectId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Returns the project snapshot.
func (c *Client) Get(ctx context.Context, projectId string, projectSnapshotId string, opts ...option.RequestOption) (*types.ProjectSnapshotExtendedResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, projectId, projectSnapshotId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Stream the audio from a Studio project snapshot.
func (c *Client) Stream(ctx context.Context, projectId string, projectSnapshotId string, request *projects.BodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost, opts ...option.RequestOption) (string, error) {
	response, err := c.WithRawResponse.Stream(ctx, projectId, projectSnapshotId, request, opts...)
	if err != nil {
		return "", err
	}
	return response.Body, nil
}

// Returns a compressed archive of the Studio project's audio.
func (c *Client) StreamArchive(ctx context.Context, projectId string, projectSnapshotId string, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.StreamArchive(ctx, projectId, projectSnapshotId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
