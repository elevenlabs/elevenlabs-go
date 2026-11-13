package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Every owner-editable field, including what callers must not hear.
type OwnerListServicesParams struct {
	ListKwargs      map[string]any `json:"list_kwargs,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OwnerListServicesParams) GetListKwargs() (value map[string]any) {
	if o == nil {
		return
	}
	return o.ListKwargs
}

func (o *OwnerListServicesParams) SetListKwargs(listKwargs map[string]any) {
	o.ListKwargs = listKwargs
}

func (o *OwnerListServicesParams) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OwnerListServicesParams) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OwnerListServicesParams to string"
	}
	return string(jsonData)
}

func (o *OwnerListServicesParams) UnmarshalJSON(data []byte) error {
	type unmarshaler OwnerListServicesParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OwnerListServicesParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o, "smb_tool_type")
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OwnerListServicesParams) MarshalJSON() ([]byte, error) {
	type embed OwnerListServicesParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
