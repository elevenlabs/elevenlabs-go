package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Every owner-editable field, available hours included.
type OwnerListAssetsParams struct {
	ListKwargs      map[string]any `json:"list_kwargs,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OwnerListAssetsParams) GetListKwargs() (value map[string]any) {
	if o == nil {
		return
	}
	return o.ListKwargs
}

func (o *OwnerListAssetsParams) SetListKwargs(listKwargs map[string]any) {
	o.ListKwargs = listKwargs
}

func (o *OwnerListAssetsParams) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OwnerListAssetsParams) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OwnerListAssetsParams to string"
	}
	return string(jsonData)
}

func (o *OwnerListAssetsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler OwnerListAssetsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OwnerListAssetsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o, "smb_tool_type")
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OwnerListAssetsParams) MarshalJSON() ([]byte, error) {
	type embed OwnerListAssetsParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
