package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// List every customer-facing agent on the workspace.
//
// The assistant uses this whenever it needs to act on a specific customer-facing
// agent (rules, config edits, etc.) so it can pick the right “agent_id“ to pass
// to mutating tools. Mirrors the “list_services“ / “list_clients“
// pattern: read once, then mutate by id.
type ListCustomerFacingAgentsParams struct {
	extraProperties map[string]interface{}
}

func (l *ListCustomerFacingAgentsParams) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListCustomerFacingAgentsParams) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListCustomerFacingAgentsParams to string"
	}
	return string(jsonData)
}

func (l *ListCustomerFacingAgentsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ListCustomerFacingAgentsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListCustomerFacingAgentsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "smb_tool_type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListCustomerFacingAgentsParams) MarshalJSON() ([]byte, error) {
	type embed ListCustomerFacingAgentsParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
