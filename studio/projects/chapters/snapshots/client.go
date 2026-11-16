package snapshots

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	chapters "github.com/elevenlabs/elevenlabs-go/studio/projects/chapters"
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

// Gets information about all the snapshots of a chapter. Each snapshot can be downloaded as audio. Whenever a chapter is converted a snapshot will automatically be created.
func (c *Client) List(ctx context.Context, projectId string, chapterId string, opts ...option.RequestOption) (*types.ChapterSnapshotsResponse, error) {
	response, err := c.WithRawResponse.List(ctx, projectId, chapterId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Returns the chapter snapshot.
func (c *Client) Get(ctx context.Context, projectId string, chapterId string, chapterSnapshotId string, opts ...option.RequestOption) (*types.ChapterSnapshotExtendedResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, projectId, chapterId, chapterSnapshotId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Stream the audio from a chapter snapshot. Use `GET /v1/studio/projects/{project_id}/chapters/{chapter_id}/snapshots` to return the snapshots of a chapter.
func (c *Client) Stream(ctx context.Context, projectId string, chapterId string, chapterSnapshotId string, request *chapters.AudioV1StudioProjectsProjectIdChaptersChapterIdSnapshotsChapterSnapshotIdStreamPost, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Stream(ctx, projectId, chapterId, chapterSnapshotId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
