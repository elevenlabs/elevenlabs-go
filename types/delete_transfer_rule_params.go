package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type DeleteTransferRuleParams struct {
	extraProperties map[string]interface{}
}

func (d *DeleteTransferRuleParams) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DeleteTransferRuleParams) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeleteTransferRuleParams to string"
	}
	return string(jsonData)
}

func (d *DeleteTransferRuleParams) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteTransferRuleParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteTransferRuleParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "smb_tool_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DeleteTransferRuleParams) MarshalJSON() ([]byte, error) {
	type embed DeleteTransferRuleParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
