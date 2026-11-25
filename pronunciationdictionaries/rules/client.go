package rules

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	pronunciationdictionaries "github.com/elevenlabs/elevenlabs-go/pronunciationdictionaries"
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

// Replaces all existing rules on the pronunciation dictionary with the provided ones.
func (c *Client) Set(ctx context.Context, pronunciationDictionaryId string, request *pronunciationdictionaries.BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost, opts ...option.RequestOption) (*types.PronunciationDictionaryRulesResponseModel, error) {
	response, err := c.WithRawResponse.Set(ctx, pronunciationDictionaryId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Add rules to the pronunciation dictionary. If a rule with the same string_to_replace already exists, it will be replaced.
func (c *Client) Add(ctx context.Context, pronunciationDictionaryId string, request *pronunciationdictionaries.PronunciationDictionary, opts ...option.RequestOption) (*types.PronunciationDictionaryRulesResponseModel, error) {
	response, err := c.WithRawResponse.Add(ctx, pronunciationDictionaryId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Remove rules from the pronunciation dictionary
func (c *Client) Remove(ctx context.Context, pronunciationDictionaryId string, request *pronunciationdictionaries.RemovePronunciationDictionaryRulesRequest, opts ...option.RequestOption) (*types.PronunciationDictionaryRulesResponseModel, error) {
	response, err := c.WithRawResponse.Remove(ctx, pronunciationDictionaryId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
