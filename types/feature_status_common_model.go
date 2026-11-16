package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type FeatureStatusCommonModel struct {
	Enabled         *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Used            *bool `json:"used,omitempty" xml:"used,omitempty"`
	extraProperties map[string]interface{}
}

func (f *FeatureStatusCommonModel) GetEnabled() (value *bool) {
	if f == nil {
		return
	}
	return f.Enabled
}

func (f *FeatureStatusCommonModel) GetUsed() (value *bool) {
	if f == nil {
		return
	}
	return f.Used
}

func (f *FeatureStatusCommonModel) SetEnabled(enabled *bool) {
	f.Enabled = enabled
}

func (f *FeatureStatusCommonModel) SetUsed(used *bool) {
	f.Used = used
}

func (f *FeatureStatusCommonModel) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FeatureStatusCommonModel) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FeatureStatusCommonModel to string"
	}
	return string(jsonData)
}

func (f *FeatureStatusCommonModel) UnmarshalJSON(data []byte) error {
	type unmarshaler FeatureStatusCommonModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FeatureStatusCommonModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FeatureStatusCommonModel) MarshalJSON() ([]byte, error) {
	type embed FeatureStatusCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
