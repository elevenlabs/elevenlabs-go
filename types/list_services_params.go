package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ListServicesParams struct {
	ListKwargs      map[string]any `json:"list_kwargs,omitempty"`
	extraProperties map[string]interface{}
}

func (l *ListServicesParams) GetListKwargs() (value map[string]any) {
	if l == nil {
		return
	}
	return l.ListKwargs
}

func (l *ListServicesParams) SetListKwargs(listKwargs map[string]any) {
	l.ListKwargs = listKwargs
}

func (l *ListServicesParams) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListServicesParams) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListServicesParams to string"
	}
	return string(jsonData)
}

func (l *ListServicesParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ListServicesParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListServicesParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "smb_tool_type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListServicesParams) MarshalJSON() ([]byte, error) {
	type embed ListServicesParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
