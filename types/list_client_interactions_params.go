package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ListClientInteractionsParams struct {
	extraProperties map[string]interface{}
}

func (l *ListClientInteractionsParams) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListClientInteractionsParams) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListClientInteractionsParams to string"
	}
	return string(jsonData)
}

func (l *ListClientInteractionsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ListClientInteractionsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListClientInteractionsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "smb_tool_type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListClientInteractionsParams) MarshalJSON() ([]byte, error) {
	type embed ListClientInteractionsParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
