package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetConversationResponseModel struct {
	AgentId             string                             `json:"agent_id" xml:"agent_id" required:"true"`
	AgentName           *string                            `json:"agent_name,omitempty" xml:"agent_name,omitempty"`
	ConversationProduct *string                            `json:"conversation_product,omitempty" xml:"conversation_product,omitempty"`
	Status              GetConversationResponseModelStatus `json:"status" xml:"status" required:"true"`
	UserId              *string                            `json:"user_id,omitempty" xml:"user_id,omitempty"`
	BranchId            *string                            `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	// The ID of the agent version used for this conversation
	VersionId *string                                 `json:"version_id,omitempty" xml:"version_id,omitempty"`
	Metadata  *ConversationHistoryMetadataCommonModel `json:"metadata" xml:"metadata" required:"true"`
	Analysis  *ConversationHistoryAnalysisCommonModel `json:"analysis,omitempty" xml:"analysis,omitempty"`
	// Number of user transcript turns with text; tool-only turns are not counted. `null` for older conversations before this field whose transcript was deleted.
	UserMessageCount *int `json:"user_message_count,omitempty" xml:"user_message_count,omitempty"`
	// Number of agent transcript turns with text; tool-only turns are not counted. `null` for older conversations before this field whose transcript was deleted.
	AgentMessageCount                *int                                           `json:"agent_message_count,omitempty" xml:"agent_message_count,omitempty"`
	VisitedAgents                    []*VisitedAgentRef                             `json:"visited_agents,omitempty" xml:"visited_agents,omitempty"`
	ConversationInitiationClientData *ConversationInitiationClientDataRequestOutput `json:"conversation_initiation_client_data,omitempty" xml:"conversation_initiation_client_data,omitempty"`
	Environment                      *string                                        `json:"environment,omitempty" xml:"environment,omitempty"`
	ConversationId                   string                                         `json:"conversation_id" xml:"conversation_id" required:"true"`
	HasAudio                         bool                                           `json:"has_audio" xml:"has_audio" required:"true"`
	HasUserAudio                     bool                                           `json:"has_user_audio" xml:"has_user_audio" required:"true"`
	HasResponseAudio                 bool                                           `json:"has_response_audio" xml:"has_response_audio" required:"true"`
	HasAuxiliaryAudio                bool                                           `json:"has_auxiliary_audio" xml:"has_auxiliary_audio" required:"true"`
	Transcript                       []*ConversationHistoryTranscriptResponseModel  `json:"transcript" xml:"transcript" required:"true"`
	// Conversation tag ids assigned to this conversation.
	TagIds []string `json:"tag_ids,omitempty" xml:"tag_ids,omitempty"`
	// OpenTelemetry trace payload when the request uses format=opentelemetry; otherwise omitted.
	OtlpTraces      map[string]any `json:"otlp_traces,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetConversationResponseModel) GetAgentId() (value string) {
	if g == nil {
		return
	}
	return g.AgentId
}

func (g *GetConversationResponseModel) GetAgentName() (value *string) {
	if g == nil {
		return
	}
	return g.AgentName
}

func (g *GetConversationResponseModel) GetConversationProduct() (value *string) {
	if g == nil {
		return
	}
	return g.ConversationProduct
}

func (g *GetConversationResponseModel) GetStatus() (value GetConversationResponseModelStatus) {
	if g == nil {
		return
	}
	return g.Status
}

func (g *GetConversationResponseModel) GetUserId() (value *string) {
	if g == nil {
		return
	}
	return g.UserId
}

func (g *GetConversationResponseModel) GetBranchId() (value *string) {
	if g == nil {
		return
	}
	return g.BranchId
}

func (g *GetConversationResponseModel) GetVersionId() (value *string) {
	if g == nil {
		return
	}
	return g.VersionId
}

func (g *GetConversationResponseModel) GetMetadata() (value *ConversationHistoryMetadataCommonModel) {
	if g == nil {
		return
	}
	return g.Metadata
}

func (g *GetConversationResponseModel) GetAnalysis() (value *ConversationHistoryAnalysisCommonModel) {
	if g == nil {
		return
	}
	return g.Analysis
}

func (g *GetConversationResponseModel) GetUserMessageCount() (value *int) {
	if g == nil {
		return
	}
	return g.UserMessageCount
}

func (g *GetConversationResponseModel) GetAgentMessageCount() (value *int) {
	if g == nil {
		return
	}
	return g.AgentMessageCount
}

func (g *GetConversationResponseModel) GetVisitedAgents() (value []*VisitedAgentRef) {
	if g == nil {
		return
	}
	return g.VisitedAgents
}

func (g *GetConversationResponseModel) GetConversationInitiationClientData() (value *ConversationInitiationClientDataRequestOutput) {
	if g == nil {
		return
	}
	return g.ConversationInitiationClientData
}

func (g *GetConversationResponseModel) GetEnvironment() (value *string) {
	if g == nil {
		return
	}
	return g.Environment
}

func (g *GetConversationResponseModel) GetConversationId() (value string) {
	if g == nil {
		return
	}
	return g.ConversationId
}

func (g *GetConversationResponseModel) GetHasAudio() (value bool) {
	if g == nil {
		return
	}
	return g.HasAudio
}

func (g *GetConversationResponseModel) GetHasUserAudio() (value bool) {
	if g == nil {
		return
	}
	return g.HasUserAudio
}

func (g *GetConversationResponseModel) GetHasResponseAudio() (value bool) {
	if g == nil {
		return
	}
	return g.HasResponseAudio
}

func (g *GetConversationResponseModel) GetHasAuxiliaryAudio() (value bool) {
	if g == nil {
		return
	}
	return g.HasAuxiliaryAudio
}

func (g *GetConversationResponseModel) GetTranscript() (value []*ConversationHistoryTranscriptResponseModel) {
	if g == nil {
		return
	}
	return g.Transcript
}

func (g *GetConversationResponseModel) GetTagIds() (value []string) {
	if g == nil {
		return
	}
	return g.TagIds
}

func (g *GetConversationResponseModel) GetOtlpTraces() (value map[string]any) {
	if g == nil {
		return
	}
	return g.OtlpTraces
}

func (g *GetConversationResponseModel) SetAgentId(agentId string) {
	g.AgentId = agentId
}

func (g *GetConversationResponseModel) SetAgentName(agentName *string) {
	g.AgentName = agentName
}

func (g *GetConversationResponseModel) SetConversationProduct(conversationProduct *string) {
	g.ConversationProduct = conversationProduct
}

func (g *GetConversationResponseModel) SetStatus(status GetConversationResponseModelStatus) {
	g.Status = status
}

func (g *GetConversationResponseModel) SetUserId(userId *string) {
	g.UserId = userId
}

func (g *GetConversationResponseModel) SetBranchId(branchId *string) {
	g.BranchId = branchId
}

func (g *GetConversationResponseModel) SetVersionId(versionId *string) {
	g.VersionId = versionId
}

func (g *GetConversationResponseModel) SetMetadata(metadata *ConversationHistoryMetadataCommonModel) {
	g.Metadata = metadata
}

func (g *GetConversationResponseModel) SetAnalysis(analysis *ConversationHistoryAnalysisCommonModel) {
	g.Analysis = analysis
}

func (g *GetConversationResponseModel) SetUserMessageCount(userMessageCount *int) {
	g.UserMessageCount = userMessageCount
}

func (g *GetConversationResponseModel) SetAgentMessageCount(agentMessageCount *int) {
	g.AgentMessageCount = agentMessageCount
}

func (g *GetConversationResponseModel) SetVisitedAgents(visitedAgents []*VisitedAgentRef) {
	g.VisitedAgents = visitedAgents
}

func (g *GetConversationResponseModel) SetConversationInitiationClientData(conversationInitiationClientData *ConversationInitiationClientDataRequestOutput) {
	g.ConversationInitiationClientData = conversationInitiationClientData
}

func (g *GetConversationResponseModel) SetEnvironment(environment *string) {
	g.Environment = environment
}

func (g *GetConversationResponseModel) SetConversationId(conversationId string) {
	g.ConversationId = conversationId
}

func (g *GetConversationResponseModel) SetHasAudio(hasAudio bool) {
	g.HasAudio = hasAudio
}

func (g *GetConversationResponseModel) SetHasUserAudio(hasUserAudio bool) {
	g.HasUserAudio = hasUserAudio
}

func (g *GetConversationResponseModel) SetHasResponseAudio(hasResponseAudio bool) {
	g.HasResponseAudio = hasResponseAudio
}

func (g *GetConversationResponseModel) SetHasAuxiliaryAudio(hasAuxiliaryAudio bool) {
	g.HasAuxiliaryAudio = hasAuxiliaryAudio
}

func (g *GetConversationResponseModel) SetTranscript(transcript []*ConversationHistoryTranscriptResponseModel) {
	g.Transcript = transcript
}

func (g *GetConversationResponseModel) SetTagIds(tagIds []string) {
	g.TagIds = tagIds
}

func (g *GetConversationResponseModel) SetOtlpTraces(otlpTraces map[string]any) {
	g.OtlpTraces = otlpTraces
}

func (g *GetConversationResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetConversationResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetConversationResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetConversationResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetConversationResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetConversationResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
