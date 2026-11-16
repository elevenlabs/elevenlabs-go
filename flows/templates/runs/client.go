package runs

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	templates "github.com/elevenlabs/elevenlabs-go/flows/templates"
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

// List this template's runs created through this API, newest first.
func (c *Client) List(ctx context.Context, templateId string, request *templates.ListRunsRequest, opts ...option.RequestOption) (*types.TemplateRunListResponse, error) {
	response, err := c.WithRawResponse.List(ctx, templateId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Start a run of a flows template. Pass `version_id` to pin a snapshot, or omit it / pass `latest` to run the latest published version. Set input values under `inputs`, keyed by input port id. The response is the run in its initial state, with every output already listed under its port id in `outputs`. Include `webhook` to receive a `flows_template_run` event carrying the finished run once its `status` is `completed` or `failed`; this is the recommended way to wait. Without one, fetch `GET /v1/flows/templates/{template_id}/runs/{run_id}` at a modest interval until the `status` is terminal.
func (c *Client) Create(ctx context.Context, templateId string, request *templates.TemplateRunCreateRequest, opts ...option.RequestOption) (*types.TemplateRunResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, templateId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Retrieve a template run: its `status`, rolled up from its outputs, and each output's own status and download URL once completed.
func (c *Client) Get(ctx context.Context, templateId string, runId string, opts ...option.RequestOption) (*types.TemplateRunResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, templateId, runId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
