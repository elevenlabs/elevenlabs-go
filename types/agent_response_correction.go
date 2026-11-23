package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentResponseCorrection struct {
	AgentResponseCorrectionEvent *AgentResponseCorrectionAgentResponseCorrectionEvent `json:"agent_response_correction_event" xml:"agent_response_correction_event" required:"true"`
	Type                         *string                                              `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties              map[string]interface{}
}

func (a *AgentResponseCorrection) GetAgentResponseCorrectionEvent() (value *AgentResponseCorrectionAgentResponseCorrectionEvent) {
	if a == nil {
		return
	}
	return a.AgentResponseCorrectionEvent
}

func (a *AgentResponseCorrection) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AgentResponseCorrection) SetAgentResponseCorrectionEvent(agentResponseCorrectionEvent *AgentResponseCorrectionAgentResponseCorrectionEvent) {
	a.AgentResponseCorrectionEvent = agentResponseCorrectionEvent
}

func (a *AgentResponseCorrection) SetType(type_ *string) {
	a.Type = type_
}

func (a *AgentResponseCorrection) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentResponseCorrection) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentResponseCorrection to string"
	}
	return string(jsonData)
}

func (a *AgentResponseCorrection) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentResponseCorrection
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentResponseCorrection(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentResponseCorrection) MarshalJSON() ([]byte, error) {
	type embed AgentResponseCorrection
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
