package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A single behavioral rule for the agent.
type AgentRule struct {
	Title       StrippedStr `json:"title" xml:"title" required:"true"`
	Description StrippedStr `json:"description" xml:"description" required:"true"`
	// How strictly the receptionist follows a rule.
	Priority *RulePriority `json:"priority,omitempty" xml:"priority,omitempty"`
	// ID of the rule.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Unix timestamp (seconds since epoch)
	CreatedAtUnix   *int `json:"created_at_unix,omitempty" xml:"created_at_unix,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentRule) GetTitle() (value StrippedStr) {
	if a == nil {
		return
	}
	return a.Title
}

func (a *AgentRule) GetDescription() (value StrippedStr) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AgentRule) GetPriority() (value *RulePriority) {
	if a == nil {
		return
	}
	return a.Priority
}

func (a *AgentRule) GetId() (value *string) {
	if a == nil {
		return
	}
	return a.Id
}

func (a *AgentRule) GetCreatedAtUnix() (value *int) {
	if a == nil {
		return
	}
	return a.CreatedAtUnix
}

func (a *AgentRule) SetTitle(title StrippedStr) {
	a.Title = title
}

func (a *AgentRule) SetDescription(description StrippedStr) {
	a.Description = description
}

func (a *AgentRule) SetPriority(priority *RulePriority) {
	a.Priority = priority
}

func (a *AgentRule) SetId(id *string) {
	a.Id = id
}

func (a *AgentRule) SetCreatedAtUnix(createdAtUnix *int) {
	a.CreatedAtUnix = createdAtUnix
}

func (a *AgentRule) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentRule) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentRule to string"
	}
	return string(jsonData)
}

func (a *AgentRule) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentRule
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentRule(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentRule) MarshalJSON() ([]byte, error) {
	type embed AgentRule
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
