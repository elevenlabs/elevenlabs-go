package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Every owner-editable field, working hours included.
type OwnerListStaffParams struct {
	ListKwargs      map[string]any `json:"list_kwargs,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OwnerListStaffParams) GetListKwargs() (value map[string]any) {
	if o == nil {
		return
	}
	return o.ListKwargs
}

func (o *OwnerListStaffParams) SetListKwargs(listKwargs map[string]any) {
	o.ListKwargs = listKwargs
}

func (o *OwnerListStaffParams) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OwnerListStaffParams) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OwnerListStaffParams to string"
	}
	return string(jsonData)
}

func (o *OwnerListStaffParams) UnmarshalJSON(data []byte) error {
	type unmarshaler OwnerListStaffParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OwnerListStaffParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o, "smb_tool_type")
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OwnerListStaffParams) MarshalJSON() ([]byte, error) {
	type embed OwnerListStaffParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
