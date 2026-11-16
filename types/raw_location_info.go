package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A single branch/location extracted from a page.
type RawLocationInfo struct {
	Name            LenientOptionalStr `json:"name,omitempty" xml:"name,omitempty"`
	Address         LenientOptionalStr `json:"address,omitempty" xml:"address,omitempty"`
	OperatingHours  map[string]any     `json:"operating_hours,omitempty"`
	extraProperties map[string]interface{}
}

func (r *RawLocationInfo) GetName() (value LenientOptionalStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *RawLocationInfo) GetAddress() (value LenientOptionalStr) {
	if r == nil {
		return
	}
	return r.Address
}

func (r *RawLocationInfo) GetOperatingHours() (value map[string]any) {
	if r == nil {
		return
	}
	return r.OperatingHours
}

func (r *RawLocationInfo) SetName(name LenientOptionalStr) {
	r.Name = name
}

func (r *RawLocationInfo) SetAddress(address LenientOptionalStr) {
	r.Address = address
}

func (r *RawLocationInfo) SetOperatingHours(operatingHours map[string]any) {
	r.OperatingHours = operatingHours
}

func (r *RawLocationInfo) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RawLocationInfo) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RawLocationInfo to string"
	}
	return string(jsonData)
}

func (r *RawLocationInfo) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RawLocationInfo) MarshalJSON() ([]byte, error) {
	type embed RawLocationInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
