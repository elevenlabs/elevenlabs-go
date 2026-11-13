package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type DeleteAgentProcedureParams struct {
	extraProperties map[string]interface{}
}

func (d *DeleteAgentProcedureParams) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DeleteAgentProcedureParams) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeleteAgentProcedureParams to string"
	}
	return string(jsonData)
}

func (d *DeleteAgentProcedureParams) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteAgentProcedureParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteAgentProcedureParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "smb_tool_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DeleteAgentProcedureParams) MarshalJSON() ([]byte, error) {
	type embed DeleteAgentProcedureParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
