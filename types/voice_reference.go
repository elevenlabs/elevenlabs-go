package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A voice this workspace can use.
type VoiceReference struct {
	type_ string
	// The ID of a voice in this workspace.
	VoiceId         string `json:"voice_id" xml:"voice_id" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (v *VoiceReference) GetVoiceId() (value string) {
	if v == nil {
		return
	}
	return v.VoiceId
}

func (v *VoiceReference) SetVoiceId(voiceId string) {
	v.VoiceId = voiceId
}

func (v *VoiceReference) Type() string {
	return v.type_
}

func (v *VoiceReference) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VoiceReference) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoiceReference to string"
	}
	return string(jsonData)
}

func (v *VoiceReference) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, v); err != nil {
		return err
	}
	type alias VoiceReference
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*v = VoiceReference(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "voice" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", v, "voice", *literals.Type)
	}
	v.type_ = "voice"
	extraProperties, err := internal.ExtractExtraProperties(data, *v, "type")
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VoiceReference) MarshalJSON() ([]byte, error) {
	type embed VoiceReference
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*v),
		Type:  "voice",
	}
	return json.Marshal(marshaler)
}
