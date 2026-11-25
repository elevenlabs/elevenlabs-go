package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentResponseCorrectionAgentResponseCorrectionEvent struct {
	OriginalAgentResponse  string `json:"original_agent_response" xml:"original_agent_response" required:"true"`
	CorrectedAgentResponse string `json:"corrected_agent_response" xml:"corrected_agent_response" required:"true"`
	EventId                int    `json:"event_id" xml:"event_id" required:"true"`
	ResponseId             string `json:"response_id" xml:"response_id" required:"true"`
	extraProperties        map[string]interface{}
}

func (a *AgentResponseCorrectionAgentResponseCorrectionEvent) GetOriginalAgentResponse() (value string) {
	if a == nil {
		return
	}
	return a.OriginalAgentResponse
}

func (a *AgentResponseCorrectionAgentResponseCorrectionEvent) GetCorrectedAgentResponse() (value string) {
	if a == nil {
		return
	}
	return a.CorrectedAgentResponse
}

func (a *AgentResponseCorrectionAgentResponseCorrectionEvent) GetEventId() (value int) {
	if a == nil {
		return
	}
	return a.EventId
}

func (a *AgentResponseCorrectionAgentResponseCorrectionEvent) GetResponseId() (value string) {
	if a == nil {
		return
	}
	return a.ResponseId
}

func (a *AgentResponseCorrectionAgentResponseCorrectionEvent) SetOriginalAgentResponse(originalAgentResponse string) {
	a.OriginalAgentResponse = originalAgentResponse
}

func (a *AgentResponseCorrectionAgentResponseCorrectionEvent) SetCorrectedAgentResponse(correctedAgentResponse string) {
	a.CorrectedAgentResponse = correctedAgentResponse
}

func (a *AgentResponseCorrectionAgentResponseCorrectionEvent) SetEventId(eventId int) {
	a.EventId = eventId
}

func (a *AgentResponseCorrectionAgentResponseCorrectionEvent) SetResponseId(responseId string) {
	a.ResponseId = responseId
}

func (a *AgentResponseCorrectionAgentResponseCorrectionEvent) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentResponseCorrectionAgentResponseCorrectionEvent) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentResponseCorrectionAgentResponseCorrectionEvent to string"
	}
	return string(jsonData)
}

func (a *AgentResponseCorrectionAgentResponseCorrectionEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentResponseCorrectionAgentResponseCorrectionEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentResponseCorrectionAgentResponseCorrectionEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentResponseCorrectionAgentResponseCorrectionEvent) MarshalJSON() ([]byte, error) {
	type embed AgentResponseCorrectionAgentResponseCorrectionEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
