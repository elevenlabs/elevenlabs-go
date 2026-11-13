package client

import (
	context "context"
	analysisitemsclient "github.com/elevenlabs/elevenlabs-go/agents/analysisitems/client"
	analyticsclient "github.com/elevenlabs/elevenlabs-go/agents/analytics/client"
	batchcalls "github.com/elevenlabs/elevenlabs-go/agents/batchcalls"
	branches "github.com/elevenlabs/elevenlabs-go/agents/branches"
	conversationsclient "github.com/elevenlabs/elevenlabs-go/agents/conversations/client"
	dashboardclient "github.com/elevenlabs/elevenlabs-go/agents/dashboard/client"
	deployments "github.com/elevenlabs/elevenlabs-go/agents/deployments"
	drafts "github.com/elevenlabs/elevenlabs-go/agents/drafts"
	exotel "github.com/elevenlabs/elevenlabs-go/agents/exotel"
	holdaudio "github.com/elevenlabs/elevenlabs-go/agents/holdaudio"
	knowledgebaseclient "github.com/elevenlabs/elevenlabs-go/agents/knowledgebase/client"
	link "github.com/elevenlabs/elevenlabs-go/agents/link"
	llm "github.com/elevenlabs/elevenlabs-go/agents/llm"
	llmusage "github.com/elevenlabs/elevenlabs-go/agents/llmusage"
	mcpserversclient "github.com/elevenlabs/elevenlabs-go/agents/mcpservers/client"
	mergeproposals "github.com/elevenlabs/elevenlabs-go/agents/mergeproposals"
	outboundcampaignsclient "github.com/elevenlabs/elevenlabs-go/agents/outboundcampaigns/client"
	phonenumbers "github.com/elevenlabs/elevenlabs-go/agents/phonenumbers"
	proceduresclient "github.com/elevenlabs/elevenlabs-go/agents/procedures/client"
	realtime "github.com/elevenlabs/elevenlabs-go/agents/realtime"
	secrets "github.com/elevenlabs/elevenlabs-go/agents/secrets"
	settings "github.com/elevenlabs/elevenlabs-go/agents/settings"
	siptrunk "github.com/elevenlabs/elevenlabs-go/agents/siptrunk"
	summaries "github.com/elevenlabs/elevenlabs-go/agents/summaries"
	testsclient "github.com/elevenlabs/elevenlabs-go/agents/tests/client"
	toolsclient "github.com/elevenlabs/elevenlabs-go/agents/tools/client"
	triagetickets "github.com/elevenlabs/elevenlabs-go/agents/triagetickets"
	twilio "github.com/elevenlabs/elevenlabs-go/agents/twilio"
	users "github.com/elevenlabs/elevenlabs-go/agents/users"
	versions "github.com/elevenlabs/elevenlabs-go/agents/versions"
	whatsapp "github.com/elevenlabs/elevenlabs-go/agents/whatsapp"
	whatsappaccounts "github.com/elevenlabs/elevenlabs-go/agents/whatsappaccounts"
	widgetclient "github.com/elevenlabs/elevenlabs-go/agents/widget/client"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Conversations     *conversationsclient.Client
	Twilio            *twilio.Client
	Exotel            *exotel.Client
	Whatsapp          *whatsapp.Client
	Summaries         *summaries.Client
	Widget            *widgetclient.Client
	Link              *link.Client
	HoldAudio         *holdaudio.Client
	KnowledgeBase     *knowledgebaseclient.Client
	LlmUsage          *llmusage.Client
	Tests             *testsclient.Client
	Users             *users.Client
	TriageTickets     *triagetickets.Client
	PhoneNumbers      *phonenumbers.Client
	Llm               *llm.Client
	Analytics         *analyticsclient.Client
	Tools             *toolsclient.Client
	AnalysisItems     *analysisitemsclient.Client
	Settings          *settings.Client
	Dashboard         *dashboardclient.Client
	Secrets           *secrets.Client
	BatchCalls        *batchcalls.Client
	SipTrunk          *siptrunk.Client
	McpServers        *mcpserversclient.Client
	WhatsappAccounts  *whatsappaccounts.Client
	Branches          *branches.Client
	Versions          *versions.Client
	Deployments       *deployments.Client
	Drafts            *drafts.Client
	MergeProposals    *mergeproposals.Client
	OutboundCampaigns *outboundcampaignsclient.Client
	Procedures        *proceduresclient.Client
	Realtime          *realtime.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse:   NewRawClient(options),
		Conversations:     conversationsclient.NewClient(options),
		Twilio:            twilio.NewClient(options),
		Exotel:            exotel.NewClient(options),
		Whatsapp:          whatsapp.NewClient(options),
		Summaries:         summaries.NewClient(options),
		Widget:            widgetclient.NewClient(options),
		Link:              link.NewClient(options),
		HoldAudio:         holdaudio.NewClient(options),
		KnowledgeBase:     knowledgebaseclient.NewClient(options),
		LlmUsage:          llmusage.NewClient(options),
		Tests:             testsclient.NewClient(options),
		Users:             users.NewClient(options),
		TriageTickets:     triagetickets.NewClient(options),
		PhoneNumbers:      phonenumbers.NewClient(options),
		Llm:               llm.NewClient(options),
		Analytics:         analyticsclient.NewClient(options),
		Tools:             toolsclient.NewClient(options),
		AnalysisItems:     analysisitemsclient.NewClient(options),
		Settings:          settings.NewClient(options),
		Dashboard:         dashboardclient.NewClient(options),
		Secrets:           secrets.NewClient(options),
		BatchCalls:        batchcalls.NewClient(options),
		SipTrunk:          siptrunk.NewClient(options),
		McpServers:        mcpserversclient.NewClient(options),
		WhatsappAccounts:  whatsappaccounts.NewClient(options),
		Branches:          branches.NewClient(options),
		Versions:          versions.NewClient(options),
		Deployments:       deployments.NewClient(options),
		Drafts:            drafts.NewClient(options),
		MergeProposals:    mergeproposals.NewClient(options),
		OutboundCampaigns: outboundcampaignsclient.NewClient(options),
		Procedures:        proceduresclient.NewClient(options),
		Realtime:          realtime.NewClient(options),
		options:           options,
		baseURL:           options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:         options.HTTPClient,
				MaxAttempts:    options.MaxAttempts,
				DisableRetries: options.DisableRetries,
			},
		),
	}
}

// Create an agent from a config object
func (c *Client) Create(ctx context.Context, request *types.BodyCreateAgentV1ConvaiAgentsCreatePost, opts ...option.RequestOption) (*types.CreateAgentResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Retrieve config for an agent
func (c *Client) Get(ctx context.Context, agentId string, request *types.GetAgentsRequest, opts ...option.RequestOption) (*types.GetAgentResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Patches an Agent settings
func (c *Client) Update(ctx context.Context, agentId string, request *types.UpdateAgentRequest, opts ...option.RequestOption) (*types.GetAgentResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete an agent
func (c *Client) Delete(ctx context.Context, agentId string, request *types.DeleteAgentsRequest, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Delete(ctx, agentId, request, opts...)
	return err
}

// Returns a list of your agents and their metadata.
func (c *Client) List(ctx context.Context, request *types.ListAgentsRequest, opts ...option.RequestOption) (*types.GetAgentsPageResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new agent by duplicating an existing one. The new agent gets a copy of every unarchived branch of the source agent, each with the same name and description and starting from the latest version on that branch. The live traffic split is copied onto the new branches, with any share held by a branch that was not copied moved to main. Version history, drafts, archived branches, past traffic splits and experiments are not copied.
func (c *Client) Duplicate(ctx context.Context, agentId string, request *types.BodyDuplicateAgentV1ConvaiAgentsAgentIdDuplicatePost, opts ...option.RequestOption) (*types.CreateAgentResponseModel, error) {
	response, err := c.WithRawResponse.Duplicate(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Run selected tests on the agent with provided configuration. If the agent configuration is provided, it will be used to override default agent configuration.
func (c *Client) RunTests(ctx context.Context, agentId string, request *types.RunAgentTestsRequestModel, opts ...option.RequestOption) (*types.GetTestSuiteInvocationResponseModel, error) {
	response, err := c.WithRawResponse.RunTests(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Provides total size and other information of RAG indexes used by knowledgebase documents
func (c *Client) RagIndexOverview(ctx context.Context, opts ...option.RequestOption) (*types.RagIndexOverviewResponseModel, error) {
	response, err := c.WithRawResponse.RagIndexOverview(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Provides information about all RAG indexes of the specified knowledgebase document.
func (c *Client) GetDocumentRagIndexes(ctx context.Context, documentationId string, opts ...option.RequestOption) (*types.RagDocumentIndexesResponseModel, error) {
	response, err := c.WithRawResponse.GetDocumentRagIndexes(ctx, documentationId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete RAG index for the knowledgebase document.
func (c *Client) DeleteDocumentRagIndex(ctx context.Context, documentationId string, ragIndexId string, opts ...option.RequestOption) (*types.RagDocumentIndexResponseModel, error) {
	response, err := c.WithRawResponse.DeleteDocumentRagIndex(ctx, documentationId, ragIndexId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
