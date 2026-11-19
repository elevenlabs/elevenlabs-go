package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Every owner-editable field, including what callers must not hear.
type OwnerListProductsParams struct {
	ListKwargs      map[string]any `json:"list_kwargs,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OwnerListProductsParams) GetListKwargs() (value map[string]any) {
	if o == nil {
		return
	}
	return o.ListKwargs
}

func (o *OwnerListProductsParams) SetListKwargs(listKwargs map[string]any) {
	o.ListKwargs = listKwargs
}

func (o *OwnerListProductsParams) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OwnerListProductsParams) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OwnerListProductsParams to string"
	}
	return string(jsonData)
}

func (o *OwnerListProductsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler OwnerListProductsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OwnerListProductsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o, "smb_tool_type")
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OwnerListProductsParams) MarshalJSON() ([]byte, error) {
	type embed OwnerListProductsParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
