package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentSummaryResponseModel struct {
	// The ID of the agent
	AgentId string `json:"agent_id" xml:"agent_id" required:"true"`
	// The name of the agent
	Name string `json:"name" xml:"name" required:"true"`
	// Voice ID assigned to this agent
	VoiceId string `json:"voice_id" xml:"voice_id" required:"true"`
	// Agent tags used to categorize the agent
	Tags []string `json:"tags" xml:"tags" required:"true"`
	// The creation time of the agent in unix seconds
	CreatedAtUnixSecs int `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	// The last time the agent's main branch or agent-level settings were updated, in unix seconds
	UpdatedAtUnixSecs int                 `json:"updated_at_unix_secs" xml:"updated_at_unix_secs" required:"true"`
	AccessInfo        *ResourceAccessInfo `json:"access_info" xml:"access_info" required:"true"`
	// The time of the most recent call in unix seconds, null if no calls have been made
	LastCallTimeUnixSecs *int `json:"last_call_time_unix_secs,omitempty" xml:"last_call_time_unix_secs,omitempty"`
	// Whether the agent is archived
	Archived        *bool `json:"archived,omitempty" xml:"archived,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentSummaryResponseModel) GetAgentId() (value string) {
	if a == nil {
		return
	}
	return a.AgentId
}

func (a *AgentSummaryResponseModel) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AgentSummaryResponseModel) GetVoiceId() (value string) {
	if a == nil {
		return
	}
	return a.VoiceId
}

func (a *AgentSummaryResponseModel) GetTags() (value []string) {
	if a == nil {
		return
	}
	return a.Tags
}

func (a *AgentSummaryResponseModel) GetCreatedAtUnixSecs() (value int) {
	if a == nil {
		return
	}
	return a.CreatedAtUnixSecs
}

func (a *AgentSummaryResponseModel) GetUpdatedAtUnixSecs() (value int) {
	if a == nil {
		return
	}
	return a.UpdatedAtUnixSecs
}

func (a *AgentSummaryResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if a == nil {
		return
	}
	return a.AccessInfo
}

func (a *AgentSummaryResponseModel) GetLastCallTimeUnixSecs() (value *int) {
	if a == nil {
		return
	}
	return a.LastCallTimeUnixSecs
}

func (a *AgentSummaryResponseModel) GetArchived() (value *bool) {
	if a == nil {
		return
	}
	return a.Archived
}

func (a *AgentSummaryResponseModel) SetAgentId(agentId string) {
	a.AgentId = agentId
}

func (a *AgentSummaryResponseModel) SetName(name string) {
	a.Name = name
}

func (a *AgentSummaryResponseModel) SetVoiceId(voiceId string) {
	a.VoiceId = voiceId
}

func (a *AgentSummaryResponseModel) SetTags(tags []string) {
	a.Tags = tags
}

func (a *AgentSummaryResponseModel) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	a.CreatedAtUnixSecs = createdAtUnixSecs
}

func (a *AgentSummaryResponseModel) SetUpdatedAtUnixSecs(updatedAtUnixSecs int) {
	a.UpdatedAtUnixSecs = updatedAtUnixSecs
}

func (a *AgentSummaryResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	a.AccessInfo = accessInfo
}

func (a *AgentSummaryResponseModel) SetLastCallTimeUnixSecs(lastCallTimeUnixSecs *int) {
	a.LastCallTimeUnixSecs = lastCallTimeUnixSecs
}

func (a *AgentSummaryResponseModel) SetArchived(archived *bool) {
	a.Archived = archived
}

func (a *AgentSummaryResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentSummaryResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentSummaryResponseModel to string"
	}
	return string(jsonData)
}

func (a *AgentSummaryResponseModel) UnmarshalJSON(data []byte) error {
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

func (a *AgentSummaryResponseModel) MarshalJSON() ([]byte, error) {
	type embed AgentSummaryResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
