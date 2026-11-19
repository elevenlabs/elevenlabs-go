package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ListStaffParams struct {
	ListKwargs      map[string]any `json:"list_kwargs,omitempty"`
	extraProperties map[string]interface{}
}

func (l *ListStaffParams) GetListKwargs() (value map[string]any) {
	if l == nil {
		return
	}
	return l.ListKwargs
}

func (l *ListStaffParams) SetListKwargs(listKwargs map[string]any) {
	l.ListKwargs = listKwargs
}

func (l *ListStaffParams) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListStaffParams) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListStaffParams to string"
	}
	return string(jsonData)
}

func (l *ListStaffParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ListStaffParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListStaffParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "smb_tool_type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListStaffParams) MarshalJSON() ([]byte, error) {
	type embed ListStaffParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
