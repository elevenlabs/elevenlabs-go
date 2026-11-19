package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentResponseClientEventData struct {
	AgentResponse   string                     `json:"agent_response" xml:"agent_response" required:"true"`
	EventId         int                        `json:"event_id" xml:"event_id" required:"true"`
	ResponseId      string                     `json:"response_id" xml:"response_id" required:"true"`
	InResponseToIds []string                   `json:"in_response_to_ids,omitempty" xml:"in_response_to_ids,omitempty"`
	ResponseType    *AgentResponseType         `json:"response_type,omitempty" xml:"response_type,omitempty"`
	Attachments     []*AgentResponseAttachment `json:"attachments,omitempty" xml:"attachments,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentResponseClientEventData) GetAgentResponse() (value string) {
	if a == nil {
		return
	}
	return a.AgentResponse
}

func (a *AgentResponseClientEventData) GetEventId() (value int) {
	if a == nil {
		return
	}
	return a.EventId
}

func (a *AgentResponseClientEventData) GetResponseId() (value string) {
	if a == nil {
		return
	}
	return a.ResponseId
}

func (a *AgentResponseClientEventData) GetInResponseToIds() (value []string) {
	if a == nil {
		return
	}
	return a.InResponseToIds
}

func (a *AgentResponseClientEventData) GetResponseType() (value *AgentResponseType) {
	if a == nil {
		return
	}
	return a.ResponseType
}

func (a *AgentResponseClientEventData) GetAttachments() (value []*AgentResponseAttachment) {
	if a == nil {
		return
	}
	return a.Attachments
}

func (a *AgentResponseClientEventData) SetAgentResponse(agentResponse string) {
	a.AgentResponse = agentResponse
}

func (a *AgentResponseClientEventData) SetEventId(eventId int) {
	a.EventId = eventId
}

func (a *AgentResponseClientEventData) SetResponseId(responseId string) {
	a.ResponseId = responseId
}

func (a *AgentResponseClientEventData) SetInResponseToIds(inResponseToIds []string) {
	a.InResponseToIds = inResponseToIds
}

func (a *AgentResponseClientEventData) SetResponseType(responseType *AgentResponseType) {
	a.ResponseType = responseType
}

func (a *AgentResponseClientEventData) SetAttachments(attachments []*AgentResponseAttachment) {
	a.Attachments = attachments
}

func (a *AgentResponseClientEventData) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentResponseClientEventData) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentResponseClientEventData to string"
	}
	return string(jsonData)
}

func (a *AgentResponseClientEventData) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentResponseClientEventData
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentResponseClientEventData(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentResponseClientEventData) MarshalJSON() ([]byte, error) {
	type embed AgentResponseClientEventData
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
