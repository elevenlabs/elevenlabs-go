package folders

import (
	context "context"
	tests "github.com/elevenlabs/elevenlabs-go/agents/tests"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
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

// Creates a folder for organizing agent tests.
func (c *Client) Create(ctx context.Context, request *tests.BodyCreateAgentTestFolderV1ConvaiAgentTestingFoldersPost, opts ...option.RequestOption) (*types.CreateAgentTestFolderResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Gets an agent test folder by ID, including its folder path.
func (c *Client) Get(ctx context.Context, folderId string, opts ...option.RequestOption) (*types.GetAgentTestFolderResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, folderId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Updates an agent test folder. Currently only supports updating the folder name.
func (c *Client) Update(ctx context.Context, folderId string, request *tests.BodyUpdateAgentTestFolderV1ConvaiAgentTestingFoldersFolderIdPatch, opts ...option.RequestOption) (*types.GetAgentTestFolderResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, folderId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Deletes an agent test folder by ID. Use force=true to delete a non-empty folder and all its contents.
func (c *Client) Delete(ctx context.Context, folderId string, request *tests.DeleteFoldersRequest, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Delete(ctx, folderId, request, opts...)
	return err
}
