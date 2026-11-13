package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ListTransferRulesParams struct {
	extraProperties map[string]interface{}
}

func (l *ListTransferRulesParams) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListTransferRulesParams) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListTransferRulesParams to string"
	}
	return string(jsonData)
}

func (l *ListTransferRulesParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ListTransferRulesParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListTransferRulesParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "smb_tool_type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListTransferRulesParams) MarshalJSON() ([]byte, error) {
	type embed ListTransferRulesParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
