package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioKeyMoment struct {
	TimestampMs     int    `json:"timestamp_ms" xml:"timestamp_ms" required:"true"`
	Type            string `json:"type" xml:"type" required:"true"`
	Description     string `json:"description" xml:"description" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AudioKeyMoment) GetTimestampMs() (value int) {
	if a == nil {
		return
	}
	return a.TimestampMs
}

func (a *AudioKeyMoment) GetType() (value string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AudioKeyMoment) GetDescription() (value string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AudioKeyMoment) SetTimestampMs(timestampMs int) {
	a.TimestampMs = timestampMs
}

func (a *AudioKeyMoment) SetType(type_ string) {
	a.Type = type_
}

func (a *AudioKeyMoment) SetDescription(description string) {
	a.Description = description
}

func (a *AudioKeyMoment) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioKeyMoment) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioKeyMoment to string"
	}
	return string(jsonData)
}

func (a *AudioKeyMoment) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AudioKeyMoment
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AudioKeyMoment(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioKeyMoment) MarshalJSON() ([]byte, error) {
	type embed AudioKeyMoment
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
