package agentbuilder

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	agentsclient "github.com/elevenlabs/elevenlabs-go/agents/client"
	"github.com/elevenlabs/elevenlabs-go/core"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/option"
)

// AgentsClient is the agents client, reached as client.Agents: the generated client, embedded,
// plus Push and Pull for agents written as code. client/client.go is frozen in .fernignore to
// construct it.
//
// Every generated method (Create, Get, Update, List, Tools, Branches...) comes from the embedded
// *agentsclient.Client.
type AgentsClient struct {
	*agentsclient.Client

	options *core.RequestOptions
	caller  *internal.Caller
}

// NewAgentsClient wraps a generated agents client built from options.
func NewAgentsClient(options *core.RequestOptions) *AgentsClient {
	return &AgentsClient{
		Client:  agentsclient.NewClient(options),
		options: options,
		caller: internal.NewCaller(&internal.CallerParams{
			Client:         options.HTTPClient,
			MaxAttempts:    options.MaxAttempts,
			DisableRetries: options.DisableRetries,
		}),
	}
}

// call makes one request to the agents or tools API with a wire-format body and returns the
// response as JSON text.
//
// Push and Pull send and read API-format JSON as it is, so settings this SDK doesn't model reach
// the API and come back intact. The generated methods can't: their response types drop unknown
// fields when marshaled again, and their requests refuse extra body properties that repeat a
// typed field. The request is made the way the generated methods make it, with the client's and
// the call's options (base URL, API key, headers, retries, HTTP client).
func (c *AgentsClient) call(ctx context.Context, method, path string, query url.Values, body any, opts []option.RequestOption) (json.RawMessage, error) {
	options := core.NewRequestOptions(opts...)
	endpointURL := internal.ResolveBaseURL(options.BaseURL, c.options.BaseURL, "https://api.elevenlabs.io") + path
	if len(query) > 0 {
		endpointURL += "?" + query.Encode()
	}
	headers := internal.MergeHeaders(c.options.ToHeader(), options.ToHeader())
	if body != nil {
		headers.Add("Content-Type", "application/json")
	}
	var response json.RawMessage
	_, err := c.caller.Call(ctx, &internal.CallParams{
		URL:             endpointURL,
		Method:          method,
		Headers:         headers,
		MaxAttempts:     options.MaxAttempts,
		DisableRetries:  options.DisableRetries,
		BodyProperties:  options.BodyProperties,
		QueryParameters: options.QueryParameters,
		Client:          options.HTTPClient,
		Request:         body,
		Response:        &response,
	})
	return response, err
}

func (c *AgentsClient) get(ctx context.Context, path string, query url.Values, opts []option.RequestOption) (json.RawMessage, error) {
	return c.call(ctx, http.MethodGet, path, query, nil, opts)
}

func toolPath(id string) string  { return internal.EncodeURL("/v1/convai/tools/%v", id) }
func agentPath(id string) string { return internal.EncodeURL("/v1/convai/agents/%v", id) }

func branchQuery(branchID string) url.Values {
	if branchID == "" {
		return nil
	}
	return url.Values{"branch_id": {branchID}}
}
