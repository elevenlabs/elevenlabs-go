package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type FinalAudioForTurn struct {
	// Signals the end of audio for a turn. Emitted after each explicit flush, new_turn/voice-switch boundary, and on connection close. All of the turn's audio, including any bytes buffered by streaming codecs (MP3, Opus, etc.), is delivered before this marker.
	IsFinalAudioForTurn *bool `json:"is_final_audio_for_turn,omitempty" xml:"is_final_audio_for_turn,omitempty"`
	extraProperties     map[string]interface{}
}

func (f *FinalAudioForTurn) GetIsFinalAudioForTurn() (value *bool) {
	if f == nil {
		return
	}
	return f.IsFinalAudioForTurn
}

func (f *FinalAudioForTurn) SetIsFinalAudioForTurn(isFinalAudioForTurn *bool) {
	f.IsFinalAudioForTurn = isFinalAudioForTurn
}

func (f *FinalAudioForTurn) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FinalAudioForTurn) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FinalAudioForTurn to string"
	}
	return string(jsonData)
}

func (f *FinalAudioForTurn) UnmarshalJSON(data []byte) error {
	type unmarshaler FinalAudioForTurn
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FinalAudioForTurn(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FinalAudioForTurn) MarshalJSON() ([]byte, error) {
	type embed FinalAudioForTurn
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
