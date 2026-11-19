package transferrules

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	receptionists "github.com/elevenlabs/elevenlabs-go/reception/receptionists"
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

// Add a phone forwarding rule to the named receptionist.
func (c *Client) Create(ctx context.Context, agentId string, request *receptionists.ReceptionCreateTransferRuleRequest, opts ...option.RequestOption) (*types.ReceptionTransferRule, error) {
	response, err := c.WithRawResponse.Create(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Replace an existing transfer rule on the named receptionist. Send every field: an omitted `post_dial_digits` clears it.
func (c *Client) Update(ctx context.Context, agentId string, ruleId string, request *receptionists.ReceptionUpdateTransferRuleRequest, opts ...option.RequestOption) (*types.ReceptionTransferRule, error) {
	response, err := c.WithRawResponse.Update(ctx, agentId, ruleId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a transfer rule from the named receptionist.
func (c *Client) Delete(ctx context.Context, agentId string, ruleId string, request *receptionists.DeleteTransferRulesRequest, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.Delete(ctx, agentId, ruleId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
