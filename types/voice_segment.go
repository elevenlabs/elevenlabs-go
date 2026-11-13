package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoiceSegment struct {
	// The voice ID used for this segment
	VoiceId string `json:"voice_id" xml:"voice_id" required:"true"`
	// Start time of this voice segment
	StartTimeSeconds float64 `json:"start_time_seconds" xml:"start_time_seconds" required:"true"`
	// End time of this voice segment
	EndTimeSeconds float64 `json:"end_time_seconds" xml:"end_time_seconds" required:"true"`
	// Start index in the characters array
	CharacterStartIndex int `json:"character_start_index" xml:"character_start_index" required:"true"`
	// End index in the characters array (exclusive)
	CharacterEndIndex int `json:"character_end_index" xml:"character_end_index" required:"true"`
	// Line of the dialogue (script) that this segment is a part of.
	DialogueInputIndex int `json:"dialogue_input_index" xml:"dialogue_input_index" required:"true"`
	extraProperties    map[string]interface{}
}

func (v *VoiceSegment) GetVoiceId() (value string) {
	if v == nil {
		return
	}
	return v.VoiceId
}

func (v *VoiceSegment) GetStartTimeSeconds() (value float64) {
	if v == nil {
		return
	}
	return v.StartTimeSeconds
}

func (v *VoiceSegment) GetEndTimeSeconds() (value float64) {
	if v == nil {
		return
	}
	return v.EndTimeSeconds
}

func (v *VoiceSegment) GetCharacterStartIndex() (value int) {
	if v == nil {
		return
	}
	return v.CharacterStartIndex
}

func (v *VoiceSegment) GetCharacterEndIndex() (value int) {
	if v == nil {
		return
	}
	return v.CharacterEndIndex
}

func (v *VoiceSegment) GetDialogueInputIndex() (value int) {
	if v == nil {
		return
	}
	return v.DialogueInputIndex
}

func (v *VoiceSegment) SetVoiceId(voiceId string) {
	v.VoiceId = voiceId
}

func (v *VoiceSegment) SetStartTimeSeconds(startTimeSeconds float64) {
	v.StartTimeSeconds = startTimeSeconds
}

func (v *VoiceSegment) SetEndTimeSeconds(endTimeSeconds float64) {
	v.EndTimeSeconds = endTimeSeconds
}

func (v *VoiceSegment) SetCharacterStartIndex(characterStartIndex int) {
	v.CharacterStartIndex = characterStartIndex
}

func (v *VoiceSegment) SetCharacterEndIndex(characterEndIndex int) {
	v.CharacterEndIndex = characterEndIndex
}

func (v *VoiceSegment) SetDialogueInputIndex(dialogueInputIndex int) {
	v.DialogueInputIndex = dialogueInputIndex
}

func (v *VoiceSegment) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VoiceSegment) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoiceSegment to string"
	}
	return string(jsonData)
}

func (v *VoiceSegment) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, v); err != nil {
		return err
	}
	type alias VoiceSegment
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*v = VoiceSegment(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VoiceSegment) MarshalJSON() ([]byte, error) {
	type embed VoiceSegment
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
