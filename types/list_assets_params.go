package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ListAssetsParams struct {
	ListKwargs      map[string]any `json:"list_kwargs,omitempty"`
	extraProperties map[string]interface{}
}

func (l *ListAssetsParams) GetListKwargs() (value map[string]any) {
	if l == nil {
		return
	}
	return l.ListKwargs
}

func (l *ListAssetsParams) SetListKwargs(listKwargs map[string]any) {
	l.ListKwargs = listKwargs
}

func (l *ListAssetsParams) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListAssetsParams) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListAssetsParams to string"
	}
	return string(jsonData)
}

func (l *ListAssetsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ListAssetsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListAssetsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "smb_tool_type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListAssetsParams) MarshalJSON() ([]byte, error) {
	type embed ListAssetsParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
