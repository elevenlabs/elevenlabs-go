package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentPlatformSettingsResponseModel struct {
	// Settings to evaluate an agent's performance.
	// Agents are evaluated against a set of criteria, with success being defined as meeting some combination of those criteria.
	Evaluation *EvaluationSettingsOutput `json:"evaluation,omitempty" xml:"evaluation,omitempty"`
	Widget     *WidgetConfig             `json:"widget,omitempty" xml:"widget,omitempty"`
	// Data collection fields, keyed by ID. Each field is stored as an analysis item attached to the agent (see analysis_items).
	DataCollection map[string]*AnalysisProperty `json:"data_collection,omitempty"`
	// Scope per data collection item ID. Missing keys default to conversation scope.
	DataCollectionScopes map[string]AnalysisScope `json:"data_collection_scopes,omitempty"`
	// Analysis items (evaluation criteria and data collection fields) attached to the agent by reference. They are the source of truth for the agent's analysis. Writing evaluation.criteria or data_collection inline is still supported: a new entry creates an item, and changing an entry publishes a new version of its attached item, which every agent using that item picks up. To attach or detach an existing item, use the analysis item attachments endpoint. None means the agent has no attached items yet, and its inline fields are used as they are.
	AnalysisItems      *AgentAnalysisItemsOutput                     `json:"analysis_items,omitempty" xml:"analysis_items,omitempty"`
	Overrides          *ConversationInitiationClientDataConfigOutput `json:"overrides,omitempty" xml:"overrides,omitempty"`
	WorkspaceOverrides *AgentWorkspaceOverridesOutput                `json:"workspace_overrides,omitempty" xml:"workspace_overrides,omitempty"`
	// Settings for agent testing configuration.
	Testing *AgentTestingSettings `json:"testing,omitempty" xml:"testing,omitempty"`
	// Whether the agent is archived
	Archived   *bool               `json:"archived,omitempty" xml:"archived,omitempty"`
	Guardrails *GuardrailsV1Output `json:"guardrails,omitempty" xml:"guardrails,omitempty"`
	// Language for all conversation analysis outputs (summaries, titles, evaluation rationales, data collection rationales). If not set, the language will be inferred from the conversation. Must be one of the supported conversation languages.
	SummaryLanguage *string `json:"summary_language,omitempty" xml:"summary_language,omitempty"`
	// When enabled, a conversation transcript is automatically translated to the viewer's application language when they open the transcript page. If not set or false, transcripts are shown in their original language unless the viewer manually selects a translation.
	AutoTranslateTranscriptToAppLanguage *bool                `json:"auto_translate_transcript_to_app_language,omitempty" xml:"auto_translate_transcript_to_app_language,omitempty"`
	Auth                                 *AuthSettings        `json:"auth,omitempty" xml:"auth,omitempty"`
	CallLimits                           *AgentCallLimits     `json:"call_limits,omitempty" xml:"call_limits,omitempty"`
	QueueingConfig                       *AgentQueueingConfig `json:"queueing_config,omitempty" xml:"queueing_config,omitempty"`
	Privacy                              *PrivacyConfigOutput `json:"privacy,omitempty" xml:"privacy,omitempty"`
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
	Alerting        *AlertingSettingsResponse `json:"alerting,omitempty" xml:"alerting,omitempty"`
	Safety          *SafetyResponseModel      `json:"safety,omitempty" xml:"safety,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentPlatformSettingsResponseModel) GetEvaluation() (value *EvaluationSettingsOutput) {
	if a == nil {
		return
	}
	return a.Evaluation
}

func (a *AgentPlatformSettingsResponseModel) GetWidget() (value *WidgetConfig) {
	if a == nil {
		return
	}
	return a.Widget
}

func (a *AgentPlatformSettingsResponseModel) GetDataCollection() (value map[string]*AnalysisProperty) {
	if a == nil {
		return
	}
	return a.DataCollection
}

func (a *AgentPlatformSettingsResponseModel) GetDataCollectionScopes() (value map[string]AnalysisScope) {
	if a == nil {
		return
	}
	return a.DataCollectionScopes
}

func (a *AgentPlatformSettingsResponseModel) GetAnalysisItems() (value *AgentAnalysisItemsOutput) {
	if a == nil {
		return
	}
	return a.AnalysisItems
}

func (a *AgentPlatformSettingsResponseModel) GetOverrides() (value *ConversationInitiationClientDataConfigOutput) {
	if a == nil {
		return
	}
	return a.Overrides
}

func (a *AgentPlatformSettingsResponseModel) GetWorkspaceOverrides() (value *AgentWorkspaceOverridesOutput) {
	if a == nil {
		return
	}
	return a.WorkspaceOverrides
}

func (a *AgentPlatformSettingsResponseModel) GetTesting() (value *AgentTestingSettings) {
	if a == nil {
		return
	}
	return a.Testing
}

func (a *AgentPlatformSettingsResponseModel) GetArchived() (value *bool) {
	if a == nil {
		return
	}
	return a.Archived
}

func (a *AgentPlatformSettingsResponseModel) GetGuardrails() (value *GuardrailsV1Output) {
	if a == nil {
		return
	}
	return a.Guardrails
}

func (a *AgentPlatformSettingsResponseModel) GetSummaryLanguage() (value *string) {
	if a == nil {
		return
	}
	return a.SummaryLanguage
}

func (a *AgentPlatformSettingsResponseModel) GetAutoTranslateTranscriptToAppLanguage() (value *bool) {
	if a == nil {
		return
	}
	return a.AutoTranslateTranscriptToAppLanguage
}

func (a *AgentPlatformSettingsResponseModel) GetAuth() (value *AuthSettings) {
	if a == nil {
		return
	}
	return a.Auth
}

func (a *AgentPlatformSettingsResponseModel) GetCallLimits() (value *AgentCallLimits) {
	if a == nil {
		return
	}
	return a.CallLimits
}

func (a *AgentPlatformSettingsResponseModel) GetQueueingConfig() (value *AgentQueueingConfig) {
	if a == nil {
		return
	}
	return a.QueueingConfig
}

func (a *AgentPlatformSettingsResponseModel) GetPrivacy() (value *PrivacyConfigOutput) {
	if a == nil {
		return
	}
	return a.Privacy
}

func (a *AgentPlatformSettingsResponseModel) GetTrustContext() (value *AgentTrustContext) {
	if a == nil {
		return
	}
	return a.TrustContext
}

func (a *AgentPlatformSettingsResponseModel) GetAnalysisLlm() (value *Llm) {
	if a == nil {
		return
	}
	return a.AnalysisLlm
}

func (a *AgentPlatformSettingsResponseModel) GetTopicDiscovery() (value *TopicDiscoverySettings) {
	if a == nil {
		return
	}
	return a.TopicDiscovery
}

func (a *AgentPlatformSettingsResponseModel) GetSentimentAnalysis() (value *SentimentAnalysisSettings) {
	if a == nil {
		return
	}
	return a.SentimentAnalysis
}

func (a *AgentPlatformSettingsResponseModel) GetAlerting() (value *AlertingSettingsResponse) {
	if a == nil {
		return
	}
	return a.Alerting
}

func (a *AgentPlatformSettingsResponseModel) GetSafety() (value *SafetyResponseModel) {
	if a == nil {
		return
	}
	return a.Safety
}

func (a *AgentPlatformSettingsResponseModel) SetEvaluation(evaluation *EvaluationSettingsOutput) {
	a.Evaluation = evaluation
}

func (a *AgentPlatformSettingsResponseModel) SetWidget(widget *WidgetConfig) {
	a.Widget = widget
}

func (a *AgentPlatformSettingsResponseModel) SetDataCollection(dataCollection map[string]*AnalysisProperty) {
	a.DataCollection = dataCollection
}

func (a *AgentPlatformSettingsResponseModel) SetDataCollectionScopes(dataCollectionScopes map[string]AnalysisScope) {
	a.DataCollectionScopes = dataCollectionScopes
}

func (a *AgentPlatformSettingsResponseModel) SetAnalysisItems(analysisItems *AgentAnalysisItemsOutput) {
	a.AnalysisItems = analysisItems
}

func (a *AgentPlatformSettingsResponseModel) SetOverrides(overrides *ConversationInitiationClientDataConfigOutput) {
	a.Overrides = overrides
}

func (a *AgentPlatformSettingsResponseModel) SetWorkspaceOverrides(workspaceOverrides *AgentWorkspaceOverridesOutput) {
	a.WorkspaceOverrides = workspaceOverrides
}

func (a *AgentPlatformSettingsResponseModel) SetTesting(testing *AgentTestingSettings) {
	a.Testing = testing
}

func (a *AgentPlatformSettingsResponseModel) SetArchived(archived *bool) {
	a.Archived = archived
}

func (a *AgentPlatformSettingsResponseModel) SetGuardrails(guardrails *GuardrailsV1Output) {
	a.Guardrails = guardrails
}

func (a *AgentPlatformSettingsResponseModel) SetSummaryLanguage(summaryLanguage *string) {
	a.SummaryLanguage = summaryLanguage
}

func (a *AgentPlatformSettingsResponseModel) SetAutoTranslateTranscriptToAppLanguage(autoTranslateTranscriptToAppLanguage *bool) {
	a.AutoTranslateTranscriptToAppLanguage = autoTranslateTranscriptToAppLanguage
}

func (a *AgentPlatformSettingsResponseModel) SetAuth(auth *AuthSettings) {
	a.Auth = auth
}

func (a *AgentPlatformSettingsResponseModel) SetCallLimits(callLimits *AgentCallLimits) {
	a.CallLimits = callLimits
}

func (a *AgentPlatformSettingsResponseModel) SetQueueingConfig(queueingConfig *AgentQueueingConfig) {
	a.QueueingConfig = queueingConfig
}

func (a *AgentPlatformSettingsResponseModel) SetPrivacy(privacy *PrivacyConfigOutput) {
	a.Privacy = privacy
}

func (a *AgentPlatformSettingsResponseModel) SetTrustContext(trustContext *AgentTrustContext) {
	a.TrustContext = trustContext
}

func (a *AgentPlatformSettingsResponseModel) SetAnalysisLlm(analysisLlm *Llm) {
	a.AnalysisLlm = analysisLlm
}

func (a *AgentPlatformSettingsResponseModel) SetTopicDiscovery(topicDiscovery *TopicDiscoverySettings) {
	a.TopicDiscovery = topicDiscovery
}

func (a *AgentPlatformSettingsResponseModel) SetSentimentAnalysis(sentimentAnalysis *SentimentAnalysisSettings) {
	a.SentimentAnalysis = sentimentAnalysis
}

func (a *AgentPlatformSettingsResponseModel) SetAlerting(alerting *AlertingSettingsResponse) {
	a.Alerting = alerting
}

func (a *AgentPlatformSettingsResponseModel) SetSafety(safety *SafetyResponseModel) {
	a.Safety = safety
}

func (a *AgentPlatformSettingsResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentPlatformSettingsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentPlatformSettingsResponseModel to string"
	}
	return string(jsonData)
}

func (a *AgentPlatformSettingsResponseModel) UnmarshalJSON(data []byte) error {
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

func (a *AgentPlatformSettingsResponseModel) MarshalJSON() ([]byte, error) {
	type embed AgentPlatformSettingsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
