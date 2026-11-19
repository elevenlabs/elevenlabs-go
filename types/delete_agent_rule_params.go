package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type DeleteAgentRuleParams struct {
	extraProperties map[string]interface{}
}

func (d *DeleteAgentRuleParams) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DeleteAgentRuleParams) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeleteAgentRuleParams to string"
	}
	return string(jsonData)
}

func (d *DeleteAgentRuleParams) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteAgentRuleParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteAgentRuleParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "smb_tool_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DeleteAgentRuleParams) MarshalJSON() ([]byte, error) {
	type embed DeleteAgentRuleParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
