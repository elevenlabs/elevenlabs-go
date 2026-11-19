package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ListProductsParams struct {
	ListKwargs      map[string]any `json:"list_kwargs,omitempty"`
	extraProperties map[string]interface{}
}

func (l *ListProductsParams) GetListKwargs() (value map[string]any) {
	if l == nil {
		return
	}
	return l.ListKwargs
}

func (l *ListProductsParams) SetListKwargs(listKwargs map[string]any) {
	l.ListKwargs = listKwargs
}

func (l *ListProductsParams) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListProductsParams) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListProductsParams to string"
	}
	return string(jsonData)
}

func (l *ListProductsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ListProductsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListProductsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "smb_tool_type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListProductsParams) MarshalJSON() ([]byte, error) {
	type embed ListProductsParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
