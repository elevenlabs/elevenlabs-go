package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentSimulatedChatTestResponseModel struct {
	SimulatedConversation []*ConversationHistoryTranscriptResponseModel `json:"simulated_conversation" xml:"simulated_conversation" required:"true"`
	Analysis              *ConversationHistoryAnalysisCommonModel       `json:"analysis" xml:"analysis" required:"true"`
	extraProperties       map[string]interface{}
}

func (a *AgentSimulatedChatTestResponseModel) GetSimulatedConversation() (value []*ConversationHistoryTranscriptResponseModel) {
	if a == nil {
		return
	}
	return a.SimulatedConversation
}

func (a *AgentSimulatedChatTestResponseModel) GetAnalysis() (value *ConversationHistoryAnalysisCommonModel) {
	if a == nil {
		return
	}
	return a.Analysis
}

func (a *AgentSimulatedChatTestResponseModel) SetSimulatedConversation(simulatedConversation []*ConversationHistoryTranscriptResponseModel) {
	a.SimulatedConversation = simulatedConversation
}

func (a *AgentSimulatedChatTestResponseModel) SetAnalysis(analysis *ConversationHistoryAnalysisCommonModel) {
	a.Analysis = analysis
}

func (a *AgentSimulatedChatTestResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentSimulatedChatTestResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentSimulatedChatTestResponseModel to string"
	}
	return string(jsonData)
}

func (a *AgentSimulatedChatTestResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentSimulatedChatTestResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentSimulatedChatTestResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentSimulatedChatTestResponseModel) MarshalJSON() ([]byte, error) {
	type embed AgentSimulatedChatTestResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
