package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioResponse struct {
	AudioEvent      *AudioResponseAudioEvent `json:"audio_event" xml:"audio_event" required:"true"`
	Type            *string                  `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AudioResponse) GetAudioEvent() (value *AudioResponseAudioEvent) {
	if a == nil {
		return
	}
	return a.AudioEvent
}

func (a *AudioResponse) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AudioResponse) SetAudioEvent(audioEvent *AudioResponseAudioEvent) {
	a.AudioEvent = audioEvent
}

func (a *AudioResponse) SetType(type_ *string) {
	a.Type = type_
}

func (a *AudioResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioResponse to string"
	}
	return string(jsonData)
}

func (a *AudioResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AudioResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AudioResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioResponse) MarshalJSON() ([]byte, error) {
	type embed AudioResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
