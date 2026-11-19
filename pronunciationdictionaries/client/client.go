package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	rules "github.com/elevenlabs/elevenlabs-go/pronunciationdictionaries/rules"
	types "github.com/elevenlabs/elevenlabs-go/types"
	"io"
)

type Client struct {
	WithRawResponse *RawClient

	Rules *rules.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Rules:           rules.NewClient(options),
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

// Creates a new pronunciation dictionary from a lexicon .PLS file
func (c *Client) CreateFromFile(ctx context.Context, request *types.AddPronunciationDictionaryRequest, opts ...option.RequestOption) (*types.AddPronunciationDictionaryResponseModel, error) {
	response, err := c.WithRawResponse.CreateFromFile(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Creates a new pronunciation dictionary from provided rules.
func (c *Client) CreateFromRules(ctx context.Context, request *types.BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost, opts ...option.RequestOption) (*types.AddPronunciationDictionaryResponseModel, error) {
	response, err := c.WithRawResponse.CreateFromRules(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get metadata for a pronunciation dictionary
func (c *Client) Get(ctx context.Context, pronunciationDictionaryId string, opts ...option.RequestOption) (*types.GetPronunciationDictionaryWithRulesResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, pronunciationDictionaryId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Partially update the pronunciation dictionary without changing the version
func (c *Client) Update(ctx context.Context, pronunciationDictionaryId string, request *types.BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch, opts ...option.RequestOption) (*types.GetPronunciationDictionaryMetadataResponse, error) {
	response, err := c.WithRawResponse.Update(ctx, pronunciationDictionaryId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a PLS file with a pronunciation dictionary version rules
func (c *Client) Download(ctx context.Context, dictionaryId string, versionId string, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Download(ctx, dictionaryId, versionId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a list of the pronunciation dictionaries you have access to and their metadata
func (c *Client) List(ctx context.Context, request *types.ListPronunciationDictionariesRequest, opts ...option.RequestOption) (*types.GetPronunciationDictionariesMetadataResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
