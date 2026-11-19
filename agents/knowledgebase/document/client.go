package document

import (
	context "context"
	knowledgebase "github.com/elevenlabs/elevenlabs-go/agents/knowledgebase"
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

// Update the source file of a file document. The document name, content, and metadata are updated to reflect the new file. Any manual content edits will be overwritten.
func (c *Client) UpdateFile(ctx context.Context, documentationId string, request *knowledgebase.BodyUpdateFileDocumentV1ConvaiKnowledgeBaseDocumentationIdUpdateFilePatch, opts ...option.RequestOption) (*knowledgebase.UpdateFileDocumentResponse, error) {
	response, err := c.WithRawResponse.UpdateFile(ctx, documentationId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Manually refresh a URL document by re-fetching its content from the source URL.
func (c *Client) Refresh(ctx context.Context, documentationId string, opts ...option.RequestOption) (*knowledgebase.RefreshDocumentResponse, error) {
	response, err := c.WithRawResponse.Refresh(ctx, documentationId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// In case the document is not RAG indexed, it triggers rag indexing task, otherwise it just returns the current status.
func (c *Client) ComputeRagIndex(ctx context.Context, documentationId string, request *knowledgebase.RagIndexRequestModel, opts ...option.RequestOption) (*types.RagDocumentIndexResponseModel, error) {
	response, err := c.WithRawResponse.ComputeRagIndex(ctx, documentationId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
