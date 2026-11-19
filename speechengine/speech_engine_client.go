// Hand-written (listed in .fernignore): the client the root client exposes as
// client.SpeechEngine. It embeds the generated Client and shadows Create, Get, Update and
// Duplicate so they return a SpeechEngineResource, like SpeechEngineClientWrapper in the
// TypeScript SDK. client/client.go is frozen in .fernignore to construct it; see there.

package speechengine

import (
	context "context"

	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

// SpeechEngineClient is the Speech Engine client, reached as client.SpeechEngine.
//
// List, Delete, Upstream and WithRawResponse come from the embedded generated Client. Create,
// Get, Update and Duplicate return a *SpeechEngineResource that can serve the engine's inbound
// connections; the plain generated methods stay reachable through the embedded Client field
// (client.SpeechEngine.Client.Create).
type SpeechEngineClient struct {
	*Client
}

// NewSpeechEngineClient wraps a generated Client built from options.
func NewSpeechEngineClient(options *core.RequestOptions) *SpeechEngineClient {
	return &SpeechEngineClient{Client: NewClient(options)}
}

// apiKey returns the key that verifies inbound connections: a per-call option.WithApiKey wins
// over the client's key, matching the key the API call itself used.
func (c *SpeechEngineClient) apiKey(opts []option.RequestOption) string {
	if call := core.NewRequestOptions(opts...); call.ApiKey != nil {
		return *call.ApiKey
	}
	if c.options != nil && c.options.ApiKey != nil {
		return *c.options.ApiKey
	}
	return ""
}

// Create creates a speech engine and returns it as a SpeechEngineResource.
func (c *SpeechEngineClient) Create(ctx context.Context, request *types.CreateSpeechEngineRequest, opts ...option.RequestOption) (*SpeechEngineResource, error) {
	response, err := c.Client.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return NewSpeechEngineResource(response.GetSpeechEngineId(), c.apiKey(opts), response), nil
}

// Get retrieves a speech engine and returns it as a SpeechEngineResource.
func (c *SpeechEngineClient) Get(ctx context.Context, speechEngineId string, opts ...option.RequestOption) (*SpeechEngineResource, error) {
	response, err := c.Client.Get(ctx, speechEngineId, opts...)
	if err != nil {
		return nil, err
	}
	return NewSpeechEngineResource(speechEngineId, c.apiKey(opts), response), nil
}

// Update partially updates a speech engine and returns it as a SpeechEngineResource.
func (c *SpeechEngineClient) Update(ctx context.Context, speechEngineId string, request *types.UpdateSpeechEngineRequest, opts ...option.RequestOption) (*SpeechEngineResource, error) {
	response, err := c.Client.Update(ctx, speechEngineId, request, opts...)
	if err != nil {
		return nil, err
	}
	return NewSpeechEngineResource(speechEngineId, c.apiKey(opts), response), nil
}

// Duplicate copies a speech engine and returns the copy as a SpeechEngineResource.
func (c *SpeechEngineClient) Duplicate(ctx context.Context, speechEngineId string, request types.DuplicateSpeechEngineRequest, opts ...option.RequestOption) (*SpeechEngineResource, error) {
	response, err := c.Client.Duplicate(ctx, speechEngineId, request, opts...)
	if err != nil {
		return nil, err
	}
	return NewSpeechEngineResource(response.GetSpeechEngineId(), c.apiKey(opts), response), nil
}

// Resource returns a SpeechEngineResource for an existing engine without an API call, so a
// server can start without fetching the engine first. Its Config is nil.
func (c *SpeechEngineClient) Resource(speechEngineId string) *SpeechEngineResource {
	return NewSpeechEngineResource(speechEngineId, c.apiKey(nil), nil)
}
