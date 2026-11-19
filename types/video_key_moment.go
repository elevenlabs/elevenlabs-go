package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VideoKeyMoment struct {
	TimestampMs     int    `json:"timestamp_ms" xml:"timestamp_ms" required:"true"`
	Type            string `json:"type" xml:"type" required:"true"`
	Description     string `json:"description" xml:"description" required:"true"`
	extraProperties map[string]interface{}
}

func (v *VideoKeyMoment) GetTimestampMs() (value int) {
	if v == nil {
		return
	}
	return v.TimestampMs
}

func (v *VideoKeyMoment) GetType() (value string) {
	if v == nil {
		return
	}
	return v.Type
}

func (v *VideoKeyMoment) GetDescription() (value string) {
	if v == nil {
		return
	}
	return v.Description
}

func (v *VideoKeyMoment) SetTimestampMs(timestampMs int) {
	v.TimestampMs = timestampMs
}

func (v *VideoKeyMoment) SetType(type_ string) {
	v.Type = type_
}

func (v *VideoKeyMoment) SetDescription(description string) {
	v.Description = description
}

func (v *VideoKeyMoment) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VideoKeyMoment) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VideoKeyMoment to string"
	}
	return string(jsonData)
}

func (v *VideoKeyMoment) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, v); err != nil {
		return err
	}
	type alias VideoKeyMoment
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*v = VideoKeyMoment(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VideoKeyMoment) MarshalJSON() ([]byte, error) {
	type embed VideoKeyMoment
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
