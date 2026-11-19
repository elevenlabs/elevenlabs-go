package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentPlatformSettingsRequestModel struct {
	// Settings to evaluate an agent's performance.
	// Agents are evaluated against a set of criteria, with success being defined as meeting some combination of those criteria.
	Evaluation *EvaluationSettingsInput `json:"evaluation,omitempty" xml:"evaluation,omitempty"`
	Widget     *WidgetConfig            `json:"widget,omitempty" xml:"widget,omitempty"`
	// Data collection fields, keyed by ID. Each field is stored as an analysis item attached to the agent (see analysis_items).
	DataCollection map[string]*AnalysisProperty `json:"data_collection,omitempty"`
	// Scope per data collection item ID. Missing keys default to conversation scope.
	DataCollectionScopes map[string]AnalysisScope `json:"data_collection_scopes,omitempty"`
	// Analysis items (evaluation criteria and data collection fields) attached to the agent by reference. They are the source of truth for the agent's analysis. Writing evaluation.criteria or data_collection inline is still supported: a new entry creates an item, and changing an entry publishes a new version of its attached item, which every agent using that item picks up. To attach or detach an existing item, use the analysis item attachments endpoint. None means the agent has no attached items yet, and its inline fields are used as they are.
	AnalysisItems      *AgentAnalysisItemsInput                     `json:"analysis_items,omitempty" xml:"analysis_items,omitempty"`
	Overrides          *ConversationInitiationClientDataConfigInput `json:"overrides,omitempty" xml:"overrides,omitempty"`
	WorkspaceOverrides *AgentWorkspaceOverridesInput                `json:"workspace_overrides,omitempty" xml:"workspace_overrides,omitempty"`
	// Settings for agent testing configuration.
	Testing *AgentTestingSettings `json:"testing,omitempty" xml:"testing,omitempty"`
	// Whether the agent is archived
	Archived   *bool              `json:"archived,omitempty" xml:"archived,omitempty"`
	Guardrails *GuardrailsV1Input `json:"guardrails,omitempty" xml:"guardrails,omitempty"`
	// Language for all conversation analysis outputs (summaries, titles, evaluation rationales, data collection rationales). If not set, the language will be inferred from the conversation. Must be one of the supported conversation languages.
	SummaryLanguage *string `json:"summary_language,omitempty" xml:"summary_language,omitempty"`
	// When enabled, a conversation transcript is automatically translated to the viewer's application language when they open the transcript page. If not set or false, transcripts are shown in their original language unless the viewer manually selects a translation.
	AutoTranslateTranscriptToAppLanguage *bool                `json:"auto_translate_transcript_to_app_language,omitempty" xml:"auto_translate_transcript_to_app_language,omitempty"`
	Auth                                 *AuthSettings        `json:"auth,omitempty" xml:"auth,omitempty"`
	CallLimits                           *AgentCallLimits     `json:"call_limits,omitempty" xml:"call_limits,omitempty"`
	QueueingConfig                       *AgentQueueingConfig `json:"queueing_config,omitempty" xml:"queueing_config,omitempty"`
	Privacy                              *PrivacyConfigInput  `json:"privacy,omitempty" xml:"privacy,omitempty"`
	// The trust context in which the agent operates.
	//
	// UNKNOWN: not yet classified (existing agents created before this feature).
	// LOW: serves untrusted external participants (e.g. customer support, sales) —
	// outputs should be vetted and tool access scoped.
	// HIGH: serves the owner (e.g. personal assistant) — full tool access is appropriate.
	TrustContext *AgentTrustContext `json:"trust_context,omitempty" xml:"trust_context,omitempty"`
	AnalysisLlm  *Llm               `json:"analysis_llm,omitempty" xml:"analysis_llm,omitempty"`
	// Per-agent topic-discovery configuration. Cadence and analysis window are
	// managed internally; this only exposes the customer-facing on/off toggle.
	TopicDiscovery    *TopicDiscoverySettings    `json:"topic_discovery,omitempty" xml:"topic_discovery,omitempty"`
	SentimentAnalysis *SentimentAnalysisSettings `json:"sentiment_analysis,omitempty" xml:"sentiment_analysis,omitempty"`
	// Agent-level alerting configuration overriding workspace settings.
	Alerting        *AlertingSettings `json:"alerting,omitempty" xml:"alerting,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentPlatformSettingsRequestModel) GetEvaluation() (value *EvaluationSettingsInput) {
	if a == nil {
		return
	}
	return a.Evaluation
}

func (a *AgentPlatformSettingsRequestModel) GetWidget() (value *WidgetConfig) {
	if a == nil {
		return
	}
	return a.Widget
}

func (a *AgentPlatformSettingsRequestModel) GetDataCollection() (value map[string]*AnalysisProperty) {
	if a == nil {
		return
	}
	return a.DataCollection
}

func (a *AgentPlatformSettingsRequestModel) GetDataCollectionScopes() (value map[string]AnalysisScope) {
	if a == nil {
		return
	}
	return a.DataCollectionScopes
}

func (a *AgentPlatformSettingsRequestModel) GetAnalysisItems() (value *AgentAnalysisItemsInput) {
	if a == nil {
		return
	}
	return a.AnalysisItems
}

func (a *AgentPlatformSettingsRequestModel) GetOverrides() (value *ConversationInitiationClientDataConfigInput) {
	if a == nil {
		return
	}
	return a.Overrides
}

func (a *AgentPlatformSettingsRequestModel) GetWorkspaceOverrides() (value *AgentWorkspaceOverridesInput) {
	if a == nil {
		return
	}
	return a.WorkspaceOverrides
}

func (a *AgentPlatformSettingsRequestModel) GetTesting() (value *AgentTestingSettings) {
	if a == nil {
		return
	}
	return a.Testing
}

func (a *AgentPlatformSettingsRequestModel) GetArchived() (value *bool) {
	if a == nil {
		return
	}
	return a.Archived
}

func (a *AgentPlatformSettingsRequestModel) GetGuardrails() (value *GuardrailsV1Input) {
	if a == nil {
		return
	}
	return a.Guardrails
}

func (a *AgentPlatformSettingsRequestModel) GetSummaryLanguage() (value *string) {
	if a == nil {
		return
	}
	return a.SummaryLanguage
}

func (a *AgentPlatformSettingsRequestModel) GetAutoTranslateTranscriptToAppLanguage() (value *bool) {
	if a == nil {
		return
	}
	return a.AutoTranslateTranscriptToAppLanguage
}

func (a *AgentPlatformSettingsRequestModel) GetAuth() (value *AuthSettings) {
	if a == nil {
		return
	}
	return a.Auth
}

func (a *AgentPlatformSettingsRequestModel) GetCallLimits() (value *AgentCallLimits) {
	if a == nil {
		return
	}
	return a.CallLimits
}

func (a *AgentPlatformSettingsRequestModel) GetQueueingConfig() (value *AgentQueueingConfig) {
	if a == nil {
		return
	}
	return a.QueueingConfig
}

func (a *AgentPlatformSettingsRequestModel) GetPrivacy() (value *PrivacyConfigInput) {
	if a == nil {
		return
	}
	return a.Privacy
}

func (a *AgentPlatformSettingsRequestModel) GetTrustContext() (value *AgentTrustContext) {
	if a == nil {
		return
	}
	return a.TrustContext
}

func (a *AgentPlatformSettingsRequestModel) GetAnalysisLlm() (value *Llm) {
	if a == nil {
		return
	}
	return a.AnalysisLlm
}

func (a *AgentPlatformSettingsRequestModel) GetTopicDiscovery() (value *TopicDiscoverySettings) {
	if a == nil {
		return
	}
	return a.TopicDiscovery
}

func (a *AgentPlatformSettingsRequestModel) GetSentimentAnalysis() (value *SentimentAnalysisSettings) {
	if a == nil {
		return
	}
	return a.SentimentAnalysis
}

func (a *AgentPlatformSettingsRequestModel) GetAlerting() (value *AlertingSettings) {
	if a == nil {
		return
	}
	return a.Alerting
}

func (a *AgentPlatformSettingsRequestModel) SetEvaluation(evaluation *EvaluationSettingsInput) {
	a.Evaluation = evaluation
}

func (a *AgentPlatformSettingsRequestModel) SetWidget(widget *WidgetConfig) {
	a.Widget = widget
}

func (a *AgentPlatformSettingsRequestModel) SetDataCollection(dataCollection map[string]*AnalysisProperty) {
	a.DataCollection = dataCollection
}

func (a *AgentPlatformSettingsRequestModel) SetDataCollectionScopes(dataCollectionScopes map[string]AnalysisScope) {
	a.DataCollectionScopes = dataCollectionScopes
}

func (a *AgentPlatformSettingsRequestModel) SetAnalysisItems(analysisItems *AgentAnalysisItemsInput) {
	a.AnalysisItems = analysisItems
}

func (a *AgentPlatformSettingsRequestModel) SetOverrides(overrides *ConversationInitiationClientDataConfigInput) {
	a.Overrides = overrides
}

func (a *AgentPlatformSettingsRequestModel) SetWorkspaceOverrides(workspaceOverrides *AgentWorkspaceOverridesInput) {
	a.WorkspaceOverrides = workspaceOverrides
}

func (a *AgentPlatformSettingsRequestModel) SetTesting(testing *AgentTestingSettings) {
	a.Testing = testing
}

func (a *AgentPlatformSettingsRequestModel) SetArchived(archived *bool) {
	a.Archived = archived
}

func (a *AgentPlatformSettingsRequestModel) SetGuardrails(guardrails *GuardrailsV1Input) {
	a.Guardrails = guardrails
}

func (a *AgentPlatformSettingsRequestModel) SetSummaryLanguage(summaryLanguage *string) {
	a.SummaryLanguage = summaryLanguage
}

func (a *AgentPlatformSettingsRequestModel) SetAutoTranslateTranscriptToAppLanguage(autoTranslateTranscriptToAppLanguage *bool) {
	a.AutoTranslateTranscriptToAppLanguage = autoTranslateTranscriptToAppLanguage
}

func (a *AgentPlatformSettingsRequestModel) SetAuth(auth *AuthSettings) {
	a.Auth = auth
}

func (a *AgentPlatformSettingsRequestModel) SetCallLimits(callLimits *AgentCallLimits) {
	a.CallLimits = callLimits
}

func (a *AgentPlatformSettingsRequestModel) SetQueueingConfig(queueingConfig *AgentQueueingConfig) {
	a.QueueingConfig = queueingConfig
}

func (a *AgentPlatformSettingsRequestModel) SetPrivacy(privacy *PrivacyConfigInput) {
	a.Privacy = privacy
}

func (a *AgentPlatformSettingsRequestModel) SetTrustContext(trustContext *AgentTrustContext) {
	a.TrustContext = trustContext
}

func (a *AgentPlatformSettingsRequestModel) SetAnalysisLlm(analysisLlm *Llm) {
	a.AnalysisLlm = analysisLlm
}

func (a *AgentPlatformSettingsRequestModel) SetTopicDiscovery(topicDiscovery *TopicDiscoverySettings) {
	a.TopicDiscovery = topicDiscovery
}

func (a *AgentPlatformSettingsRequestModel) SetSentimentAnalysis(sentimentAnalysis *SentimentAnalysisSettings) {
	a.SentimentAnalysis = sentimentAnalysis
}

func (a *AgentPlatformSettingsRequestModel) SetAlerting(alerting *AlertingSettings) {
	a.Alerting = alerting
}

func (a *AgentPlatformSettingsRequestModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentPlatformSettingsRequestModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentPlatformSettingsRequestModel to string"
	}
	return string(jsonData)
}

func (a *AgentPlatformSettingsRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentPlatformSettingsRequestModel) MarshalJSON() ([]byte, error) {
	type embed AgentPlatformSettingsRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
